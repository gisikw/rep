# crane

A thin CLI dispatcher for headless AI agent invocations. Crane normalizes the interface across agent providers (Claude Code, Cursor, Ollama, etc.), picks the best available one based on a preference list, and gets out of the way.

## Usage

```bash
crane "refactor the auth module to use JWT"
echo "fix the failing test in user_test.go" | crane
crane --provider claude --model opus "review this PR"
crane --dir /path/to/project "run the tests and fix failures"
```

## Configuration

Crane reads `~/.config/crane/config.toml`:

```toml
# Provider preference order — first available wins
providers = ["cursor", "claude", "ollama"]

# Per-provider defaults
[claude]
model = "opus"
allow_all = true

[cursor]
allow_all = true

[ollama]
model = "qwen2.5-coder:32b"
```

## Provider Support

| Provider | Binary | Status |
|----------|--------|--------|
| Claude Code | `claude` | Supported |
| Cursor | `cursor-agent` / `agent` | Supported |
| Ollama | TBD | Planned |

## Install

```bash
go install github.com/gisikw/crane@latest
```

Or with Nix:

```bash
nix build
```
