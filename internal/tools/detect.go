package tools

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type versionProbe struct {
	args []string
}

type ToolStatus struct {
	Label   string
	Found   bool
	Version string
}

func DetectCoreTools() []ToolStatus {
	return []ToolStatus{
		detect("OpenCode", "opencode"),
		detect("Gentle-AI", "gentle-ai"),
		detectPowerShell7(),
		detect("Windows Terminal", "wt"),
		detect("Neovim", "nvim"),
		detect("Git", "git"),
	}
}

func DetectPrerequisitesCore() []ToolStatus {
	return []ToolStatus{
		detectPowerShell7(),
		detect("Windows Terminal", "wt"),
		detect("Neovim", "nvim"),
		detect("Git", "git"),
	}
}

func detectPowerShell7() ToolStatus {
	status := detect("PowerShell 7", "pwsh")
	if !status.Found {
		return status
	}
	major := parseMajorVersion(status.Version)
	if major > 0 && major < 7 {
		status.Found = false
		status.Version = "versión menor a 7"
	}
	return status
}

func detect(label, bin string) ToolStatus {
	if _, err := exec.LookPath(bin); err != nil {
		return ToolStatus{Label: label, Found: false}
	}

	version := detectVersion(bin)
	return ToolStatus{Label: label, Found: true, Version: version}
}

func detectVersion(bin string) string {
	for _, probe := range probesFor(bin) {
		cmd := exec.Command(bin, probe.args...)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			continue
		}
		line := strings.TrimSpace(strings.Split(out.String(), "\n")[0])
		if line == "" {
			continue
		}
		return normalizeVersionLine(line)
	}
	return "detectado"
}

func probesFor(bin string) []versionProbe {
	switch strings.ToLower(bin) {
	case "wt":
		return nil
	case "git":
		return []versionProbe{{args: []string{"--version"}}}
	case "pwsh", "nvim", "opencode", "gentle-ai":
		return []versionProbe{{args: []string{"--version"}}, {args: []string{"-v"}}}
	default:
		return []versionProbe{{args: []string{"--version"}}, {args: []string{"-v"}}}
	}
}

var spaces = regexp.MustCompile(`\s+`)

func normalizeVersionLine(s string) string {
	s = strings.TrimSpace(s)
	s = spaces.ReplaceAllString(s, " ")
	return s
}

func parseMajorVersion(s string) int {
	fields := strings.Fields(s)
	for _, f := range fields {
		f = strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(f, "v"), "V"), "version")
		parts := strings.Split(f, ".")
		if len(parts) == 0 {
			continue
		}
		major, err := strconv.Atoi(parts[0])
		if err == nil {
			return major
		}
	}
	return 0
}
