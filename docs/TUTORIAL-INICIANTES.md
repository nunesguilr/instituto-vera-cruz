# 🎓 Tutorial Completo - Hugo para Iniciantes

## 🌟 Bem-vindo!

Este guia vai te ensinar a gerenciar o site da **Sociedade de Vera Cruz** usando Hugo, mesmo sem experiência em programação.

---

## 📋 Pré-requisitos

### O que você precisa ter instalado:

1. **Hugo Extended** (gerador de sites estáticos)
   - Windows: [Download aqui](https://github.com/gohugoio/hugo/releases)
   - Procure por: `hugo_extended_X.XX.X_windows-amd64.zip`
   - Extraia e adicione ao PATH

2. **Editor de Texto** (para editar arquivos)
   - Visual Studio Code (recomendado): [Download](https://code.visualstudio.com/)
   - Ou Notepad++: [Download](https://notepad-plus-plus.org/)

3. **Terminal/PowerShell** (já vem no Windows)

---

## 🚀 Passo 1: Iniciar o Site

### 1.1 Abrir o Terminal
1. Pressione `Win + R`
2. Digite: `powershell`
3. Pressione Enter

### 1.2 Navegar até a pasta do projeto
```powershell
cd C:\Users\guilherme\instituto-vera-cruz\sociedade-veracruz
```

### 1.3 Iniciar o servidor de desenvolvimento
```powershell
hugo server -D
```

### 1.4 Abrir no navegador
Acesse: **http://localhost:1313**

✅ **Pronto!** O site está rodando localmente.

---

## 👥 Passo 2: Adicionar um Novo Membro

### 2.1 Abrir o arquivo de dados
1. Navegue até: `sociedade-veracruz/data/`
2. Abra o arquivo: `membros.json`

### 2.2 Entender a estrutura
```json
{
  "reitores_e_colaboradores": [
    {
      "id": "1",
      "nome": "Nome do Membro",
      "titulo": "Cargo",
      "biografia_curta": "Descrição...",
      "link_externo": "#",
      "imagem_perfil": "/images/membros/foto.jpg"
    }
  ]
}
```

### 2.3 Adicionar novo membro
1. Copie um bloco completo (de `{` até `}`)
2. Cole ANTES do último `]`
3. Adicione uma vírgula `,` depois do bloco anterior
4. Edite os dados:

```json
{
  "reitores_e_colaboradores": [
    {
      "id": "1",
      "nome": "Membro Existente",
      "titulo": "Cargo",
      "biografia_curta": "...",
      "link_externo": "#",
      "imagem_perfil": "/images/membros/existente.jpg"
    },
    {
      "id": "2",
      "nome": "NOVO MEMBRO",
      "titulo": "Novo Cargo",
      "biografia_curta": "Nova descrição...",
      "link_externo": "#",
      "imagem_perfil": "/images/membros/novo.jpg"
    }
  ]
}
```

### 2.4 Adicionar foto do membro
1. Coloque a foto em: `sociedade-veracruz/static/images/membros/`
2. Nomeie algo como: `novo-membro.jpg`
3. No JSON, use: `"/images/membros/novo-membro.jpg"`

### 2.5 Salvar e ver resultado
1. Salve o arquivo `membros.json`
2. O navegador recarrega automaticamente!
3. Vá para: http://localhost:1313/membros

---

## ✏️ Passo 3: Editar um Membro Existente

### 3.1 Localizar o membro
1. Abra `membros.json`
2. Encontre o membro pelo nome ou ID

### 3.2 Editar informações
Você pode alterar qualquer campo:

```json
{
  "id": "3",
  "nome": "Nome Atualizado",           ← Mudar nome
  "titulo": "Novo Cargo",              ← Mudar cargo
  "biografia_curta": "Nova bio...",    ← Mudar descrição
  "link_externo": "https://novo.com",  ← Mudar link
  "imagem_perfil": "/images/membros/nova-foto.jpg"  ← Mudar foto
}
```

### 3.3 Salvar
1. Salve o arquivo
2. Veja as mudanças no navegador

---

## 🗑️ Passo 4: Remover um Membro

### 4.1 Localizar o membro
Abra `membros.json` e encontre o bloco do membro

### 4.2 Deletar o bloco completo
```json
{
  "reitores_e_colaboradores": [
    {
      "id": "1",
      "nome": "Manter Este",
      ...
    },
    {                              ← DELETE tudo isso
      "id": "2",                   ← incluindo as chaves
      "nome": "Remover Este",      ← e a vírgula depois
      ...
    },                             ← até aqui
    {
      "id": "3",
      "nome": "Manter Este Também",
      ...
    }
  ]
}
```

### 4.3 Corrigir vírgulas
⚠️ **IMPORTANTE**: Não pode ter vírgula depois do último item!

**ERRADO:**
```json
[
  { "id": "1", ... },
  { "id": "2", ... },  ← vírgula depois do último
]
```

**CORRETO:**
```json
[
  { "id": "1", ... },
  { "id": "2", ... }   ← SEM vírgula
]
```

---

## 🎨 Passo 5: Alterar Cores do Site

### 5.1 Abrir configuração
Arquivo: `hugo.toml`

### 5.2 Localizar parâmetros de cor
```toml
[params]
    color_primary = "#8C1C1C"      # Vinho profundo
    color_secondary = "#C8AA6E"    # Ouro envelhecido
    color_background = "#141210"   # Fundo escuro
    color_text = "#E7E5E4"         # Texto claro
```

### 5.3 Alterar cores
Use códigos hexadecimais (encontre em: [coolors.co](https://coolors.co))

```toml
color_primary = "#2E5090"    # Azul escuro
color_secondary = "#D4AF37"  # Dourado
```

### 5.4 Salvar e ver resultado
Hugo recarrega automaticamente!

---

## 📝 Passo 6: Criar um Novo Artigo

### 6.1 Usar comando Hugo
No terminal:
```powershell
hugo new artigos/meu-novo-artigo.md
```

### 6.2 Editar o artigo
1. Abra: `content/artigos/meu-novo-artigo.md`
2. Edite o cabeçalho (front matter):

```yaml
---
title: "Título do Artigo"
date: 2025-12-15
draft: false
---

Conteúdo do artigo começa aqui...

## Subtítulo

Mais conteúdo...
```

### 6.3 Publicar
1. Mude `draft: false`
2. Salve o arquivo
3. Veja em: http://localhost:1313/artigos

---

## 🚢 Passo 7: Gerar Site para Publicação

### 7.1 Parar o servidor
No terminal, pressione: `Ctrl + C`

### 7.2 Gerar site otimizado
```powershell
hugo --gc --minify
```

### 7.3 Resultado
Pasta gerada: `public/`

**Esta pasta contém TODO o site pronto para publicação!**

---

## 📦 Passo 8: Publicar o Site (Netlify)

### 8.1 Criar conta no Netlify
1. Acesse: [netlify.com](https://netlify.com)
2. Cadastre-se (grátis)

### 8.2 Conectar GitHub
1. No Netlify, clique "New site from Git"
2. Conecte sua conta GitHub
3. Selecione o repositório

### 8.3 Configurar build
```
Build command: hugo --gc --minify
Publish directory: public
```

### 8.4 Deploy!
1. Clique em "Deploy site"
2. Aguarde 1-2 minutos
3. Seu site está online! 🎉

---

## ❓ Resolução de Problemas

### Problema: Hugo não é reconhecido
**Solução**: Adicione Hugo ao PATH do Windows
1. Baixe Hugo Extended
2. Extraia para: `C:\Hugo\bin\`
3. Adicione ao PATH nas variáveis de ambiente

### Problema: JSON inválido
**Solução**: Use um validador
1. Copie o conteúdo de `membros.json`
2. Cole em: [jsonlint.com](https://jsonlint.com)
3. Corrija os erros apontados

### Problema: Imagem não aparece
**Solução**: Verifique o caminho
- ✅ Correto: `/images/membros/foto.jpg`
- ❌ Errado: `images/membros/foto.jpg` (sem `/` no início)
- ❌ Errado: `/static/images/...` (não inclua `static/`)

### Problema: Mudanças não aparecem
**Solução**:
1. Limpe o cache: `hugo --gc`
2. Reinicie o servidor
3. Force refresh no navegador: `Ctrl + Shift + R`

---

## 🎯 Comandos Essenciais

### Desenvolvimento
```powershell
# Iniciar servidor
hugo server -D

# Parar servidor
Ctrl + C
```

### Conteúdo
```powershell
# Novo artigo
hugo new artigos/titulo.md

# Novo estudo
hugo new estudos/titulo.md
```

### Publicação
```powershell
# Limpar cache
hugo --gc

# Gerar site
hugo --gc --minify
```

---

## 📚 Recursos Úteis

### Documentação
- [Hugo Docs](https://gohugo.io/documentation/)
- [Markdown Guide](https://www.markdownguide.org/)
- [JSON Tutorial](https://www.json.org/)

### Ferramentas Online
- [JSONLint](https://jsonlint.com/) - Validar JSON
- [Coolors](https://coolors.co/) - Escolher cores
- [PlaceHolder](https://placehold.co/) - Imagens temporárias

### Hospedagem Grátis
- [Netlify](https://netlify.com)
- [Vercel](https://vercel.com)
- [GitHub Pages](https://pages.github.com)

---

## ✅ Checklist Rápido

### Antes de Começar
- [ ] Hugo Extended instalado
- [ ] Editor de texto instalado
- [ ] Terminal aberto na pasta correta

### Ao Editar
- [ ] JSON válido (sem erros de sintaxe)
- [ ] Imagens na pasta `static/images/`
- [ ] Caminhos corretos (com `/` no início)
- [ ] Arquivo salvo antes de testar

### Antes de Publicar
- [ ] Todos os artigos com `draft: false`
- [ ] Imagens otimizadas (< 500KB cada)
- [ ] Build testado: `hugo --gc --minify`
- [ ] Pasta `public/` verificada

---

## 🎓 Conclusão

Parabéns! Agora você sabe:
- ✅ Iniciar o servidor Hugo
- ✅ Adicionar/editar/remover membros
- ✅ Criar artigos
- ✅ Alterar cores
- ✅ Publicar o site

**Dica final**: Sempre teste localmente antes de publicar!

---

*Sociedade de Vera Cruz - Hugo Made Simple* 🎉
