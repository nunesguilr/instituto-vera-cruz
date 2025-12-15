package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Membro representa a estrutura de um membro
type Membro struct {
	Nome           string `json:"nome"`
	Titulo         string `json:"titulo"`
	Email          string `json:"email,omitempty"`
	BiografiaCurta string `json:"biografia_curta"`
	LinkExterno    string `json:"link_externo,omitempty"`
	Foto           string `json:"foto,omitempty"`
	DataCadastro   string `json:"data_cadastro"`
}

// Response representa a resposta padrão da API
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	dataDir     = "./data"
	membrosFile = "./data/membros.json"
	uploadDir   = "./static/images/membros"
	maxFileSize = 2 << 20 // 2MB
)

func main() {
	// Criar diretórios necessários
	createDirs()

	// Configurar rotas
	http.HandleFunc("/api/membros", corsMiddleware(handleMembros))
	http.HandleFunc("/api/membros/upload", corsMiddleware(handleUpload))
	
	// Servir arquivos estáticos do Hugo
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Iniciar servidor
	port := ":8080"
	log.Printf("🚀 Servidor iniciado em http://localhost%s\n", port)
	log.Printf("📊 API disponível em http://localhost%s/api/membros\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// createDirs cria os diretórios necessários se não existirem
func createDirs() {
	dirs := []string{dataDir, uploadDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Erro ao criar diretório %s: %v", dir, err)
		}
	}
	
	// Criar arquivo membros.json se não existir
	if _, err := os.Stat(membrosFile); os.IsNotExist(err) {
		emptyData := []Membro{}
		data, _ := json.MarshalIndent(emptyData, "", "  ")
		if err := os.WriteFile(membrosFile, data, 0644); err != nil {
			log.Fatalf("Erro ao criar arquivo membros.json: %v", err)
		}
	}
}

// corsMiddleware adiciona headers CORS
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}

// handleMembros gerencia as requisições de membros
func handleMembros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case "GET":
		getMembros(w, r)
	case "POST":
		createMembro(w, r)
	case "DELETE":
		deleteMembro(w, r)
	default:
		sendResponse(w, http.StatusMethodNotAllowed, false, "Método não permitido", nil)
	}
}

// getMembros retorna todos os membros
func getMembros(w http.ResponseWriter, r *http.Request) {
	membros, err := loadMembros()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, false, "Erro ao carregar membros", nil)
		return
	}
	
	sendResponse(w, http.StatusOK, true, "Membros carregados com sucesso", membros)
}

// createMembro cria um novo membro
func createMembro(w http.ResponseWriter, r *http.Request) {
	var novoMembro Membro
	
	if err := json.NewDecoder(r.Body).Decode(&novoMembro); err != nil {
		sendResponse(w, http.StatusBadRequest, false, "Dados inválidos", nil)
		return
	}
	
	// Validar campos obrigatórios
	if err := validateMembro(&novoMembro); err != nil {
		sendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}
	
	// Adicionar data de cadastro
	novoMembro.DataCadastro = time.Now().Format("2006-01-02")
	
	// Carregar membros existentes
	membros, err := loadMembros()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, false, "Erro ao carregar membros", nil)
		return
	}
	
	// Adicionar novo membro
	membros = append(membros, novoMembro)
	
	// Salvar no arquivo
	if err := saveMembros(membros); err != nil {
		sendResponse(w, http.StatusInternalServerError, false, "Erro ao salvar membro", nil)
		return
	}
	
	log.Printf("✅ Novo membro cadastrado: %s\n", novoMembro.Nome)
	sendResponse(w, http.StatusCreated, true, "Membro cadastrado com sucesso", novoMembro)
}

// deleteMembro remove um membro
func deleteMembro(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		sendResponse(w, http.StatusBadRequest, false, "Nome do membro não fornecido", nil)
		return
	}
	
	membros, err := loadMembros()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, false, "Erro ao carregar membros", nil)
		return
	}
	
	// Filtrar membros (remover o especificado)
	newMembros := []Membro{}
	found := false
	for _, m := range membros {
		if m.Nome != nome {
			newMembros = append(newMembros, m)
		} else {
			found = true
		}
	}
	
	if !found {
		sendResponse(w, http.StatusNotFound, false, "Membro não encontrado", nil)
		return
	}
	
	if err := saveMembros(newMembros); err != nil {
		sendResponse(w, http.StatusInternalServerError, false, "Erro ao salvar alterações", nil)
		return
	}
	
	log.Printf("🗑️  Membro removido: %s\n", nome)
	sendResponse(w, http.StatusOK, true, "Membro removido com sucesso", nil)
}

// handleUpload gerencia o upload de fotos
func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		sendResponse(w, http.StatusMethodNotAllowed, false, "Método não permitido", nil)
		return
	}
	
	// Limitar tamanho do upload
	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		sendResponse(w, http.StatusBadRequest, false, "Arquivo muito grande (máx. 2MB)", nil)
		return
	}
	
	file, header, err := r.FormFile("foto")
	if err != nil {
		sendResponse(w, http.StatusBadRequest, false, "Erro ao processar arquivo", nil)
		return
	}
	defer file.Close()
	
	// Validar tipo de arquivo
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		sendResponse(w, http.StatusBadRequest, false, "Apenas imagens são permitidas", nil)
		return
	}
	
	// Gerar nome único para o arquivo
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().Unix(), ext)
	filepath := filepath.Join(uploadDir, filename)
	
	// Criar arquivo de destino
	dst, err := os.Create(filepath)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, false, "Erro ao salvar arquivo", nil)
		return
	}
	defer dst.Close()
	
	// Copiar conteúdo
	if _, err := io.Copy(dst, file); err != nil {
		sendResponse(w, http.StatusInternalServerError, false, "Erro ao salvar arquivo", nil)
		return
	}
	
	// Retornar caminho da imagem
	imagePath := fmt.Sprintf("/images/membros/%s", filename)
	log.Printf("📷 Imagem enviada: %s\n", filename)
	
	sendResponse(w, http.StatusOK, true, "Upload realizado com sucesso", map[string]string{
		"path": imagePath,
	})
}

// loadMembros carrega os membros do arquivo JSON
func loadMembros() ([]Membro, error) {
	data, err := os.ReadFile(membrosFile)
	if err != nil {
		return nil, err
	}
	
	var membros []Membro
	if err := json.Unmarshal(data, &membros); err != nil {
		return nil, err
	}
	
	return membros, nil
}

// saveMembros salva os membros no arquivo JSON
func saveMembros(membros []Membro) error {
	data, err := json.MarshalIndent(membros, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(membrosFile, data, 0644)
}

// validateMembro valida os campos obrigatórios
func validateMembro(m *Membro) error {
	if strings.TrimSpace(m.Nome) == "" {
		return fmt.Errorf("nome é obrigatório")
	}
	
	if strings.TrimSpace(m.Titulo) == "" {
		return fmt.Errorf("título é obrigatório")
	}
	
	if strings.TrimSpace(m.BiografiaCurta) == "" {
		return fmt.Errorf("biografia é obrigatória")
	}
	
	if len(m.BiografiaCurta) > 200 {
		return fmt.Errorf("biografia deve ter no máximo 200 caracteres")
	}
	
	return nil
}

// sendResponse envia uma resposta JSON padronizada
func sendResponse(w http.ResponseWriter, status int, success bool, message string, data interface{}) {
	w.WriteHeader(status)
	response := Response{
		Success: success,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}