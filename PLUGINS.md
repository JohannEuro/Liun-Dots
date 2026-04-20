# PLUGINS Y FRAMEWORKS DEL ENTORNO

Este documento detalla todas las herramientas clave que componen tu entorno de desarrollo, su propósito y cómo se interrelacionan para darte un entorno rápido y estético.

## 1. Shell y Consola

### Starship 🚀

- **¿Qué es?** Un prompt de terminal ultrarrápido y súper personalizable escrito en Rust.
- **¿Qué hace?** Es el encargado de pintar la información en tu terminal. Muestra tu directorio actual, estado de Git, y a la derecha te muestra con **iconitos** las versiones de los lenguajes que estás usando (Node `⬢`, Go `🐹`, Bun `🥟`, etc.), cuánto tardó el último comando en ejecutarse (`⏳`), y la hora (`🕒`).
- **Configuración:** Vive en `~/.config/starship.toml`.

### Zoxide (z) ⚡

- **¿Qué es?** Un reemplazo inteligente y mucho más rápido para el comando `cd` de toda la vida.
- **¿Qué hace?** Recuerda los directorios que más visitás. En vez de escribir `cd ~/work/Angular-18-interview`, simplemente escribís `z angular` y te lleva al instante.

### FZF (Fuzzy Finder) 🔍

- **¿Qué es?** Un buscador difuso de línea de comandos.
- **¿Qué hace?** Se integra con tu historial. Cuando presionás `Ctrl + R`, en vez de la búsqueda fea por defecto, te abre una lista interactiva donde podés tipear pedazos de comandos viejos y te los encuentra al instante.

## 2. Multiplexores de Terminal

### Tmux / Zellij 🖥️

- **¿Qué son?** Multiplexores de terminal. (Zellij es más moderno y con interfaz, Tmux es más clásico y rápido).
- **¿Qué hacen?** Te permiten tener múltiples "ventanas" o "paneles" (splits) dentro de una sola ventana de tu terminal de Windows. Además, mantienen las sesiones vivas. Si cerrás la terminal y la volvés a abrir, podés "attachear" (reconectarte) y todo sigue exactamente como lo dejaste.

## 3. Editor de Código (Neovim + LazyVim) 📝

### Neovim (nvim)

- **¿Qué es?** El editor de texto hiper-optimizado que usás en la terminal.

### LazyVim 💤

- **¿Qué es?** Un framework (una configuración base) para Neovim.
- **¿Qué hace?** En lugar de que tengas que configurar Neovim desde cero instalando 50 plugins a mano, LazyVim ya viene con una base hermosa, temas estéticos, autocompletado, soporte para LSP (errores de código) y atajos listos para usar.

### mini.files (Plugin de Neovim) 📂

- **¿Qué es?** Un explorador de archivos.
- **¿Qué hace?** Cuando apretás `<leader>e` en Neovim, te abre una vista tipo columna para navegar tus carpetas y archivos como si fuera un explorador de sistema tradicional, pero sin salir del editor.

## 4. Inteligencia Artificial y Flujo de Trabajo

### OpenCode (oc / op-fast) 🤖

- **¿Qum es?** La herramienta principal de IA integrada directamente en tu terminal.
- **¿Qué hace?** Lee tu contexto, entiende tus proyectos y ejecuta comandos por vos.
  - El comando normal (`oc` o `opencode .`) levanta agentes y memoria persistente, ideal para proyectos completos.
  - El comando rápido (`op-fast`) levanta la versión pura, sin cargar los plugins pesados, ideal para preguntas rápidas en `~/work/Sessions`.

### Engram 🧠

- **¿Qué es?** El sistema de memoria persistente para OpenCode.
- **¿Qué hace?** Actúa como el cerebro de tu asistente. Recuerda las decisiones arquitectónicas que tomaron, los errores que ya resolvieron y los atajos que configuraste a lo largo de diferentes sesiones. Así la IA "no se olvida" de lo que hablaron ayer.
- **¿Cómo se correlacionan?** Engram corre como un servidor en segundo plano (vía MCP - Model Context Protocol) al que OpenCode se conecta al arrancar.

### Background Agents (Agent Teams) 🕵️

- **¿Qué es?** Un plugin de OpenCode.
- **¿Qué hace?** En el modo normal (Spec-Driven Development), coordina a un "Orquestador" que delega tareas a sub-agentes especializados (uno para explorar, otro para escribir código, otro para verificar).
