# Windows Workflow v1.0 Review

## QA Checklist

### Shell / Startup
- [ ] Abrir Windows Terminal desde cero y confirmar que entra directo a PowerShell 7
- [ ] Ejecutar `devcheck` y revisar que no falte ninguna tool core
- [ ] Ejecutar `reload-profile` y confirmar que no muestra errores
- [ ] Confirmar que el startup del shell se siente rapido y sin flashes raros

### Windows Terminal / Panes
- [ ] `Ctrl + Shift + d` divide el pane actual
- [ ] `Ctrl + Shift + h/j/k/l` mueve el foco entre panes
- [ ] `Ctrl + Alt + h/j/k/l` redimensiona panes
- [ ] `Ctrl + Shift + z` hace zoom y unzoom del pane
- [ ] Abrir/cerrar varios panes y confirmar que no se traba ni deja estados raros

### Project Flow
- [ ] `croot` lista roots reales y permite entrar rapido
- [ ] `cproj` encuentra proyectos reales dentro de las roots definidas
- [ ] `vproj` entra a un proyecto y abre `nvim .`
- [ ] Probar roots con espacios largos en path y confirmar que no se rompen

### Neovim
- [ ] `nvim .` abre rapido
- [ ] `Space Space` abre busqueda de archivos sin delay molesto
- [ ] `Space /` busca texto correctamente
- [ ] `Space e` abre Oil
- [ ] `q` cierra Oil
- [ ] `Space q` cierra ventana
- [ ] `Space Q` sale de todo sin dejarte atrapado
- [ ] `Ctrl + h/j/k/l` mueve entre ventanas dentro de Neovim

### Git / CLI
- [ ] `lg` abre lazygit sin errores
- [ ] `ff <texto>` encuentra archivos
- [ ] `grep <texto>` encuentra texto
- [ ] `vim` abre Neovim
- [ ] `nv` abre Neovim en la carpeta actual

---

## Review - Prioridades

### 1. Alta prioridad - Unificar package managers
Hoy el workflow mezcla Scoop y winget.

**Que mejorar**
- Idealmente dejar CLI core en un solo manager, preferentemente Scoop

**Por que importa**
- Simplifica mantenimiento
- Simplifica actualizaciones
- Hace mas limpio el futuro repo GitHub

**Tools a revisar primero**
- Neovim
- ripgrep
- fd
- lazygit

### 2. Alta prioridad - Validar ergonomia real de panes
La arquitectura ya funciona, pero falta validar si los atajos elegidos son los definitivos.

**Que mejorar**
- Confirmar si `Ctrl + Shift + h/j/k/l` y `Ctrl + Alt + h/j/k/l` se sienten naturales en uso diario

**Por que importa**
- Si los panes se sienten incomodos, el workflow pierde fluidez real

### 3. Alta prioridad - Consolidar flujo de proyectos
`croot`, `cproj` y `vproj` ya funcionan, pero son una pieza central del sistema.

**Que mejorar**
- Ver si las roots actuales son suficientes
- Ver si `max-depth 2` alcanza para todos tus proyectos

**Por que importa**
- Abrir proyectos rapido es parte del corazon de la experiencia "mantequilla"

### 4. Media prioridad - Pulir Neovim para experiencia de principiante
La base esta muy bien, pero todavia hay pequeñas fricciones de aprendizaje.

**Que mejorar**
- Ver si faltan keymaps obvios
- Ver si la salida del editor y de Oil es suficientemente intuitiva
- Ver si `which-key` ayuda lo suficiente o hay que guiar mejor algunos flujos

**Por que importa**
- La herramienta puede ser buena, pero si se siente crptica, te frena

### 5. Media prioridad - Limpiar configuracion de Windows Terminal
Ahora funciona, pero la config todavia esta orientada a dejar todo andando, no a quedar hermosa para repo.

**Que mejorar**
- Revisar defaults innecesarios
- Ordenar keybindings
- Dejar el archivo mas legible y mantenible

**Por que importa**
- Cuando esto vaya a GitHub, la claridad del setup tambien importa

### 6. Baja prioridad - Definir el siguiente bloque funcional de Neovim
Todavia no entra LSP, autocompletado, formatter ni testing integrado.

**Que mejorar**
- Diseñar una v1.1 de Neovim con crecimiento controlado

**Por que importa**
- La v1 actual es excelente como base, pero todavia no es un IDE completo

### 7. Baja prioridad - Preparar versionable para GitHub
No hacerlo todavia, pero ya hay que pensarlo con cabeza.

**Que tener en cuenta**
- Separar config de usuario vs config compartible
- Documentar instalacion
- Documentar decisiones de arquitectura
- Evitar rutas hardcodeadas si luego queres hacerlo portable

---

## Notas para la siguiente etapa
- No personalizar pesado antes de terminar QA
- No agregar AI todavia
- No agregar LSP todavia sin medir impacto
- Priorizar experiencia percibida antes que sumar features
- Cuando todo este en verde, recien ahi pasar a limpieza final y repo GitHub
