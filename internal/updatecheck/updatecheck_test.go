package updatecheck

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNormalizeVersion(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "plain semver", in: "0.1.0", want: "0.1.0"},
		{name: "tag with v", in: "v0.1.0", want: "0.1.0"},
		{name: "trim spaces", in: " v1.2.3\n", want: "1.2.3"},
		{name: "empty", in: " ", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeVersion(tt.in); got != tt.want {
				t.Fatalf("got=%q want=%q", got, tt.want)
			}
		})
	}
}

func TestIsNewer(t *testing.T) {
	tests := []struct {
		name    string
		remote  string
		current string
		want    bool
	}{
		{name: "patch update", remote: "0.1.1", current: "0.1.0", want: true},
		{name: "minor update", remote: "0.2.0", current: "0.1.9", want: true},
		{name: "major update", remote: "1.0.0", current: "0.9.9", want: true},
		{name: "same version", remote: "0.1.0", current: "0.1.0", want: false},
		{name: "remote older", remote: "0.1.0", current: "0.2.0", want: false},
		{name: "tag prefix", remote: "v0.1.1", current: "0.1.0", want: true},
		{name: "invalid remote", remote: "latest", current: "0.1.0", want: false},
		{name: "invalid current", remote: "0.2.0", current: "dev", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNewer(tt.remote, tt.current); got != tt.want {
				t.Fatalf("got=%t want=%t", got, tt.want)
			}
		})
	}
}

func TestCacheRoundTripAndFreshness(t *testing.T) {
	tmp := t.TempDir()
	oldHome := os.Getenv("HOME")
	oldUserProfile := os.Getenv("USERPROFILE")
	t.Cleanup(func() {
		_ = os.Setenv("HOME", oldHome)
		_ = os.Setenv("USERPROFILE", oldUserProfile)
	})
	_ = os.Setenv("HOME", tmp)
	_ = os.Setenv("USERPROFILE", tmp)

	err := SaveCache(CacheData{LastChecked: time.Now().UTC(), LatestVersion: "1.2.3"})
	if err != nil {
		t.Fatalf("SaveCache error: %v", err)
	}

	got, err := LoadCache()
	if err != nil {
		t.Fatalf("LoadCache error: %v", err)
	}
	if got.LatestVersion != "1.2.3" {
		t.Fatalf("got latest=%q", got.LatestVersion)
	}

	status, ok, err := StatusFromCache()
	if err != nil {
		t.Fatalf("StatusFromCache error: %v", err)
	}
	if !ok {
		t.Fatalf("expected fresh cache")
	}
	if !status.Behind {
		t.Fatalf("expected behind=true for cache latest newer than current")
	}

	err = SaveCache(CacheData{LastChecked: time.Now().UTC().Add(-25 * time.Hour), LatestVersion: "9.9.9"})
	if err != nil {
		t.Fatalf("SaveCache stale error: %v", err)
	}
	_, ok, err = StatusFromCache()
	if err != nil {
		t.Fatalf("StatusFromCache stale error: %v", err)
	}
	if ok {
		t.Fatalf("expected stale cache ignored")
	}

	expectedFile := filepath.Join(tmp, ".liun-dots", "cache", "update-check.json")
	if _, err := os.Stat(expectedFile); err != nil {
		t.Fatalf("expected cache file at %s: %v", expectedFile, err)
	}
}

func TestCurrentVersionOverridePrecedence(t *testing.T) {
	oldBuild := BuildVersion
	oldEnv := os.Getenv("LIUN_DOTS_VERSION")
	t.Cleanup(func() {
		BuildVersion = oldBuild
		_ = os.Setenv("LIUN_DOTS_VERSION", oldEnv)
	})

	BuildVersion = "1.0.0"
	_ = os.Setenv("LIUN_DOTS_VERSION", "")
	if got := CurrentVersion(); got != "1.0.0" {
		t.Fatalf("expected build version, got %q", got)
	}

	_ = os.Setenv("LIUN_DOTS_VERSION", "2.0.0")
	if got := CurrentVersion(); got != "2.0.0" {
		t.Fatalf("expected env override, got %q", got)
	}
}
