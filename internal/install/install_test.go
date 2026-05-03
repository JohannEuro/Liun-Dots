package install

import (
	"os"
	"path/filepath"
	"testing"
)

func seedSources(t *testing.T, root string) {
	t.Helper()
	files := map[string]string{
		filepath.Join("powershell", "Microsoft.PowerShell_profile.ps1"): "Set-StrictMode -Version Latest",
		filepath.Join("windows-terminal", "settings.json"):           "{\"theme\":\"dark\"}",
		filepath.Join("nvim", "init.lua"):                            "vim.opt.number = true",
	}
	for rel, body := range files {
		full := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestRunModeFullInstallsAllTargets(t *testing.T) {
	home := t.TempDir()
	source := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("LOCALAPPDATA", filepath.Join(home, "AppData", "Local"))
	seedSources(t, source)

	res, err := Run(source, ModeFull)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.Installed) != 3 {
		t.Fatalf("installed=%d want=3", len(res.Installed))
	}
	if len(res.Skipped) != 0 {
		t.Fatalf("skipped=%d want=0", len(res.Skipped))
	}
	if len(res.MissingFiles) != 0 {
		t.Fatalf("missing=%d want=0", len(res.MissingFiles))
	}
	if res.BackupDir == "" {
		t.Fatal("expected backup dir")
	}
}

func TestRunModeSafeSkipsExistingFiles(t *testing.T) {
	home := t.TempDir()
	source := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("LOCALAPPDATA", filepath.Join(home, "AppData", "Local"))
	seedSources(t, source)

	// Primera pasada para crear destino.
	_, err := Run(source, ModeFull)
	if err != nil {
		t.Fatalf("unexpected error in full install: %v", err)
	}

	res, err := Run(source, ModeSafe)
	if err != nil {
		t.Fatalf("unexpected error in safe install: %v", err)
	}
	if len(res.Installed) != 0 {
		t.Fatalf("installed=%d want=0", len(res.Installed))
	}
	if len(res.Skipped) != 3 {
		t.Fatalf("skipped=%d want=3", len(res.Skipped))
	}
}

func TestRunReportsMissingSourceWithoutFailing(t *testing.T) {
	home := t.TempDir()
	source := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("LOCALAPPDATA", filepath.Join(home, "AppData", "Local"))

	if err := os.MkdirAll(filepath.Join(source, "powershell"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "powershell", "Microsoft.PowerShell_profile.ps1"), []byte("ok"), 0o644); err != nil {
		t.Fatal(err)
	}

	res, err := Run(source, ModeFull)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.MissingFiles) != 2 {
		t.Fatalf("missing=%d want=2", len(res.MissingFiles))
	}
}

func TestRunReturnsErrorWhenDestinationCannotBeCreated(t *testing.T) {
	home := t.TempDir()
	source := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("LOCALAPPDATA", filepath.Join(home, "AppData", "Local"))
	seedSources(t, source)

	blocked := filepath.Join(home, "scoop", "persist")
	if err := os.MkdirAll(filepath.Dir(blocked), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(blocked, []byte("bloqueado"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := Run(source, ModeFull)
	if err == nil {
		t.Fatal("expected error when destination parent is a file")
	}
}
