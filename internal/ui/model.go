package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"liun-dots/internal/backup"
	"liun-dots/internal/install"
	"liun-dots/internal/tools"
	"liun-dots/internal/updatecheck"
)

type state int

const (
	stateMain state = iota
	stateIntegrations
	statePrecheck
)

type installMode int

const (
	installNone installMode = iota
	installFull
	installSafe
)

type Model struct {
	cursor         int
	width          int
	state          state
	status         string
	upgradeInfo    string
	intCursor      int
	openCode       bool
	gentleAI       bool
	sourceRoot     string
	mainOptions    []string
	intOptions     []string
	preCursor      int
	preOptions     []string
	pendingInstall installMode
	toolsSnapshot  []tools.ToolStatus
	precheckCore   []tools.ToolStatus
}

var (
	badgeInfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("24")).
			Padding(0, 1)

	badgeSuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("22")).
			Padding(0, 1)

	badgeWarnStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("16")).
			Background(lipgloss.Color("214")).
			Padding(0, 1)

	badgeUpdateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("161")).
			Padding(0, 1)

	logoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("117")).
			Bold(true)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("254")).
			Bold(true)

	taglineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("246"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	statusLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("250")).
				Bold(true)

	statusValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252"))

	optionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("251"))

	selectedOptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("117")).
				Bold(true)

	activeCursorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("117")).
				Bold(true)

	successTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("78"))

	warnTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("214"))

	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1, 2)
)

func NewModel() Model {
	upgradeInfo := ""
	if s, ok, err := updatecheck.StatusFromCache(); err == nil && ok && s.Behind {
		upgradeInfo = fmt.Sprintf("Nueva versión v%s disponible", s.Remote)
	}

	m := Model{
		openCode:    true,
		upgradeInfo: upgradeInfo,
		sourceRoot:  detectSourceRoot(),
		mainOptions: []string{
			"Instalar todo (sobrescribe + backup)",
			"Instalación segura (solo faltantes)",
			"Recomendados de IA (OpenCode / Gentle-AI)",
			"Actualizar Liun-Dots (verificación manual)",
			"Recuperar backup (rollback)",
			"Salir",
		},
		intOptions: []string{"OpenCode (recomendado)", "Gentle-AI (opcional)", "Aplicar selección y volver"},
		preOptions: []string{"Volver atrás", "Continuar con la instalación"},
	}
	m.refreshToolSnapshot()
	return m
}

func (m Model) Init() tea.Cmd { return tea.ClearScreen }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.state == stateMain && m.cursor > 0 {
				m.cursor--
			}
			if m.state == stateIntegrations && m.intCursor > 0 {
				m.intCursor--
			}
			if m.state == statePrecheck && m.preCursor > 0 {
				m.preCursor--
			}
		case "down", "j":
			if m.state == stateMain && m.cursor < len(m.mainOptions)-1 {
				m.cursor++
			}
			if m.state == stateIntegrations && m.intCursor < len(m.intOptions)-1 {
				m.intCursor++
			}
			if m.state == statePrecheck && m.preCursor < len(m.preOptions)-1 {
				m.preCursor++
			}
		case "esc":
			if m.state == stateIntegrations {
				m.state = stateMain
			}
			if m.state == statePrecheck {
				m.returnToMainFromPrecheck()
			}
		case "enter":
			if m.state == stateMain {
				return m.handleMainSelection()
			}
			if m.state == statePrecheck {
				return m.handlePrecheckSelection()
			}
			m.handleIntegrationSelection()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}

func (m Model) View() string {
	if m.state == stateIntegrations {
		return m.integrationsView()
	}
	if m.state == statePrecheck {
		return m.precheckView()
	}
	b := strings.Builder{}
	b.WriteString(renderHeader(m.width))
	b.WriteString("\n\n")
	b.WriteString(statusLabelStyle.Render("Estado general: ") + m.topStatusBadge())
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Elegí una opción con flechas + Enter."))
	b.WriteString("\n\n")
	for i, opt := range m.mainOptions {
		cursor := "  "
		lineStyle := optionStyle
		if m.cursor == i {
			cursor = activeCursorStyle.Render("▸ ")
			lineStyle = selectedOptionStyle
		}
		b.WriteString(lineStyle.Render(fmt.Sprintf("%s%s", cursor, opt)) + "\n")
	}
	b.WriteString("\n")
	if m.status != "" {
		b.WriteString("\n" + statusLabelStyle.Render("Resultado: ") + statusValueStyle.Render(m.status) + "\n")
	}
	b.WriteString("\n" + helpStyle.Render("Tip: podés salir con q.") + "\n")
	return panelStyle.Render(b.String())
}

