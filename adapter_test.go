package main

import "testing"

func TestClaudeAdapter_Build(t *testing.T) {
	a := ClaudeAdapter{}
	cmd := a.Build(InvokeRequest{
		Prompt:       "hello",
		Model:        "opus",
		SystemPrompt: "be nice",
		Dir:          "/tmp",
		AllowAll:     true,
	})

	if cmd.Path == "" {
		t.Fatal("expected non-empty path")
	}
	args := cmd.Args[1:] // skip binary name
	assertContains(t, args, "--dangerously-skip-permissions")
	assertContains(t, args, "--model")
	assertContains(t, args, "opus")
	assertContains(t, args, "--append-system-prompt")
	assertContains(t, args, "be nice")
	if cmd.Dir != "/tmp" {
		t.Errorf("dir = %q, want /tmp", cmd.Dir)
	}
}

func TestClaudeAdapter_NoPermissions(t *testing.T) {
	a := ClaudeAdapter{}
	cmd := a.Build(InvokeRequest{
		Prompt:   "hello",
		AllowAll: false,
	})
	args := cmd.Args[1:]
	for _, arg := range args {
		if arg == "--dangerously-skip-permissions" {
			t.Error("should not include --dangerously-skip-permissions when AllowAll is false")
		}
	}
}

func TestCursorAdapter_InlinesSystemPrompt(t *testing.T) {
	a := CursorAdapter{}
	cmd := a.Build(InvokeRequest{
		Prompt:       "do stuff",
		SystemPrompt: "context here",
	})
	args := cmd.Args[1:]
	// System prompt should be inlined into -p value, not a separate flag
	for _, arg := range args {
		if arg == "--append-system-prompt" {
			t.Error("cursor should not use --append-system-prompt")
		}
	}
	// The -p value should contain both system prompt and prompt
	if len(args) > 1 && args[0] == "-p" {
		if args[1] != "context here\n\ndo stuff" {
			t.Errorf("prompt = %q, want system prompt inlined", args[1])
		}
	}
}

func TestGetAdapter_Defaults(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"claude", "main.ClaudeAdapter"},
		{"cursor", "main.CursorAdapter"},
		{"unknown", "main.GenericAdapter"},
	}
	for _, tt := range tests {
		a := GetAdapter(tt.name)
		got := typeName(a)
		if got != tt.expected {
			t.Errorf("GetAdapter(%q) = %s, want %s", tt.name, got, tt.expected)
		}
	}
}

func TestExitCode(t *testing.T) {
	code := exitCode(nil)
	if code != 1 {
		t.Errorf("exitCode(nil) = %d, want 1", code)
	}
}

func assertContains(t *testing.T, args []string, val string) {
	t.Helper()
	for _, a := range args {
		if a == val {
			return
		}
	}
	t.Errorf("args %v missing %q", args, val)
}

func typeName(v interface{}) string {
	switch v.(type) {
	case ClaudeAdapter:
		return "main.ClaudeAdapter"
	case CursorAdapter:
		return "main.CursorAdapter"
	case GenericAdapter:
		return "main.GenericAdapter"
	default:
		return "unknown"
	}
}
