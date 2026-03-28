package cmd

import (
	"strings"
	"testing"
)

func TestPromptingCommandOutputsGuide(t *testing.T) {
	restore, read := captureStdout(t)
	defer restore()

	rootCmd.SetArgs([]string{"prompting"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("prompting command failed: %v", err)
	}

	out := read()
	if !strings.Contains(out, "# maclabs prompting guide") {
		t.Fatalf("unexpected output: %q", out)
	}
}
