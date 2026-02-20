package main

import "testing"

func TestResolveProvider_Explicit(t *testing.T) {
	p, err := ResolveProvider("cursor", []string{"claude"})
	if err != nil {
		t.Fatal(err)
	}
	if p != "cursor" {
		t.Errorf("expected cursor, got %q", p)
	}
}

func TestResolveProvider_NoProviders(t *testing.T) {
	_, err := ResolveProvider("", []string{"nonexistent-binary-xyz"})
	if err == nil {
		t.Error("expected error when no provider is available")
	}
}

func TestProviderBinaries(t *testing.T) {
	tests := []struct {
		name     string
		expected []string
	}{
		{"claude", []string{"claude"}},
		{"cursor", []string{"cursor-agent", "agent"}},
		{"ollama", []string{"ollama"}},
		{"custom", []string{"custom"}},
	}
	for _, tt := range tests {
		bins := providerBinaries(tt.name)
		if len(bins) != len(tt.expected) {
			t.Errorf("%s: got %v, want %v", tt.name, bins, tt.expected)
			continue
		}
		for i := range bins {
			if bins[i] != tt.expected[i] {
				t.Errorf("%s[%d]: got %q, want %q", tt.name, i, bins[i], tt.expected[i])
			}
		}
	}
}
