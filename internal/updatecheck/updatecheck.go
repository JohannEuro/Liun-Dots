package updatecheck

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const defaultVersion = "0.1.0"

var BuildVersion = ""

type Status struct {
	Current string
	Remote  string
	Behind  bool
}

type CacheData struct {
	LastChecked   time.Time `json:"last_checked"`
	LatestVersion string    `json:"latest_version"`
}

const CacheTTL = 24 * time.Hour

func Check() (Status, error) {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/JohannEuro/Liun-Dots/releases/latest")
	if err != nil {
		return Status{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Status{Current: CurrentVersion()}, fmt.Errorf("todavía no hay una release publicada en GitHub")
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return Status{}, fmt.Errorf("la verificación con GitHub falló: %s", resp.Status)
	}

	var payload struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return Status{}, err
	}

	remote := NormalizeVersion(payload.TagName)
	current := NormalizeVersion(CurrentVersion())
	status := Status{Current: current, Remote: remote, Behind: IsNewer(remote, current)}
	_ = SaveCache(CacheData{LastChecked: time.Now().UTC(), LatestVersion: remote})
	return status, nil
}

func StatusFromCache() (Status, bool, error) {
	cache, err := LoadCache()
	if err != nil {
		if os.IsNotExist(err) {
			return Status{}, false, nil
		}
		return Status{}, false, err
	}
	if cache.LastChecked.IsZero() || time.Since(cache.LastChecked) > CacheTTL {
		return Status{}, false, nil
	}

	current := NormalizeVersion(CurrentVersion())
	remote := NormalizeVersion(cache.LatestVersion)
	if remote == "" {
		return Status{}, false, nil
	}

	return Status{Current: current, Remote: remote, Behind: IsNewer(remote, current)}, true, nil
}

func CurrentVersion() string {
	if v := NormalizeVersion(os.Getenv("LIUN_DOTS_VERSION")); v != "" {
		return v
	}
	if v := NormalizeVersion(BuildVersion); v != "" {
		return v
	}
	return defaultVersion
}

func LoadCache() (CacheData, error) {
	file, err := cacheFilePath()
	if err != nil {
		return CacheData{}, err
	}
	b, err := os.ReadFile(file)
	if err != nil {
		return CacheData{}, err
	}
	var c CacheData
	if err := json.Unmarshal(b, &c); err != nil {
		return CacheData{}, err
	}
	return c, nil
}

func SaveCache(c CacheData) error {
	file, err := cacheFilePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, b, 0o644)
}

func cacheFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".liun-dots", "cache", "update-check.json"), nil
}

func NormalizeVersion(v string) string {
	v = strings.TrimSpace(v)
	v = strings.TrimPrefix(v, "v")
	return v
}

func IsNewer(remote, current string) bool {
	remoteParts, okRemote := parseSemver(remote)
	currentParts, okCurrent := parseSemver(current)
	if !okRemote || !okCurrent {
		return false
	}
	for i := range remoteParts {
		if remoteParts[i] > currentParts[i] {
			return true
		}
		if remoteParts[i] < currentParts[i] {
			return false
		}
	}
	return false
}

func parseSemver(v string) ([3]int, bool) {
	var out [3]int
	parts := strings.Split(NormalizeVersion(v), ".")
	if len(parts) != 3 {
		return out, false
	}
	for i, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			return out, false
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return out, false
		}
		out[i] = n
	}
	return out, true
}
