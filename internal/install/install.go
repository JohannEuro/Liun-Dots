package install

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"liun-dots/internal/backup"
	"liun-dots/internal/config"
)

type Mode int

const (
	ModeFull Mode = iota
	ModeSafe
)

type Result struct {
	Mode         Mode
	BackupDir    string
	Installed    []string
	Skipped      []string
	MissingFiles []string
}

func Run(sourceRoot string, mode Mode) (Result, error) {
	targets, err := config.DefaultTargets()
	if err != nil {
		return Result{}, err
	}

	paths := make([]string, 0, len(targets))
	for _, t := range targets {
		paths = append(paths, t.Destination)
	}
	bkpDir, _, err := backup.Create(paths, time.Now())
	if err != nil {
		return Result{}, err
	}

	res := Result{Mode: mode, BackupDir: bkpDir}

	for _, t := range targets {
		src := filepath.Join(sourceRoot, t.SourceRel)
		if _, err := os.Stat(src); err != nil {
			res.MissingFiles = append(res.MissingFiles, t.SourceRel)
			continue
		}

		if mode == ModeSafe {
			if _, err := os.Stat(t.Destination); err == nil {
				res.Skipped = append(res.Skipped, t.Destination)
				continue
			}
		}

		b, err := os.ReadFile(src)
		if err != nil {
			return res, err
		}
		if err := os.MkdirAll(filepath.Dir(t.Destination), 0o755); err != nil {
			return res, err
		}
		if err := os.WriteFile(t.Destination, b, 0o644); err != nil {
			return res, err
		}
		res.Installed = append(res.Installed, fmt.Sprintf("%s -> %s", t.SourceRel, t.Destination))
	}

	return res, nil
}
