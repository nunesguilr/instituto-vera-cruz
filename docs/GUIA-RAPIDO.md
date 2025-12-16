# 🚀 Guia Rápido - Hugo Estático

## Comandos Principais

### Iniciar Servidor de Desenvolvimento
```bash
cd sociedade-veracruz
hugo server -D
```
Acesse: http://localhost:1313

### Build para Produção
```bash
cd sociedade-veracruz
hugo --gc
```
Site gerado em: `public/`

### Limpar Cache
```bash
hugo --gc --cleanDestinationDir
```

## 📝 Editar Membros

1. Abra: `data/membros.json`
2. Edite o JSON:
```json
{
  "reitores_e_colaboradores": [
    {
      "id": "1",
      "nome": "Nome",
      "titulo": "Cargo",
      "biografia_curta": "Descrição...",
      "link_externo": "#",
      "imagem_perfil": "/images/membros/foto.jpg"
    }
  ]
}
```
3. Salve e o Hugo recarrega automaticamente!

## 📸 Adicionar Fotos

1. Coloque a foto em: `static/images/membros/`
2. Use no JSON: `"imagem_perfil": "/images/membros/sua-foto.jpg"`

## 🎨 Editar Estilos

1. Vá para: `assets/scss/modules/`
2. Edite os arquivos `.scss`
3. Hugo recompila automaticamente

## ✅ Sistema 100% Estático

- ❌ **Não precisa** de servidor Go
- ❌ **Não precisa** de Node.js
- ❌ **Não precisa** de banco de dados
- ✅ **Só precisa** de Hugo
- ✅ **Deploy simples** em qualquer hospedagem estática

## 📦 Deploy

### Netlify (Recomendado)
1. Conecte seu repositório GitHub
2. Configure:
   - Build command: `hugo --gc --minify`
   - Publish directory: `public`
3. Deploy automático!

### Vercel
1. Import repositório
2. Framework: Hugo
3. Deploy!

### GitHub Pages
1. Adicione workflow em `.github/workflows/hugo.yml`
2. Push para GitHub
3. Ative Pages nas configurações

---

*Sistema simples, rápido e eficiente! 🚀*
