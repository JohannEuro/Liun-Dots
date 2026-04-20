# Homebrew en Linuxbrew
eval (/home/linuxbrew/.linuxbrew/bin/brew shellenv)

# Binarios locales
fish_add_path ~/.local/bin

# Prompt limpio
set -g fish_greeting ""
set -gx EDITOR nvim
set -gx VISUAL nvim

# Inicialización de herramientas
if type -q zoxide
    zoxide init fish | source
end

if type -q fzf
    fzf --fish | source
end

if type -q atuin
    atuin init fish | source
end

if type -q carapace
    carapace _carapace | source
end

if type -q starship
    starship init fish | source
end

# Alias modernos
alias ls="eza --color=always --group-directories-first"
alias ll="eza -la --color=always --group-directories-first"
alias cat="bat --style=plain --paging=never"
alias tm="tmux"
alias zj="zellij"

# --- Paleta Nude Pro (Clara) ---
# set -l text 5E524A normal
# set -l subtle 8F7F73 brblack
# set -l rose C87C6A red
# set -l sand C9A67A yellow
# set -l sky 8CA3B0 blue
# set -l plum AF8EA6 magenta
# set -l mint 86A7A5 cyan
# set -l sel E7D9CF normal

# --- Paleta Nude Dark ---
set -l text d6c4b6 normal
set -l subtle 8099a6 brblack
set -l rose c97e7b red
set -l sand c2a383 yellow
set -l sky 8099a6 blue
set -l plum e8a8cf magenta
set -l mint 7d9e92 cyan
set -l sel a89a8e normal

set -g fish_color_normal $text
set -g fish_color_command $sky
set -g fish_color_keyword $plum
set -g fish_color_quote $sand
set -g fish_color_redirection $text
set -g fish_color_end $sand
set -g fish_color_error $rose
set -g fish_color_param $text
set -g fish_color_comment $subtle
set -g fish_color_selection --background=$sel
set -g fish_color_search_match --background=$sel
set -g fish_color_operator $mint
set -g fish_color_escape $plum
set -g fish_color_autosuggestion $subtle

set -g fish_pager_color_progress $subtle
set -g fish_pager_color_prefix $sky
set -g fish_pager_color_completion $text
set -g fish_pager_color_description $subtle

# Aliases de flujo
alias oc='opencode .'
alias op-fast='opencode --pure'
alias v='nvim'
alias d='zellij --layout work_nude.kdl'
alias dev='zellij --layout work_nude.kdl --cwd'
