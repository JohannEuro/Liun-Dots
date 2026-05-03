package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type InstallTarget struct {
	Name        string
	SourceRel   string
	Destination string
}

func Expand(path string) (string, error) {
	resolved := path
	home, _ := os.UserHomeDir()
	localAppData := os.Getenv("LOCALAPPDATA")

	if home != "" {
		resolved = strings.ReplaceAll(resolved, "$HOME", home)
	}
	if localAppData != "" {
		resolved = strings.ReplaceAll(resolved, "$LOCALAPPDATA", localAppData)
		resolved = strings.ReplaceAll(resolved, "$env:LOCALAPPDATA", localAppData)
	}

	if strings.Contains(resolved, "$") {
		return "", fmt.Errorf("unresolved variable in path: %s", path)
	}

	return filepath.Clean(resolved), nil
}

func DefaultTargets() ([]InstallTarget, error) {
	t := []InstallTarget{
		{Name: "PowerShell profile", SourceRel: filepath.FromSlash("powershell/Microsoft.PowerShell_profile.ps1"), Destination: "$HOME\\scoop\\persist\\pwsh\\Microsoft.PowerShell_profile.ps1"},
		{Name: "Windows Terminal settings", SourceRel: filepath.FromSlash("windows-terminal/settings.json"), Destination: "$HOME\\scoop\\persist\\windows-terminal\\settings\\settings.json"},
		{Name: "Neovim init", SourceRel: filepath.FromSlash("nvim/init.lua"), Destination: "$env:LOCALAPPDATA\\nvim\\init.lua"},
	}

	for i := range t {
		expanded, err := Expand(t[i].Destination)
		if err != nil {
			return nil, err
		}
		t[i].Destination = expanded
	}

	return t, nil
}