func (m Model) handleMainSelection() (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		m.pendingInstall = installFull
		m.preCursor = 0
		m.refreshPrecheckSnapshot()
		m.state = statePrecheck
	case 1:
		m.pendingInstall = installSafe
		m.preCursor = 0
		m.refreshPrecheckSnapshot()
		m.state = statePrecheck
	case 2:
		m.state = stateIntegrations
	case 3:
		s, err := updatecheck.Check()
		if err != nil {
			m.status = "Falló la verificación de actualización: " + err.Error()
			return m, nil
		}
		if s.Behind {
			m.upgradeInfo = fmt.Sprintf("Nueva versión v%s disponible", s.Remote)
			m.status = "Hay una actualización disponible. Ejecutá: scoop update liun-dots"
		} else {
			m.upgradeInfo = ""
			m.status = "Ya está actualizado. Todo OK."
		}
	case 4:
		dir, count, err := backup.RestoreLatest()
		if err != nil {
			m.status = "Falló la restauración: " + err.Error()
			return m, nil
		}
		m.refreshToolSnapshot()
		m.status = fmt.Sprintf("Recuperación de backup (rollback) finalizada desde %s (%d archivos restaurados).\n\n%s", dir, count, toolsSummary(m.toolsSnapshot))
	case 5:
		return m, tea.Quit
	}
	return m, nil
}

func toolsSummary(snapshot []tools.ToolStatus) string {
	b := strings.Builder{}
	b.WriteString("Estado local de herramientas detectadas:\n")
	for _, t := range snapshot {
		if t.Found {
			if t.Version != "" {
				b.WriteString(fmt.Sprintf("- %s: Sí (%s)\n", t.Label, t.Version))
			} else {
				b.WriteString(fmt.Sprintf("- %s: Sí\n", t.Label))
			}
			continue
		}
		b.WriteString(fmt.Sprintf("- %s: No detectado\n", t.Label))
	}
	return strings.TrimSpace(b.String())
}

func (m *Model) handleIntegrationSelection() {
	switch m.intCursor {
	case 0:
		m.openCode = !m.openCode
	case 1:
		m.gentleAI = !m.gentleAI
	case 2:
		m.status = fmt.Sprintf("Integraciones seleccionadas - OpenCode: %t, Gentle-AI: %t", m.openCode, m.gentleAI)
		m.state = stateMain
	}
}

func (m Model) integrationsView() string {
	b := strings.Builder{}
	b.WriteString(renderHeader(m.width))
	b.WriteString("\n\n")
	b.WriteString(statusLabelStyle.Render("Estado general: ") + m.topStatusBadge())
	b.WriteString("\n")
	b.WriteString(titleStyle.Render("Herramientas recomendadas / integraciones de IA"))
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Alterná con Enter y aplicá en la última opción."))
	b.WriteString("\n\n")
	for i, opt := range m.intOptions {
		cursor := "  "
		lineStyle := optionStyle
		if m.intCursor == i {
			cursor = activeCursorStyle.Render("▸ ")
			lineStyle = selectedOptionStyle
		}
		flag := ""
		if i == 0 {
			flag = boolFlag(m.openCode)
		}
		if i == 1 {
			flag = boolFlag(m.gentleAI)
		}
		b.WriteString(lineStyle.Render(fmt.Sprintf("%s%s %s", cursor, flag, opt)) + "\n")
	}
	b.WriteString("\n" + helpStyle.Render("Esc para volver.") + "\n")
	return panelStyle.Render(b.String())
}

