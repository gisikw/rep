package main

import "testing"

func TestParseArgs_PromptFromPositional(t *testing.T) {
	opts, err := ParseArgs([]string{"do", "the", "thing"}, Config{})
	if err != nil {
		t.Fatal(err)
	}
	if opts.Prompt != "do the thing" {
		t.Errorf("prompt = %q, want %q", opts.Prompt, "do the thing")
	}
}

func TestParseArgs_Flags(t *testing.T) {
	opts, err := ParseArgs([]string{
		"--provider", "cursor",
		"--model", "gpt-4",
		"--dir", "/tmp/test",
		"--system-prompt", "be helpful",
		"do stuff",
	}, Config{})
	if err != nil {
		t.Fatal(err)
	}
	if opts.Provider != "cursor" {
		t.Errorf("provider = %q", opts.Provider)
	}
	if opts.Model != "gpt-4" {
		t.Errorf("model = %q", opts.Model)
	}
	if opts.Dir != "/tmp/test" {
		t.Errorf("dir = %q", opts.Dir)
	}
	if opts.SystemPrompt != "be helpful" {
		t.Errorf("system-prompt = %q", opts.SystemPrompt)
	}
	if opts.Prompt != "do stuff" {
		t.Errorf("prompt = %q", opts.Prompt)
	}
}

func TestParseArgs_AllowAllDefault(t *testing.T) {
	opts, err := ParseArgs([]string{"hello"}, Config{})
	if err != nil {
		t.Fatal(err)
	}
	if !opts.AllowAll {
		t.Error("expected AllowAll to default to true")
	}
}

func TestParseArgs_WithPermissions(t *testing.T) {
	opts, err := ParseArgs([]string{"--with-permissions", "hello"}, Config{})
	if err != nil {
		t.Fatal(err)
	}
	if opts.AllowAll {
		t.Error("expected AllowAll to be false with --with-permissions")
	}
}

func TestParseArgs_UnknownFlag(t *testing.T) {
	_, err := ParseArgs([]string{"--bogus"}, Config{})
	if err == nil {
		t.Error("expected error for unknown flag")
	}
}

func TestParseArgs_EmptyPrompt(t *testing.T) {
	opts, err := ParseArgs([]string{}, Config{})
	if err != nil {
		t.Fatal(err)
	}
	if opts.Prompt != "" {
		t.Errorf("expected empty prompt, got %q", opts.Prompt)
	}
}
