package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config holds crane's configuration.
type Config struct {
	Providers []string                   `toml:"providers"`
	Provider  map[string]ProviderOptions `toml:"provider"`
}

// ProviderOptions holds per-provider defaults.
type ProviderOptions struct {
	Model    string `toml:"model"`
	AllowAll *bool  `toml:"allow_all"`
}

// ProviderConfig returns the config for a given provider, or an empty config.
func (c Config) ProviderConfig(name string) ProviderOptions {
	if c.Provider == nil {
		return ProviderOptions{}
	}
	return c.Provider[name]
}

// LoadConfig reads ~/.config/crane/config.toml. Returns defaults if missing.
func LoadConfig() (Config, error) {
	cfg := Config{
		Providers: []string{"claude"},
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, nil // Use defaults
	}

	path := filepath.Join(home, ".config", "crane", "config.toml")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil // Use defaults
		}
		return cfg, err
	}

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	if len(cfg.Providers) == 0 {
		cfg.Providers = []string{"claude"}
	}

	return cfg, nil
}
