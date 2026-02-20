package main

import (
	"fmt"
	"os/exec"
)

// ResolveProvider picks the first available provider from the preference list.
// If explicit is non-empty, it's used directly (no availability check).
func ResolveProvider(explicit string, preference []string) (string, error) {
	if explicit != "" {
		return explicit, nil
	}

	for _, name := range preference {
		if providerAvailable(name) {
			return name, nil
		}
	}

	return "", fmt.Errorf("no available provider (tried: %v)", preference)
}

// providerAvailable checks if the provider's binary is on PATH.
func providerAvailable(name string) bool {
	bins := providerBinaries(name)
	for _, bin := range bins {
		if _, err := exec.LookPath(bin); err == nil {
			return true
		}
	}
	return false
}

// providerBinaries returns the binary names to search for a given provider.
func providerBinaries(name string) []string {
	switch name {
	case "claude":
		return []string{"claude"}
	case "cursor":
		return []string{"cursor-agent", "agent"}
	case "ollama":
		return []string{"ollama"}
	default:
		return []string{name}
	}
}
