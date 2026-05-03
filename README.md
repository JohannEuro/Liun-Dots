# Liun Dots

`Liun-Dots` prepara un entorno Windows limpio, rápido y cómodo para trabajar con **PowerShell 7**, **Windows Terminal** y **Neovim**. La idea no es complicarte la vida: es darte una base seria, entendible y fácil de recuperar si algo no te convence.

---

## ¿Qué es exactamente?

Liun-Dots no intenta tocar toda tu PC ni convertirse en un framework gigante de dotfiles.

Hoy se enfoca en tres piezas concretas:

- tu perfil de **PowerShell**
- tu configuración de **Windows Terminal**
- tu `init.lua` de **Neovim**

Todo eso se aplica desde una **TUI en español**, con backup automático antes de sobrescribir archivos.

---

## ¿Qué modifica en tu PC?

Liun-Dots trabaja sobre estas rutas:

- `powershell/Microsoft.PowerShell_profile.ps1` → `%USERPROFILE%\scoop\persist\pwsh\Microsoft.PowerShell_profile.ps1`
- `windows-terminal/settings.json` → `%USERPROFILE%\scoop\persist\windows-terminal\settings\settings.json`
- `nvim/init.lua` → `%LOCALAPPDATA%\nvim\init.lua`

Antes de tocar cualquiera de esos archivos, crea un backup en:

- `%USERPROFILE%\.liun-dots\backups\<timestamp>\`

Si después querés volver atrás, la TUI tiene la opción **Recuperar backup (rollback)**.

### ¿Y si falla a mitad de instalación?

Puede pasar (por permisos, ruta bloqueada o archivo en uso). En ese caso:

1. Liun-Dots **ya dejó creado el backup** antes de escribir.
2. Revisá el mensaje de error en la TUI para detectar qué ruta falló.
3. Ejecutá **"Recuperar backup (rollback)"** para volver al último estado guardado.
4. Corregí la causa (permiso, herramienta faltante, archivo bloqueado) y volvé a correr la instalación.

En resumen: si hay falla parcial, la recuperación recomendada es rollback inmediato y reintento limpio.

---

## ¿Qué NO hace?

- No instala un IDE completo.
- No te obliga a usar IA.
- No corre procesos ocultos en segundo plano.
- No revisa internet cada vez que abrís PowerShell.
- No pisa tus archivos sin crear backup antes.

---

## Prerequisitos (PowerShell, copiar y pegar)

> Abrí PowerShell como usuario normal. No hace falta hacer todo como administrador.

### 1) Obligatorio

```powershell
Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
```

```powershell
iwr -useb get.scoop.sh | iex
```

```powershell
scoop install git
```

```powershell
scoop install pwsh
```

```powershell
scoop install windows-terminal
```

```powershell
scoop install neovim
```

### 2) Recomendado

OpenCode no es obligatorio para que Liun-Dots funcione, pero sí forma parte de la experiencia recomendada.

```powershell
scoop install opencode
```

### 3) Opcional (avanzado)

Gentle-AI es para quien quiera sumar una capa más avanzada de automatización y asistentes. No hace falta para usar Liun-Dots bien.

---

## Flujo Scoop esperado

```powershell
scoop bucket add liun-dots https://github.com/JohannEuro/Liun-Dots
```

```powershell
scoop install liun-dots
```

```powershell
liun-dots
```

Eso abre el instalador TUI paso a paso.

---

## Cómo se organiza Liun-Dots

### Obligatorio

- Scoop
- Git
- PowerShell 7
- Windows Terminal
- Neovim

### Recomendado

- OpenCode

### Opcional (avanzado)

- Gentle-AI

---

## Instalador TUI: qué vas a ver

- interfaz en español
- navegación con **flechas + Enter**
- pantalla limpia al entrar y al salir
- estilo visual sobrio, pensado para que se sienta cómodo y claro

Opciones principales:

- **Instalación completa**: sobrescribe configs soportadas + backup obligatorio.
- **Instalación segura**: solo copia faltantes y respeta lo que ya exista.
- **Herramientas recomendadas / IA**: toggles para OpenCode y Gentle-AI.
- **Actualizar Liun-Dots**: búsqueda manual de actualizaciones.
- **Recuperar backup (rollback)**: restaura el último backup.

### Chequeo previo de prerequisitos

Antes de ejecutar **Instalación completa** o **Instalación segura**, vas a ver una pantalla de chequeo con el estado de herramientas core como:

- PowerShell 7
- Windows Terminal
- Git
- Neovim

Si falta una herramienta, la TUI muestra un comando concreto para resolverlo. Por ejemplo:

```powershell
scoop install git
```

También te explica qué puede quedar incompleto si decidís continuar igual.

### Resumen final

Al finalizar instalación o recuperación de backup (rollback), Liun-Dots muestra:

- qué se aplicó
- qué se omitió
- dónde quedó el backup
- qué herramientas detectó en tu equipo
- cuál es el siguiente paso recomendado

También vas a ver si herramientas como **OpenCode** o **Gentle-AI** están detectadas en tu equipo, junto con su versión cuando sea posible.

---

## Actualizaciones (v1)

- No hay checks en background.
- No hay checks al iniciar PowerShell.
- La búsqueda es **manual** desde la opción de la TUI.
- La verificación usa **GitHub Releases API** y compara con la versión local de la app.
- El resultado se guarda en cache local por 24 horas en:
  - `%USERPROFILE%\.liun-dots\cache\update-check.json`

La cache guarda:

- `last_checked`
- `latest_version`

Si la cache sigue fresca y ya sabe que existe una versión nueva, el badge superior de la TUI puede avisarlo sin volver a hacer una consulta de red.

Para actualizar cuando haya una release nueva:

```powershell
scoop update liun-dots
```

---

## Si algo falta, ¿qué hago?

La regla es simple:

- si falta **Git** → `scoop install git`
- si falta **PowerShell 7** → `scoop install pwsh`
- si falta **Windows Terminal** → `scoop install windows-terminal`
- si falta **Neovim** → `scoop install neovim`
- si querés la experiencia recomendada con IA → `scoop install opencode`

Después de instalar lo que falta, abrí `liun-dots` otra vez y repetí el chequeo.

---

## Seguridad

Liun-Dots está pensado para Windows. Si estás en un entorno corporativo o sensible, revisá rutas, permisos y contenido antes de aplicarlo.
