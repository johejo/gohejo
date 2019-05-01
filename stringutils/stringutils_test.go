package stringutils

import (
	"testing"
)

func TestStrip(t *testing.T) {
	t.Parallel()
	var (
		actual   string
		expected string
	)
	const jsonStr = `
{
  "key": "value"
}
`
	actual = Strip(jsonStr)
	expected = `{"key":"value"}`

	if actual != expected {
		t.Errorf("failed to strip json: actual=%s, expected=%s", actual, expected)
	}
}
