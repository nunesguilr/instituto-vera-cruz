# 📁 Estrutura do Projeto - Sociedade de Vera Cruz

## Sistema: 100% Hugo Estático (SEM Servidor Backend)

```
instituto-vera-cruz/
├── .gitignore
├── README.md                    # Documentação principal
├── GUIA-RAPIDO.md              # Guia rápido de comandos
├── ESTRUTURA-PROJETO.md        # Este arquivo - Estrutura do projeto
│
└── sociedade-veracruz/         # 🏛 SITE HUGO
    │
    ├── hugo.toml               # ⚙️ Configuração do Hugo
    ├── AUTOMACAO.md            # 📖 Guia de automação
    │
    ├── archetypes/             # 📄 Templates de conteúdo
    │   └── default.md
    │
    ├── assets/                 # 🎨 Assets (SCSS, JS)
    │   └── scss/
    │       ├── modules/        # Módulos SCSS
    │       │   ├── _variables.scss
    │       │   ├── _buttons.scss
    │       │   ├── _layout.scss
    │       │   └── ...
    │       └── main.scss       # Arquivo principal
    │
    ├── content/                # 📝 Conteúdo Markdown
    │   ├── _index.md           # Homepage
    │   ├── artigos/            # Artigos
    │   ├── estudos/            # Estudos
    │   └── sobre/              # Sobre
    │
    ├── data/                   # 📊 DADOS JSON
    │   └── membros.json        # 👥 Dados dos membros
    │
    ├── layouts/                # 🎭 Templates HTML
    │   ├── _default/
    │   │   ├── baseof.html     # Template base
    │   │   ├── list.html       # Lista de páginas
    │   │   └── single.html     # Página individual
    │   │
    │   ├── index.html          # Homepage
    │   │
    │   ├── membros/
    │   │   └── list.html       # Página de membros
    │   │
    │   └── partials/
    │       ├── header.html
    │       ├── footer.html
    │       ├── hero.html
    │       └── components/
    │           └── member_card.html  # Card de membro
    │
    ├── static/                 # 📦 Arquivos estáticos
    │   └── images/
    │       ├── logo-white.png
    │       ├── hero-church.png
    │       └── membros/        # 📸 Fotos dos membros
    │           ├── guilherme.jpg
    │           ├── gabriel.png
    │           ├── eduardo.jpg
    │           └── jandilson.png
    │
    ├── public/                 # 🚀 Site compilado (gerado)
    │   └── (arquivos HTML/CSS/JS gerados)
    │
    └── resources/              # 🔧 Cache do Hugo
        └── _gen/
```

## 🎯 Arquivos Principais

### Configuração
- `hugo.toml` - Configuração do site (título, URL, parâmetros)

### Dados
- `data/membros.json` - Lista de membros (editável)

### Templates
- `layouts/membros/list.html` - Página de membros
- `layouts/partials/components/member_card.html` - Card individual

### Estilos
- `assets/scss/main.scss` - Importa todos os módulos
- `assets/scss/modules/` - Módulos SCSS organizados

### Conteúdo
- `content/` - Páginas em Markdown
- `static/` - Imagens e arquivos estáticos

## 🔄 Fluxo de Build

1. **Edição**: Edite arquivos `.md`, `.json`, `.scss`
2. **Hugo processa**: `hugo server -D` ou `hugo`
3. **Geração**: Hugo compila tudo para `public/`
4. **Deploy**: Suba a pasta `public/` para hospedagem

## ✅ Vantagens da Estrutura

- 📁 **Organizado**: Separação clara de dados, templates e estilos
- 🚀 **Rápido**: Build em milissegundos
- 🔧 **Manutenível**: Edite JSON, não código
- 🎨 **Flexível**: SCSS modular e reutilizável
- 📦 **Simples**: Sem dependências complexas

## 🛠 Edição Rápida

### Adicionar Membro
📝 Edite: `data/membros.json`

### Alterar Estilos
🎨 Edite: `assets/scss/modules/_variables.scss`

### Criar Conteúdo
✍️ Execute: `hugo new artigos/novo-artigo.md`

### Ver Mudanças
👁️ Execute: `hugo server -D`

---

**Sistema completamente estático - Sem servidor backend necessário! 🎉**
