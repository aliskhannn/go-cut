//go:build integration
// +build integration

package integration

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const binName = "gocut"

func buildBinary(t *testing.T) string {
	t.Helper()
	binPath := filepath.Join(t.TempDir(), binName)
	cmd := exec.Command("go", "build", "-o", binPath, "github.com/aliskhannn/go-cut/cmd/gocut") // путь к main
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build binary: %v\n%s", err, string(out))
	}
	return binPath
}

func runCmd(t *testing.T, bin string, stdin string, args ...string) string {
	t.Helper()
	cmd := exec.Command(bin, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if stdin != "" {
		cmd.Stdin = bytes.NewBufferString(stdin)
	}
	if err := cmd.Run(); err != nil {
		t.Fatalf("command failed: %v", err)
	}
	return out.String()
}

func TestCut_Integration(t *testing.T) {
	bin := buildBinary(t)

	tmpFile := filepath.Join(t.TempDir(), "testfile.txt")
	content := "a:b:c:d\ne:f:g:h\ni:j:k:l"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("basic fields", func(t *testing.T) {
		out := runCmd(t, bin, "", "-f", "1,3", "-d", ":", tmpFile)
		expected := "a:c\ne:g\ni:k\n"
		if out != expected {
			t.Fatalf("got %q, want %q", out, expected)
		}
	})

	t.Run("fields range", func(t *testing.T) {
		out := runCmd(t, bin, "", "-f", "2-4", "-d", ":", tmpFile)
		expected := "b:c:d\nf:g:h\nj:k:l\n"
		if out != expected {
			t.Fatalf("got %q, want %q", out, expected)
		}
	})

	t.Run("separated only", func(t *testing.T) {
		tmpFile2 := filepath.Join(t.TempDir(), "testfile2.txt")
		content2 := "x:y\nno_delimiter_line\np:q"
		if err := os.WriteFile(tmpFile2, []byte(content2), 0644); err != nil {
			t.Fatal(err)
		}
		out := runCmd(t, bin, "", "-f", "2", "-d", ":", "-s", tmpFile2)
		expected := "y\nq\n"
		if out != expected {
			t.Fatalf("got %q, want %q", out, expected)
		}
	})

	t.Run("stdin input", func(t *testing.T) {
		stdin := "1,2,3\n4,5,6"
		out := runCmd(t, bin, stdin, "-f", "2", "-d", ",")
		expected := "2\n5\n"
		if out != expected {
			t.Fatalf("got %q, want %q", out, expected)
		}
	})
}
