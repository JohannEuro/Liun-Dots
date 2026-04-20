# Liun-Dots

¡Bienvenido a mi configuración personal! Entorno de desarrollo profesional, rápido y estético.

## 🛠️ Tecnologías Incluidas
Este setup configura un entorno completo basado en estas herramientas:

- **Shell:** `fish` (con `starship` para el prompt y `zoxide`/`fzf` para navegación).
- **Editor:** `neovim` (configurado con LazyVim).
- **Multiplexor:** `zellij` (gestión de ventanas y pestañas).
- **Terminal:** `wezterm`.
- **Utilidades:** `eza` (lista archivos), `bat` (lee archivos), `lazygit` (gestión git).
- **Lenguajes:** `Node.js`, `Rust` (vía rustup), `Go`, `Python`.

---

## 🚀 Guía de Instalación (One-Shot)

### 1. Clonar el repositorio
```bash
git clone https://github.com/JohannEuro/Liun-Dots.git ~/.Liun-Dots
```

### 2. Instalar dependencias según tu sistema

#### Opción A: Fedora Linux (Recomendado)
```bash
# Instalar herramientas básicas, lenguajes y compiladores
sudo dnf install -y \
  fish neovim zellij starship eza bat fzf zoxide gh nodejs \
  git gcc gcc-c++ make curl wget python3

# Instalar Rust (vía rustup)
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Instalar WezTerm (sigue las instrucciones oficiales: https://wezfurlong.org/wezterm/install/linux.html)
```

#### Opción B: Ubuntu / WSL
```bash
# Instalación vía Linuxbrew
brew install fish neovim zellij starship eza bat fzf gh zoxide node rust go
```

### 3. Instalación automática (PUM y listo)
Ejecutá el script incluido para configurar todo:
```bash
chmod +x ~/.Liun-Dots/install.sh
~/.Liun-Dots/install.sh
```

### 4. Finalización
1. **Cambiar shell a Fish:**
   ```bash
   chsh -s $(which fish)
   ```
2. **Reiniciar terminal:** Cerrá y abrila de nuevo.

---
⚠️ **Nota para el arquitecto:** Si tu usuario no es `liun`, recordá buscar y reemplazar `/home/liun/` por tu ruta de usuario en los archivos de configuración (`.lua`, `.toml`, etc.) después de instalar.
