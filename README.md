# Instituto Vera Cruz - Sociedade de Vera Cruz

Bem-vindo ao repositório oficial da **Sociedade de Vera Cruz**. Este projeto é o portal digital de uma instituição focada em conhecimento, tradição e cultura, desenvolvido com uma estética clássica, sóbria e institucional.

## 🏛 Sobre o Projeto

O site foi construído para refletir a solidez e a seriedade da instituição. A identidade visual combina elementos de design clássico (tipografia serifada, cores terrosas e douradas) com técnicas modernas de desenvolvimento web (Glassmorphism, CSS Grid, Design Responsivo).

**Estética Principal:**
*   **Cores:** Vinho Profundo (`#8C1C1C`), Ouro Envelhecido (`#C8AA6E`), Fundo Pergaminho (`#FDFBF7`) e Tons de Café/Sépia.
*   **Tipografia:** Times New Roman (para títulos e assinaturas) e fontes de sistema limpas para leitura.
*   **Ambiente:** Uma mistura de "Vaticano encontra Universidade de Salamanca", com influências do Romantismo Alemão.

## 🛠 Tecnologias Utilizadas

*   **[Hugo](https://gohugo.io/):** Gerador de sites estáticos (SSG) rápido e flexível - **100% estático, sem necessidade de servidor backend**.
*   **SCSS (Sass):** Pré-processador CSS utilizado para estilização modular.
*   **HTML5 Semântico:** Estrutura acessível e otimizada para SEO.
*   **Git:** Controle de versão.
*   **Dados JSON:** Membros e conteúdo gerenciados através de arquivos estáticos em `data/`.

## 📂 Estrutura do Projeto

O código principal do site encontra-se na pasta `sociedade-veracruz`.

```
vera-cruz/
├── sociedade-veracruz/      # Raiz do Site Hugo
│   ├── assets/
│   │   └── scss/
│   │       ├── modules/     # Arquivos SCSS modulares (Partials)
│   │       │   ├── _variables.scss  # Cores e configs
│   │       │   ├── _buttons.scss    # Estilos de botões
│   │       │   ├── _layout.scss     # Grid e containers
│   │       │   └── ...
│   │       └── main.scss    # Arquivo principal (Importador)
│   ├── content/             # Conteúdo Markdown (Páginas e Artigos)
│   ├── layouts/             # Templates HTML (Partials, Shortcodes)
│   ├── static/              # Imagens e Favicons
│   └── hugo.toml            # Configuração do Hugo
└── README.md
```

## 🚀 Como Rodar o Projeto

### Início Rápido

```bash
# Clone o repositório
git clone https://github.com/nunesguilr/instituto-vera-cruz.git

# Entre na pasta do site
cd instituto-vera-cruz/sociedade-veracruz

# Inicie o servidor
hugo server -D
```

Acesse: **http://localhost:1313**

### 📚 Documentação Completa

Toda a documentação está organizada na pasta **[`docs/`](docs/):**

- **Para iniciantes:** [`docs/TUTORIAL-INICIANTES.md`](docs/TUTORIAL-INICIANTES.md)
- **Guia rápido:** [`docs/GUIA-RAPIDO.md`](docs/GUIA-RAPIDO.md)
- **Documentação técnica:** [`docs/AUTOMACAO.md`](docs/AUTOMACAO.md)
- **Índice completo:** [`docs/INDICE.md`](docs/INDICE.md)

## 💻 Guia de Desenvolvimento e Contribuição

### Editando Estilos (SCSS)
Não edite o CSS final. Trabalhamos exclusivamente na pasta `assets/scss/modules`.
*   Para ajustar cores: edite `_variables.scss`.
*   Para mudar botões: edite `_buttons.scss`.
*   Para layout estrutural: edite `_layout.scss`.

O `main.scss` serve apenas para importar esses módulos. O Hugo recompila as mudanças automaticamente.

### Criando Conteúdo
Para adicionar uma nova página ou artigo:
```bash
hugo new artigos/meu-novo-artigo.md
```
Isso criará um arquivo na pasta `content` com o front-matter padrão.

### Gerenciando Membros
Os membros são gerenciados através do arquivo `data/membros.json`. Para adicionar ou editar membros:

1. Abra o arquivo `sociedade-veracruz/data/membros.json`
2. Edite o array `reitores_e_colaboradores` seguindo o padrão:
   ```json
   {
     "id": "5",
     "nome": "Nome do Membro",
     "titulo": "Cargo ou Especialidade",
     "biografia_curta": "Breve descrição sobre o membro (máx. 200 caracteres)",
     "link_externo": "#",
     "imagem_perfil": "/images/membros/foto.jpg"
   }
   ```
3. As fotos devem ser colocadas em `static/images/membros/`
4. O Hugo irá automaticamente atualizar a página de membros

### Padrões de Código
*   **CSS:** Use unidades relativas (`rem`) em vez de `px` sempre que possível.
*   **Classes:** Utilize nomes semânticos (BEM leve) e variáveis SASS para cores.
*   **Imagens:** Coloque imagens na pasta `static/images` e use caminhos absolutos (`/images/foto.jpg`).
*   **Dados:** Todos os dados dinâmicos devem estar em arquivos JSON na pasta `data/`.

## 🤝 Contribuindo

1.  Crie uma branch para sua feature (`git checkout -b feature/minha-melhoria`).
2.  Faça commit das suas mudanças (`git commit -m 'Melhoria na home'`).
3.  Faça push para a branch (`git push origin feature/minha-melhoria`).
4.  Abra um Pull Request.

---
*Sociedade de Vera Cruz - Ad Veritatem Per Scientiam*

