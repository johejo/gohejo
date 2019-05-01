package logutils

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("failed to os.Pipe(): %s", err.Error())
	}

	stdout := os.Stdout
	os.Stdout = w

	logger := New()
	logger.Print("hello")

	os.Stdout = stdout
	defer func() {
		if err := w.Close(); err != nil {
			t.Errorf("failed to close writer: %s", err.Error())
		}
	}()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Errorf("failed to copy output buffer: %s", err.Error())
	}

	actual := buf.String()
	message := "hello"
	if strings.HasSuffix(actual, message) {
		t.Errorf("failed to output to stdout: actual=%s", actual)
	}
}
