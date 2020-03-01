package utils

import "strings"

// StringsJoinWithBuilder join strings with strings.Builder
func StringsJoinWithBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
