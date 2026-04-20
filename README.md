# Liun-Dots

¡Bienvenido a mi configuración personal! Aquí encontrarás todo lo necesario para tener un entorno de desarrollo profesional, rápido y estético, igual al mío.

## 🚀 Guía de Instalación (Linux/WSL)

Esta guía asume que estás usando una distribución basada en Debian/Ubuntu (como WSL o Ubuntu nativo) y que usas **Linuxbrew**.

### 1. Requisitos previos
Primero, asegurate de tener instalado **Git** y **Linuxbrew**. Si no tenés Brew, instalalo desde [brew.sh](https://brew.sh/).

### 2. Clonar el repositorio
Abrí tu terminal y ejecutá:
```bash
git clone https://github.com/JohannEuro/Liun-Dots.git ~/.Liun-Dots
```

### 3. Instalar dependencias
Instalamos todo lo necesario vía `brew`:
```bash
brew install fish neovim zellij starship wezterm eza bat fzf gh zoxide
```

### 4. Vincular las configuraciones (Symlinks)
Para que Linux lea los archivos de este repo como si fueran suyos, creamos enlaces simbólicos:

```bash
# Crear carpetas de config si no existen
mkdir -p ~/.config

# Vincular cada configuración
ln -s ~/.Liun-Dots/config/nvim ~/.config/nvim
ln -s ~/.Liun-Dots/config/fish ~/.config/fish
ln -s ~/.Liun-Dots/config/wezterm ~/.config/wezterm
ln -s ~/.Liun-Dots/config/zellij ~/.config/zellij
ln -s ~/.Liun-Dots/config/starship.toml ~/.config/starship.toml
```

### 5. Finalización
1. **Establecer Fish como shell:**
   ```bash
   chsh -s $(which fish)
   ```
2. **Reiniciar terminal:** Cerrá todo y volvé a abrir para que cargue la configuración nueva.

---
⚠️ **Nota para el arquitecto:** Algunos archivos de configuración contienen rutas absolutas apuntando a `/home/liun`. Si tu nombre de usuario es distinto, asegúrate de buscar y reemplazar `/home/liun/` por `/home/tu-usuario/` en los archivos `.lua` o `.toml` antes de iniciar.