func (m Model) precheckView() string {
	b := strings.Builder{}
	b.WriteString(renderHeader(m.width))
	b.WriteString("\n\n")
	b.WriteString(statusLabelStyle.Render("Chequeo previo: ") + m.precheckBadge())
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Acá ves qué está listo y qué falta antes de instalar."))
	b.WriteString("\n\n")
	b.WriteString(titleStyle.Render("Prerequisitos detectados"))
	b.WriteString("\n")
	b.WriteString(precheckLines(m.precheckCore))
	b.WriteString("\n\n")
	b.WriteString(precheckImpactSummary(m.precheckCore))
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Podés continuar igual, pero puede quedar algo incompleto hasta instalar lo faltante."))
	b.WriteString("\n\n")
	for i, opt := range m.preOptions {
		cursor := "  "
		lineStyle := optionStyle
		if m.preCursor == i {
			cursor = activeCursorStyle.Render("▸ ")
			lineStyle = selectedOptionStyle
		}
		b.WriteString(lineStyle.Render(fmt.Sprintf("%s%s", cursor, opt)) + "\n")
	}
	b.WriteString("\n" + helpStyle.Render("Esc para volver al menú principal."))
	return panelStyle.Render(b.String())
}

func precheckLines(core []tools.ToolStatus) string {
	b := strings.Builder{}
	for _, t := range core {
		if t.Found {
			version := ""
			if t.Version != "" {
				version = " (" + t.Version + ")"
			}
			b.WriteString(successTextStyle.Render(fmt.Sprintf("- ✅ %s: detectado%s", t.Label, version)) + "\n")
			continue
		}
		reason := ""
		if t.Version != "" {
			reason = " (" + t.Version + ")"
		}
		b.WriteString(warnTextStyle.Render(fmt.Sprintf("- ⚠️  %s: faltante%s", t.Label, reason)) + "\n")
		if hint := prereqHint(t.Label); hint != "" {
			b.WriteString(fmt.Sprintf("    ↳ Sugerencia: %s\n", hint))
		}
	}
	return strings.TrimSpace(b.String())
}

func (m Model) precheckBadge() string {
	for _, t := range m.precheckCore {
		if !t.Found {
			return badgeWarnStyle.Render("Faltan prerequisitos")
		}
	}
	return badgeSuccessStyle.Render("Todo listo para instalar")
}

func (m Model) handlePrecheckSelection() (tea.Model, tea.Cmd) {
	if m.preCursor == 0 {
		m.returnToMainFromPrecheck()
		return m, nil
	}

	mode := install.ModeSafe
	modeLabel := "segura"
	errLabel := "segura"
	if m.pendingInstall == installFull {
		mode = install.ModeFull
		modeLabel = "completa"
		errLabel = "completa"
	}

	res, err := install.Run(m.sourceRoot, mode)
	if err != nil {
		m.status = "Falló la instalación " + errLabel + ": " + err.Error()
		m.returnToMainFromPrecheck()
		return m, nil
	}
	m.refreshToolSnapshot()
	m.status = installationSummary(modeLabel, res, m.toolsSnapshot)
	m.returnToMainFromPrecheck()
	return m, nil
}

func (m *Model) returnToMainFromPrecheck() {
	m.state = stateMain
	m.preCursor = 0
	m.pendingInstall = installNone
}

func (m *Model) refreshToolSnapshot() {
	m.toolsSnapshot = tools.DetectCoreTools()
}

func (m *Model) refreshPrecheckSnapshot() {
	m.precheckCore = tools.DetectPrerequisitesCore()
}

func prereqHint(label string) string {
	switch label {
	case "Git":
		return "scoop install git"
	case "PowerShell 7":
		return "scoop install pwsh"
	case "Windows Terminal":
		return "scoop install windows-terminal"
	case "Neovim":
		return "scoop install neovim"
	default:
		return ""
	}
}

func precheckImpactSummary(core []tools.ToolStatus) string {
	missing := 0
	for _, t := range core {
		if !t.Found {
			missing++
		}
	}
	if missing == 0 {
		return "Todo listo: Liun-Dots puede aplicar la configuración con backup automático y modo seguro si lo elegís."
	}
	return fmt.Sprintf("Faltan %d herramienta(s) core: si continuás, Liun-Dots igual crea backup e instala archivos, pero algunas partes del entorno pueden no funcionar hasta que las instales.", missing)
}

