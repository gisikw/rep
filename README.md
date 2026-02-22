# crane

A thin CLI dispatcher for headless AI agent invocations. Crane normalizes the interface across agent providers (Claude Code, Cursor, OpenCode, etc.), picks the best available one based on a preference list, and gets out of the way.

## Usage

```bash
crane "refactor the auth module to use JWT"
echo "fix the failing test in user_test.go" | crane
crane --provider claude --model opus "review this PR"
crane --provider opencode --model "ollama/qwen3:32b" "fix the build"
crane --dir /path/to/project "run the tests and fix failures"
```

## Configuration

Crane reads `~/.config/crane/config.toml`:

```toml
# Provider preference order — first available wins
providers = ["claude", "cursor", "opencode"]

# Per-provider defaults
[claude]
model = "opus"
allow_all = true

[cursor]
allow_all = true

[opencode]
model = "ollama/qwen3-coder-next:latest"
```

## Provider Support

| Provider | Binary | Notes |
|----------|--------|-------|
| Claude Code | `claude` | Subscription-compliant invocation via official CLI |
| Cursor | `cursor-agent` / `agent` | Subscription-compliant invocation via official CLI |
| OpenCode | `opencode` | Local/remote models via Ollama, OpenRouter, etc. |

OpenCode handles tool calls and agentic workflows for models that don't have their own CLI harness. Configure the model endpoint (e.g. a remote Ollama server) in OpenCode's own config at `~/.config/opencode/config.json`.

## Install

```bash
just install   # builds and symlinks to ~/.local/bin
```

Or build manually:

```bash
just build     # produces ./crane
```
