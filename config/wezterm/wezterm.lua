local wezterm = require 'wezterm'
local config = wezterm.config_builder()

-- 1. WSL Ubuntu como entorno por defecto
config.default_prog = { 'wsl.exe', '/home/linuxbrew/.linuxbrew/bin/fish', '-l' }
-- config.default_cwd = "~"  -- Comentado para evitar que use C:\Users\Admin

-- 2. ACELERACIÓN (Volvemos a OpenGL por estabilidad en Windows)
config.front_end = "OpenGL"

-- Ocultar advertencias de fuentes faltantes (Nerd Fonts)
config.warn_about_missing_glyphs = false

-- 3. RENDIMIENTO VISUAL
config.max_fps = 120
config.animation_fps = 120

-- 4. COMPATIBILIDAD WSL + TERMINFO
-- Evita: "missing or unsuitable terminal: wezterm" en shells dentro de WSL
if wezterm.target_triple:find("windows") then
  config.term = "xterm-256color"
else
  config.term = "wezterm"
end

config.enable_csi_u_key_encoding = true

-- 5. ESTÉTICA GENTLEMAN REPLICADA (KANAGAWA custom, no dependemos de schema nativo)
config.colors = {
  foreground = "#d6c4b6",
  background = "#1a1614",

  cursor_bg = "#c97e7b",
  cursor_fg = "#1a1614",
  cursor_border = "#c97e7b",

  selection_fg = "#1a1614",
  selection_bg = "#a89a8e",

  ansi = {
    "#1a1614",
    "#c97e7b",
    "#7d9e92",
    "#c2a383",
    "#8099a6",
    "#a6889d",
    "#8099a6",
    "#d6c4b6",
  },

  brights = {
    "#26211e",
    "#d98c89",
    "#8eb3a6",
    "#d6b492",
    "#91adc0",
    "#b899af",
    "#91adc0",
    "#e8d8c8",
  },
}

-- Colores de tab bar integrados al tema nude dark
config.colors.tab_bar = {
  background = "#1a1614",
  active_tab = {
    bg_color = "#26211e",
    fg_color = "#d6c4b6",
    intensity = "Bold",
  },
  inactive_tab = {
    bg_color = "#1a1614",
    fg_color = "#a89a8e",
  },
  inactive_tab_hover = {
    bg_color = "#26211e",
    fg_color = "#d6c4b6",
  },
  new_tab = {
    bg_color = "#1a1614",
    fg_color = "#a89a8e",
  },
  new_tab_hover = {
    bg_color = "#26211e",
    fg_color = "#d6c4b6",
  },
}

config.font_size = 10

-- Sin transparencias para evitar cuelgues en algunos drivers Windows
config.window_background_opacity = 1.0
config.win32_system_backdrop = "Disable"

-- Neovim/LSP undercurl
config.underline_thickness = 2
config.underline_position = -2

-- Scrollback amplio para debugging
config.scrollback_lines = 10000

-- Inputs estables (evita teclas muertas y compose extraños)
config.use_dead_keys = false
config.send_composed_key_when_left_alt_is_pressed = false
config.send_composed_key_when_right_alt_is_pressed = false

-- Ventana limpia
config.window_decorations = "TITLE | RESIZE"
config.enable_tab_bar = true
config.hide_tab_bar_if_only_one_tab = false
config.use_fancy_tab_bar = false
config.window_padding = {
  top = 6,
  right = 8,
  bottom = 6,
  left = 8,
}

return config
