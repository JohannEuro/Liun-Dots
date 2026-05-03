package config

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestExpandReplacesHomeAndCleansPath(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)

	got, err := Expand("$HOME\\folder\\..\\config\\profile.ps1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := filepath.Join(home, "config", "profile.ps1")
	if got != want {
		t.Fatalf("got=%q want=%q", got, want)
	}
}

func TestExpandReplacesLocalAppDataVariants(t *testing.T) {
	home := t.TempDir()
	local := filepath.Join(home, "AppData", "Local")
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("LOCALAPPDATA", local)

	for _, in := range []string{"$LOCALAPPDATA\\nvim\\init.lua", "$env:LOCALAPPDATA\\nvim\\init.lua"} {
		t.Run(in, func(t *testing.T) {
			got, err := Expand(in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			want := filepath.Join(local, "nvim", "init.lua")
			if got != want {
				t.Fatalf("got=%q want=%q", got, want)
			}
		})
	}
}

func TestExpandErrorsWhenVariableIsUnresolved(t *testing.T) {
	t.Setenv("LOCALAPPDATA", "")

	_, err := Expand("$UNKNOWN\\x")
	if err == nil {
		t.Fatal("expected unresolved variable error")
	}
	if !strings.Contains(err.Error(), "unresolved variable") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDefaultTargetsReturnsExpandedDestinations(t *testing.T) {
	home := t.TempDir()
	local := filepath.Join(home, "AppData", "Local")
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("LOCALAPPDATA", local)

	targets, err := DefaultTargets()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(targets) != 3 {
		t.Fatalf("targets=%d want=3", len(targets))
	}

	for _, target := range targets {
		if strings.Contains(target.Destination, "$") {
			t.Fatalf("destination still has unresolved variable: %q", target.Destination)
		}
	}
}
