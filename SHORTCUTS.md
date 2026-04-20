# SHORTCUTS (Flujo completo)

> Regla de oro: si un atajo no está acá, NO existe.

## 1) Shell

### Comandos/aliases

- `oc` → abre OpenCode en la carpeta actual (`opencode .`).
- `op-fast` → abre OpenCode sin plugins ni framework, para máxima velocidad (`opencode --pure`).
- `v` → abre Neovim (`nvim`).
- `d` → abre flujo de proyecto (`dev`): crea/entra a sesión tmux y abre nvim.
- `dev <ruta>` → lo mismo que `d`, pero para una ruta específica.

### Sesiones Rápidas con OpenCode (Preguntas generales)

- **Carpeta base:** `~/work/Sessions`
- **Uso:** Usá esta carpeta para abrir OpenCode cuando solo quieras hacer preguntas a la IA y no necesites indexar un proyecto. Al estar en Linux y vacía, inicia instantáneamente sin colgarse.
- **Ejemplo:** `cd ~/work/Sessions && opencode .` o `cd ~/work/Sessions && op-fast .`

### Atajos de edición y búsqueda en shell

- `Ctrl + R` → búsqueda fuzzy en historial con `fzf`.
- `Tab` → autocompletado inteligente.
- `↑ / ↓` → navegar historial.

## 2) tmux (prefix = `Ctrl + a`)

### Concepto clave

- **Prefix**: tecla de “comando” de tmux. Primero `Ctrl+a`, luego la tecla de acción.

### Splits y navegación

- `Ctrl+a` luego `v` → split vertical (izquierda/derecha).
- `Ctrl+a` luego `d` → split horizontal (arriba/abajo).
- `Ctrl+a` luego `h` → mover al panel izquierdo.
- `Ctrl+a` luego `j` → mover al panel inferior.
- `Ctrl+a` luego `k` → mover al panel superior.
- `Ctrl+a` luego `l` → mover al panel derecho.

### Config

- `Ctrl+a` luego `r` → recargar `~/.tmux.conf`.

### Clipboard en WSL

- En copy-mode (vi), al seleccionar con mouse y soltar, copia a Windows (`clip.exe`).
- En copy-mode (vi), `Enter` también copia a Windows (`clip.exe`).

## 3) Zellij (Alternativa a tmux)

### Concepto clave

- Zellij tiene diferentes "modos" indicados abajo (Normal, Locked, Pane, Tab, etc.).
- `Ctrl + g` → Bloquea (Locked) o desbloquea (Normal) los atajos de Zellij.

### Navegación entre pestañas (Tabs)

- `Alt + [`, `Alt + h`, `Alt + ←` (Flecha Izquierda) → Pestaña anterior.
- `Alt + ]`, `Alt + l`, `Alt + →` (Flecha Derecha) → Siguiente pestaña.

## 4) Gestión de Ventanas en Neovim

> Todo empieza con `Ctrl + w`. Luego usás la tecla de acción.

- `Ctrl + w` + `o` → Pantalla completa (cierra todo lo demás).
- `Ctrl + w` + `q` → Cerrar ventana actual.
- `Ctrl + w` + `H` → Mover ventana a la izquierda.
- `Ctrl + w` + `K` → Mover ventana hacia arriba.
- `Ctrl + w` + `+` → Aumentar tamaño.
- `Ctrl + w` + `-` → Disminuir tamaño.

*Nota: Usá `Ctrl + h/j/k/l` para navegar entre ventanas (saltar de una a otra), y `Ctrl + w` + `MAYÚSCULAS(H/J/K/L)` para mover la ventana físicamente.*

## 5) Neovim (LazyVim casi vanilla)

### Concepto clave

- **Leader** = `Space` (en LazyVim por defecto).

### Atajos agregados por tu setup

- `<leader>e` → abrir explorador de archivos `mini.files`.
- En `mini.files`:
  - `l` o `→` (Flecha Derecha) o `L` → Entrar a carpeta o abrir archivo.
  - `h` o `←` (Flecha Izquierda) o `H` → Volver a la carpeta padre (retroceder).
- `<leader>oo` → abrir OpenCode en un split de terminal, en la ruta del buffer actual.

### Atajos básicos imprescindibles

- `:w` → guardar archivo.
- `:q` → cerrar ventana actual.
- `:qa` → salir de Neovim.
- `:wq` → guardar y salir.

## 6) Flujo recomendado (paso a paso)

1. Abrís terminal en el proyecto (Siempre dentro de la partición de Linux `~` para máxima velocidad, NUNCA en `/mnt/c/`).
2. Ejecutás `d` (o `dev <ruta>`).
3. tmux te abre/adjunta sesión del proyecto y entra en nvim.
4. En nvim:
   - `<leader>e` para navegar carpetas/archivos,
   - `<leader>oo` para abrir opencode en split en esa ruta.

## 7) Notas técnicas

- OpenCode queda fijado por shim en `~/.local/bin/opencode` (y Linuxbrew).
- Clipboard WSL en tmux va a `clip.exe`.
- Las carpetas en Windows (`/mnt/c/`) son hasta 100x más lentas para indexar que el home de Linux (`~`). Evitá desarrollar o correr OpenCode ahí directo.