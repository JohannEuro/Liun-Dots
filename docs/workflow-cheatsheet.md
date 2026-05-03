# Windows Workflow v1 - Cheatsheet

## Flujo mental principal

La regla base es simple:

- Si estas en PowerShell, usas comandos del shell.
- Si estas en Neovim/Oil, usas keymaps de Neovim.
- Si queres trabajar con IA en una carpeta, abris OpenCode en esa carpeta.
- Si una carpeta no aparece en el selector, agregas su root con `addroot`.

## Que usar segun donde estas

### Estoy en PowerShell y quiero abrir un proyecto
```powershell
cproj
```
Elegis proyecto y entras a la carpeta.

```powershell
vproj
```
Elegis proyecto y lo abris en Neovim.

```powershell
oproj
```
Elegis proyecto y lo abris en OpenCode.

### Estoy en PowerShell y ya estoy parado en la carpeta correcta
```powershell
nv
```
Abre Neovim en la carpeta actual.

```powershell
oc
```
Abre OpenCode en la carpeta actual.

```powershell
lg
```
Abre lazygit en la carpeta actual.

### Estoy en Neovim/Oil y quiero abrir OpenCode ahi
```text
Space o c
```
Abre OpenCode en un pane nuevo usando la carpeta actual de Neovim/Oil.

### Estoy en Neovim/Oil y quiero copiar la ruta actual
```text
Space y p
```
Copia la ruta actual al portapapeles.

### Quiero que una carpeta aparezca en croot/cproj/vproj/oproj
```powershell
addroot "D:\Ruta\A\Mis\Proyectos"
```

Despues podes verificar con:

```powershell
roots
```

### No se donde estoy
En PowerShell:
```powershell
pwd
```

En Neovim:
```vim
:pwd
```

En Oil, la carpeta visible es la carpeta actual que usan `Space oc` y `Space yp`.

---

## Shell
- `vim` -> abrir Neovim
- `nv` -> abrir Neovim en la carpeta actual (`nvim .`)
- `lg` -> abrir lazygit
- `ll` -> listar archivos
- `ff <texto>` -> buscar archivos con fd
- `grep <texto>` -> buscar texto con ripgrep
- `croot` -> elegir raiz de proyectos
- `cproj` -> elegir proyecto
- `vproj` -> elegir proyecto y abrirlo en Neovim
- `oc` -> abrir OpenCode en la carpeta actual
- `oc <ruta>` -> abrir OpenCode en una ruta especifica
- `oproj` -> elegir proyecto y abrirlo en OpenCode
- `roots` -> listar roots registradas
- `addroot <ruta>` -> agregar una root nueva para `croot/cproj/vproj/oproj`
- `openf <ruta>` -> abrir archivo/carpeta con la aplicacion por defecto de Windows
- `devcheck` -> verificar entorno base
- `reload-profile` -> recargar perfil actual

## Windows Terminal
### Panes
- `Ctrl + Shift + d` -> dividir pane
- `Ctrl + Shift + h/j/k/l` -> mover foco entre panes
- `Ctrl + Alt + h/j/k/l` -> redimensionar panes
- `Ctrl + Shift + z` -> zoom / unzoom del pane

### Tabs y ventanas utiles
- `Ctrl + Shift + t` -> nueva tab
- `Ctrl + Tab` -> tab siguiente
- `Ctrl + Shift + Tab` -> tab anterior
- `Ctrl + Shift + w` -> cerrar pane o tab

### Scrollback / leer mensajes sin mouse
- `Ctrl + Shift + Up` -> subir scroll
- `Ctrl + Shift + Down` -> bajar scroll
- `Ctrl + Shift + PgUp` -> subir una pagina
- `Ctrl + Shift + PgDn` -> bajar una pagina
- `Ctrl + Shift + Home` -> ir arriba del todo
- `Ctrl + Shift + End` -> ir abajo del todo
- `Ctrl + Shift + f` -> buscar texto en la terminal

## Neovim - modos
- `i` -> entrar en modo insert desde NORMAL
- `a` -> insertar despues del cursor
- `o` -> abrir linea debajo e insertar
- `Esc` -> salir de insert y volver a NORMAL
- `v` -> entrar en visual character-wise
- `V` -> entrar en visual por lineas
- `Ctrl + v` -> visual block
- `:` -> abrir linea de comandos

## Neovim - movimiento basico
- `h` -> izquierda
- `j` -> abajo
- `k` -> arriba
- `l` -> derecha
- `w` -> siguiente palabra
- `b` -> palabra anterior
- `0` -> inicio de linea
- `$` -> fin de linea
- `gg` -> inicio del archivo
- `G` -> fin del archivo
- `Ctrl + u` -> subir media pagina
- `Ctrl + d` -> bajar media pagina
- `zz` -> centrar linea actual

## Neovim - edicion basica
- `x` -> borrar caracter
- `dd` -> borrar linea
- `yy` -> copiar linea
- `p` -> pegar despues
- `u` -> undo
- `Ctrl + r` -> redo
- `:w` -> guardar
- `:q` -> salir
- `:wq` -> guardar y salir
- `:q!` -> salir sin guardar

## Neovim - busqueda
- `/texto` -> buscar hacia adelante
- `n` -> siguiente resultado
- `N` -> resultado anterior
- `Space Space` -> buscar archivos
- `Space /` -> buscar texto en proyecto
- `Space fb` -> buffers
- `Space fr` -> recientes
- `Space oc` -> abrir OpenCode en la carpeta actual / carpeta de Oil
- `Space yp` -> copiar ruta actual al portapapeles

