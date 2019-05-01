package stringutils

import "strings"

// Strip removes "\n", "\r", " ".
func Strip(s string) string {
	return strings.NewReplacer("\n", "", "\r", "", " ", "").Replace(s)
}
