package ui

import (
	"os"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"liun-dots/internal/tools"
)

func TestIntegrationToggleFlow(t *testing.T) {
	isolateHome(t)
	m := NewModel()
	m.cursor = 2

	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	if m.state != stateIntegrations {
		t.Fatalf("expected integrations state")
	}

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	if m.openCode {
		t.Fatalf("expected OpenCode toggled off")
	}
}

func TestMainViewInSpanishAndLogo(t *testing.T) {
	isolateHome(t)
	m := NewModel()
	v := m.View()

	if !strings.Contains(v, "LIUNDOTS") {
		t.Fatalf("expected LIUNDOTS logo/header in view")
	}

	if !strings.Contains(v, "Instalar todo (sobrescribe + backup)") {
		t.Fatalf("expected Spanish main option text")
	}

	if strings.Contains(v, "Press q to quit") {
		t.Fatalf("unexpected English text in main view")
	}

	if !strings.Contains(v, "Todo listo") {
		t.Fatalf("expected premium status badge in Spanish")
	}

	if !strings.Contains(v, "╭") {
		t.Fatalf("expected main panel border")
	}
}

func TestMainViewShowsUpgradeBadgeWhenAvailable(t *testing.T) {
	isolateHome(t)
	m := NewModel()
	m.upgradeInfo = "Nueva versión v1.2.3 disponible"

	v := m.View()
	if !strings.Contains(v, "Nueva versión v1.2.3 disponible") {
		t.Fatalf("expected upgrade badge text in Spanish")
	}
}

func TestInstallOpensPrecheckBeforeRunning(t *testing.T) {
	isolateHome(t)
	m := NewModel()
	m.cursor = 0

	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)

	if m.state != statePrecheck {
		t.Fatalf("expected precheck state before installation")
	}

	v := m.View()
	if !strings.Contains(v, "Prerequisitos detectados") {
		t.Fatalf("expected precheck summary title")
	}
	if !strings.Contains(v, "Volver atrás") {
		t.Fatalf("expected back option")
	}
	if !strings.Contains(v, "Continuar con la instalación") {
		t.Fatalf("expected continue option")
	}
}

func TestPrecheckBackReturnsMainMenu(t *testing.T) {
	isolateHome(t)
	m := NewModel()
	m.cursor = 1

	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	if m.state != statePrecheck {
		t.Fatalf("expected precheck state")
	}

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	if m.state != stateMain {
		t.Fatalf("expected returning to main menu")
	}
}

func isolateHome(t *testing.T) {
	t.Helper()
	tmp := t.TempDir()
	oldHome := os.Getenv("HOME")
	oldUserProfile := os.Getenv("USERPROFILE")
	t.Cleanup(func() {
		_ = os.Setenv("HOME", oldHome)
		_ = os.Setenv("USERPROFILE", oldUserProfile)
	})
	_ = os.Setenv("HOME", tmp)
	_ = os.Setenv("USERPROFILE", tmp)
}

func TestPrecheckLinesShowsConcreteSuggestionForMissingCoreTool(t *testing.T) {
	lines := precheckLines([]tools.ToolStatus{{Label: "Git", Found: false}})
	if !strings.Contains(lines, "scoop install git") {
		t.Fatalf("expected concrete install suggestion for missing Git")
	}
}

func TestPrecheckImpactSummaryExplainsContinueBehavior(t *testing.T) {
	s := precheckImpactSummary([]tools.ToolStatus{{Label: "Git", Found: false}})
	if !strings.Contains(s, "si continuás") {
		t.Fatalf("expected continue behavior explained")
	}
}
