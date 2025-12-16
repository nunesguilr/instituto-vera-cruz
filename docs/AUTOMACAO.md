# Automação do Sistema - Apenas Hugo

## 🎯 Visão Geral

Este sistema foi projetado para funcionar **100% estaticamente** usando apenas o **Hugo** como gerador de sites. Não há necessidade de servidor backend (como Go, Node.js, etc.).

## 📊 Como Funciona

### Estrutura de Dados

Os dados são gerenciados através de arquivos JSON estáticos na pasta `data/`:

```
sociedade-veracruz/
├── data/
│   └── membros.json          # Dados dos membros
├── layouts/
│   ├── membros/
│   │   └── list.html         # Template da página de membros
│   └── partials/
│       └── components/
│           └── member_card.html  # Componente do card de membro
└── static/
    └── images/
        └── membros/          # Fotos dos membros
```

### Fluxo de Dados

1. **Dados em JSON**: Os membros são definidos em `data/membros.json`
2. **Hugo processa**: Durante o build, Hugo lê o arquivo JSON
3. **Templates renderizam**: Os templates Hugo (`list.html` e `member_card.html`) renderizam os dados
4. **Site estático gerado**: Hugo gera HTML estático na pasta `public/`

## 🔧 Gerenciando Membros

### Estrutura do membros.json

```json
{
  "reitores_e_colaboradores": [
    {
      "id": "1",
      "nome": "Nome do Membro",
      "titulo": "Cargo ou Especialidade",
      "biografia_curta": "Descrição breve (máx. 200 caracteres)",
      "link_externo": "https://link-externo.com",
      "imagem_perfil": "/images/membros/foto.jpg"
    }
  ]
}
```

### Campos Disponíveis

- **id** (obrigatório): Identificador único do membro
- **nome** (obrigatório): Nome completo
- **titulo** (obrigatório): Cargo ou especialidade
- **biografia_curta** (obrigatório): Breve descrição
- **link_externo** (opcional): Link para perfil externo ou página pessoal
- **imagem_perfil** (opcional): Caminho para a foto (relativo a `/static/`)

### Adicionando um Novo Membro

1. Abra o arquivo `data/membros.json`
2. Adicione um novo objeto no array `reitores_e_colaboradores`:

```json
{
  "id": "5",
  "nome": "Novo Membro",
  "titulo": "Especialista em Filosofia",
  "biografia_curta": "Descrição do novo membro...",
  "link_externo": "#",
  "imagem_perfil": "/images/membros/novo-membro.jpg"
}
```

3. Salve a foto do membro em `static/images/membros/`
4. Execute `hugo server -D` para ver as alterações

### Editando um Membro Existente

1. Localize o membro pelo `id` ou `nome` no arquivo `membros.json`
2. Edite os campos desejados
3. Salve o arquivo
4. O Hugo irá recarregar automaticamente (se estiver rodando com `hugo server`)

### Removendo um Membro

1. Localize o membro no arquivo `membros.json`
2. Delete o objeto completo (incluindo as chaves `{}`)
3. Certifique-se de manter a estrutura JSON válida (vírgulas, colchetes, etc.)
4. Salve o arquivo

## 🚀 Comandos Hugo

### Desenvolvimento Local

```bash
cd sociedade-veracruz
hugo server -D
```

- Inicia servidor de desenvolvimento
- Acesse: `http://localhost:1313`
- Recarrega automaticamente ao detectar mudanças

### Build de Produção

```bash
cd sociedade-veracruz
hugo
```

- Gera site estático na pasta `public/`
- Otimiza assets (CSS, JS)
- Pronto para deploy

### Limpar Cache

```bash
hugo --gc
```

- Remove arquivos não utilizados
- Limpa cache de recursos

## 📝 Padrão do Sistema

### Templates Hugo

O sistema usa a seguinte estrutura de templates:

1. **`layouts/membros/list.html`**: Página principal de membros
   - Carrega dados de `.Site.Data.membros`
   - Itera sobre `reitores_e_colaboradores`
   - Renderiza cada membro usando o partial

2. **`layouts/partials/components/member_card.html`**: Card individual
   - Recebe dados de um membro
   - Renderiza foto, nome, título e biografia
   - Adiciona link externo se disponível

### Acessando Dados nos Templates

```go-html-template
{{ $membrosData := .Site.Data.membros }}
{{ range $membrosData.reitores_e_colaboradores }}
  {{ partial "components/member_card.html" . }}
{{ end }}
```

## 🎨 Customização

### Alterando Estilos

Os estilos dos cards de membros estão em:
- `assets/scss/modules/_components.scss` (se existir)
- Ou inline nos templates

### Adicionando Novos Campos

1. Adicione o campo no `membros.json`:
```json
{
  "id": "1",
  "nome": "...",
  "novo_campo": "valor"
}
```

2. Acesse no template:
```go-html-template
{{ .novo_campo }}
```

## ✅ Vantagens desta Abordagem

- ✨ **Sem servidor backend**: Hospedagem simples e barata
- 🚀 **Performance**: Sites estáticos são extremamente rápidos
- 🔒 **Segurança**: Sem banco de dados ou APIs vulneráveis
- 💰 **Custo**: Hospedagem gratuita (Netlify, Vercel, GitHub Pages)
- 🔧 **Fácil manutenção**: Apenas edite arquivos JSON

## 📦 Deploy

O site pode ser deployado em qualquer serviço de hospedagem estática:

- **Netlify**: Conecte o repositório GitHub
- **Vercel**: Import do GitHub
- **GitHub Pages**: Configure GitHub Actions
- **Servidor tradicional**: Copie a pasta `public/`

## 🆘 Troubleshooting

### Membros não aparecem

1. Verifique se o JSON está válido (use um validador online)
2. Confirme que o caminho está correto: `data/membros.json`
3. Verifique se a estrutura `reitores_e_colaboradores` existe

### Imagens não carregam

1. Confirme que as imagens estão em `static/images/membros/`
2. Use caminhos absolutos: `/images/membros/foto.jpg`
3. Nomes de arquivos são case-sensitive

### Alterações não aparecem

1. Pare o servidor Hugo (`Ctrl+C`)
2. Execute `hugo --gc` para limpar cache
3. Inicie novamente com `hugo server -D`

---

*Sociedade de Vera Cruz - Sistema 100% Estático com Hugo*
