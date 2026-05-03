package backup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreateWritesManifestAndCopiesExistingFiles(t *testing.T) {
	home := t.TempDir()
	t.Setenv("USERPROFILE", home)
	t.Setenv("HOME", home)

	target := filepath.Join(home, "target.txt")
	if err := os.WriteFile(target, []byte("abc"), 0o644); err != nil {
		t.Fatal(err)
	}

	dir, m, err := Create([]string{target, filepath.Join(home, "missing.txt")}, time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Entries) != 2 {
		t.Fatalf("entries=%d", len(m.Entries))
	}

	manifestPath := filepath.Join(dir, "manifest.json")
	raw, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatal(err)
	}
	var got Manifest
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if !got.Entries[0].Existed {
		t.Fatal("expected first entry existed=true")
	}
}

func TestSanitizeRelativePath(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "windows absolute", in: `C:\Users\A\x.txt`, want: filepath.Join("Users", "A", "x.txt")},
		{name: "unix style", in: `/Users/a/x.txt`, want: filepath.Join("Users", "a", "x.txt")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sanitizeRelativePath(tt.in)
			if got != tt.want {
				t.Fatalf("got=%q want=%q", got, tt.want)
			}
		})
	}
}