func installationSummary(mode string, res install.Result, snapshot []tools.ToolStatus) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("Instalación %s finalizada.\n", mode))
	b.WriteString(fmt.Sprintf("- Backup creado en: %s\n", res.BackupDir))
	b.WriteString(fmt.Sprintf("- Archivos aplicados: %d\n", len(res.Installed)))
	b.WriteString(fmt.Sprintf("- Archivos omitidos: %d\n", len(res.Skipped)))
	b.WriteString(fmt.Sprintf("- Fuentes faltantes del paquete: %d\n", len(res.MissingFiles)))
	b.WriteString("\nSiguiente paso sugerido: abrí una terminal nueva y probá 'pwsh', 'git --version' y 'nvim --version'.\n\n")
	b.WriteString(toolsSummary(snapshot))
	return strings.TrimSpace(b.String())
}

func renderHeader(width int) string {
	logo := strings.Join([]string{
		"██╗     ██╗██╗   ██╗███╗   ██╗██████╗  ██████╗ ████████╗███████╗",
		"██║     ██║██║   ██║████╗  ██║██╔══██╗██╔═══██╗╚══██╔══╝██╔════╝",
		"██║     ██║██║   ██║██╔██╗ ██║██║  ██║██║   ██║   ██║   ███████╗",
		"██║     ██║██║   ██║██║╚██╗██║██║  ██║██║   ██║   ██║   ╚════██║",
		"███████╗██║╚██████╔╝██║ ╚████║██████╔╝╚██████╔╝   ██║   ███████║",
		"╚══════╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═════╝  ╚═════╝    ╚═╝   ╚══════╝",
	}, "\n")
	if width > 0 && width < 90 {
		logo = strings.Join([]string{
			"██╗     ██╗██╗   ██╗███╗   ██╗",
			"██║     ██║██║   ██║████╗  ██║",
			"██║     ██║██║   ██║██╔██╗ ██║",
			"██║     ██║██║   ██║██║╚██╗██║",
			"███████╗██║╚██████╔╝██║ ╚████║",
			"╚══════╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝",
			"██████╗  ██████╗ ████████╗███████╗",
			"██╔══██╗██╔═══██╗╚══██╔══╝██╔════╝",
			"██║  ██║██║   ██║   ██║   ███████╗",
			"██║  ██║██║   ██║   ██║   ╚════██║",
			"██████╔╝╚██████╔╝   ██║   ███████║",
			"╚═════╝  ╚═════╝    ╚═╝   ╚══════╝",
		}, "\n")
	}

	accent := helpStyle.Render("◆") + " " + titleStyle.Render("LIUNDOTS") + "" + helpStyle.Render("◆")
	tagline := taglineStyle.Render("Instalador guiado, claro y profesional para PowerShell, Windows Terminal y Neovim")
	header := logoStyle.Render(logo) + "\n" + accent + "\n" + tagline
	return header
}

func boolFlag(v bool) string {
	if v {
		return "[x]"
	}
	return "[ ]"
}

func (m Model) topStatusBadge() string {
	if m.upgradeInfo != "" {
		return badgeUpdateStyle.Render(m.upgradeInfo)
	}
	return badgeInfoStyle.Render("Todo listo")
}

func detectSourceRoot() string {
	exe, err := os.Executable()
	if err == nil {
		base := filepath.Dir(exe)
		if hasSources(base) {
			return base
		}
	}
	cwd, err := os.Getwd()
	if err == nil && hasSources(cwd) {
		return cwd
	}
	return "."
}

func hasSources(root string) bool {
	_, aErr := os.Stat(filepath.Join(root, "powershell", "Microsoft.PowerShell_profile.ps1"))
	_, bErr := os.Stat(filepath.Join(root, "windows-terminal", "settings.json"))
	_, cErr := os.Stat(filepath.Join(root, "nvim", "init.lua"))
	return aErr == nil && bErr == nil && cErr == nil
}