## Neovim - ventanas
- `Ctrl + h/j/k/l` -> mover foco entre ventanas
- `Space q` -> cerrar ventana actual
- `Space Q` -> salir de todo sin guardar

## Oil explorer
- `Space e` -> abrir Oil
- `j / k` -> moverte por la lista
- `Enter` -> abrir archivo o entrar a carpeta
- `-` -> subir al directorio padre
- `g.` -> mostrar/ocultar archivos ocultos
- `q` -> cerrar Oil
- `:w` -> aplicar cambios si renombraste/editaste entradas
- `Space oc` -> abrir OpenCode en la carpeta que estas viendo en Oil
- `Space yp` -> copiar la ruta de la carpeta que estas viendo en Oil

## Abrir archivos del sistema desde terminal
- `openf archivo.pdf` -> abre el PDF con la app por defecto (si Edge es default, abre en Edge)
- `openf archivo.docx` -> abre con Word si Word es la app por defecto
- `openf archivo.xlsx` -> abre con Excel si Excel es la app por defecto
- `openf .` -> abre la carpeta actual en el Explorador de Windows

## Roots, proyectos y condiciones
### Que es una root
Una root es una carpeta base donde vivien tus proyectos.

Ejemplos:
- `D:\Storage\L\Programacion`
- `D:\Users\Admin\AndroidStudioProjects`

### `croot`
Muestra SOLO las roots registradas que existen en disco.

### `cproj`
Busca carpetas dentro de tus roots con profundidad maxima 2.
Tambien incluye la root misma como opcion.

### `vproj`
Hace lo mismo que `cproj`, pero abre el proyecto elegido en Neovim.

### `oproj`
Hace lo mismo que `cproj`, pero abre el proyecto elegido en OpenCode.

### Agregar una nueva root
```powershell
addroot "D:\ruta\a\mis\proyectos"
```

Despues de eso, esa ruta aparece en:
- `roots`
- `croot`
- `cproj`
- `vproj`
- `oproj`

### Ver roots actuales
```powershell
roots
```

### Abrir OpenCode
```powershell
oc
```
Abre OpenCode en la carpeta actual.

```powershell
oc "D:\Storage\L\Programacion"
```
Abre OpenCode en esa ruta.

## OpenCode / leer sin mouse
- Como OpenCode corre dentro de Windows Terminal, para subir y bajar historial usa los atajos de scrollback de arriba.
- Para buscar texto visible en la terminal usa `Ctrl + Shift + f`.

## Nota
Primero consolidar este flujo. Despues viene auditoria, QA y solo luego personalizacion pesada.

---

# Aprendizaje por fases

## Nivel 1 - Supervivencia diaria

### Shell
- `cproj` -> elegir proyecto
- `vproj` -> elegir proyecto y abrirlo en Neovim
- `nv` -> abrir `nvim .`
- `lg` -> abrir lazygit
- `openf <ruta>` -> abrir archivo/carpeta con la app por defecto

### Windows Terminal
- `Ctrl + Shift + d` -> dividir pane
- `Ctrl + Shift + h/j/k/l` -> mover foco entre panes

### Neovim
- `i` -> entrar en insert
- `Esc` -> volver a NORMAL
- `h j k l` -> moverte
- `:w` -> guardar
- `:q` -> salir
- `:q!` -> salir sin guardar
- `Space e` -> abrir Oil
- `q` -> cerrar Oil
- `Enter` en Oil -> abrir archivo/carpeta
- `-` en Oil -> subir carpeta

### Objetivo del Nivel 1
- Abrir proyecto
- Abrir archivo
- Editar
- Guardar
- Salir

## Nivel 2 - Productividad real

### Shell
- `ff <texto>` -> buscar archivos
- `grep <texto>` -> buscar texto
- `reload-profile` -> recargar config
- `devcheck` -> revisar entorno

### Windows Terminal
- `Ctrl + Alt + h/j/k/l` -> redimensionar panes
- `Ctrl + Shift + z` -> zoom pane
- `Ctrl + Shift + Up/Down` -> scroll sin mouse
- `Ctrl + Shift + f` -> buscar en terminal

### Neovim
- `w` / `b` -> moverte por palabras
- `0` / `$` -> inicio/fin de linea
- `gg` / `G` -> inicio/fin de archivo
- `u` -> undo
- `Ctrl + r` -> redo
- `/texto` -> buscar
- `n` / `N` -> siguiente / anterior
- `Space Space` -> buscar archivos
- `Space /` -> buscar texto en proyecto
- `Space fb` -> buffers
- `Space fr` -> recientes
- `Ctrl + h/j/k/l` -> mover foco entre ventanas

### Objetivo del Nivel 2
- Buscar rapido
- Moverte por el proyecto sin mouse
- Empezar a sentir velocidad real

## Nivel 3 - Control del editor

### Neovim modos
- `v` -> visual
- `V` -> visual por lineas
- `Ctrl + v` -> bloque visual

### Edicion
- `x` -> borrar caracter
- `dd` -> borrar linea
- `yy` -> copiar linea
- `p` -> pegar
- `o` -> nueva linea abajo e insertar
- `a` -> insertar despues del cursor

### Flujo de cierre
- `Space q` -> cerrar ventana
- `Space Q` -> salir de todo sin guardar
- `:wq` -> guardar y salir

### Objetivo del Nivel 3
- Editar mas rapido
- Depender menos del mouse
- Consolidar memoria muscular

## Regla de estudio
- Primero dominar Nivel 1
- Despues usar Nivel 2 todos los dias
- Recién cuando eso salga natural, pasar a Nivel 3
