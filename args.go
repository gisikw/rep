package main

import (
	"fmt"
	"strings"
)

// Options holds the parsed command-line options.
type Options struct {
	Prompt       string
	Provider     string // Empty means "use config preference order"
	Model        string
	Dir          string
	SystemPrompt string
	AllowAll     bool
}

// ParseArgs parses CLI arguments into Options.
func ParseArgs(args []string, cfg Config) (Options, error) {
	opts := Options{
		AllowAll: true, // Default: skip permissions
	}

	var positional []string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--provider":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("--provider requires a value")
			}
			i++
			opts.Provider = args[i]
		case "--model":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("--model requires a value")
			}
			i++
			opts.Model = args[i]
		case "--dir":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("--dir requires a value")
			}
			i++
			opts.Dir = args[i]
		case "--system-prompt":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("--system-prompt requires a value")
			}
			i++
			opts.SystemPrompt = args[i]
		case "--no-permissions":
			opts.AllowAll = true
		case "--with-permissions":
			opts.AllowAll = false
		case "--help", "-h":
			return opts, fmt.Errorf("help requested")
		default:
			if strings.HasPrefix(args[i], "--") {
				return opts, fmt.Errorf("unknown flag: %s", args[i])
			}
			positional = append(positional, args[i])
		}
	}

	opts.Prompt = strings.Join(positional, " ")
	return opts, nil
}
