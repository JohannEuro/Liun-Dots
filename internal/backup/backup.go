package backup

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type Entry struct {
	OriginalPath string `json:"original_path"`
	BackupPath   string `json:"backup_path"`
	Existed      bool   `json:"existed"`
}

type Manifest struct {
	Timestamp string  `json:"timestamp"`
	Entries   []Entry `json:"entries"`
}

func BackupRoot() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".liun-dots", "backups"), nil
}

func Create(paths []string, now time.Time) (string, Manifest, error) {
	root, err := BackupRoot()
	if err != nil {
		return "", Manifest{}, err
	}
	ts := now.Format("20060102-150405")
	dir := filepath.Join(root, ts)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", Manifest{}, err
	}

	m := Manifest{Timestamp: now.UTC().Format(time.RFC3339), Entries: make([]Entry, 0, len(paths))}
	for _, p := range paths {
		rel := sanitizeRelativePath(p)
		dst := filepath.Join(dir, rel)
		entry := Entry{OriginalPath: p, BackupPath: dst}

		if _, err := os.Stat(p); err == nil {
			entry.Existed = true
			if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
				return "", Manifest{}, err
			}
			b, err := os.ReadFile(p)
			if err != nil {
				return "", Manifest{}, err
			}
			if err := os.WriteFile(dst, b, 0o644); err != nil {
				return "", Manifest{}, err
			}
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", Manifest{}, err
		}

		m.Entries = append(m.Entries, entry)
	}

	manifestPath := filepath.Join(dir, "manifest.json")
	raw, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", Manifest{}, err
	}
	if err := os.WriteFile(manifestPath, raw, 0o644); err != nil {
		return "", Manifest{}, err
	}

	return dir, m, nil
}

func RestoreLatest() (string, int, error) {
	root, err := BackupRoot()
	if err != nil {
		return "", 0, err
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return "", 0, err
	}
	if len(entries) == 0 {
		return "", 0, os.ErrNotExist
	}

	latest := ""
	for _, e := range entries {
		if e.IsDir() && e.Name() > latest {
			latest = e.Name()
		}
	}
	if latest == "" {
		return "", 0, os.ErrNotExist
	}

	manifestPath := filepath.Join(root, latest, "manifest.json")
	raw, err := os.ReadFile(manifestPath)
	if err != nil {
		return "", 0, err
	}
	var m Manifest
	if err := json.Unmarshal(raw, &m); err != nil {
		return "", 0, err
	}

	restored := 0
	for _, e := range m.Entries {
		if !e.Existed {
			_ = os.Remove(e.OriginalPath)
			continue
		}
		b, err := os.ReadFile(e.BackupPath)
		if err != nil {
			return "", restored, err
		}
		if err := os.MkdirAll(filepath.Dir(e.OriginalPath), 0o755); err != nil {
			return "", restored, err
		}
		if err := os.WriteFile(e.OriginalPath, b, 0o644); err != nil {
			return "", restored, err
		}
		restored++
	}

	return filepath.Join(root, latest), restored, nil
}

func sanitizeRelativePath(p string) string {
	vol := filepath.VolumeName(p)
	p = p[len(vol):]
	p = filepath.Clean(p)
	for len(p) > 0 && (p[0] == '\\' || p[0] == '/') {
		p = p[1:]
	}
	return p
}
