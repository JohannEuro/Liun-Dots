package tools

import "testing"

func TestProbesFor(t *testing.T) {
	tests := []struct {
		name string
		bin  string
		want int
	}{
		{name: "windows terminal avoids probes", bin: "wt", want: 0},
		{name: "git uses one safe probe", bin: "git", want: 1},
		{name: "generic uses safe probes", bin: "nvim", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := probesFor(tt.bin)
			if len(got) != tt.want {
				t.Fatalf("got %d probes, want %d", len(got), tt.want)
			}
		})
	}
}
