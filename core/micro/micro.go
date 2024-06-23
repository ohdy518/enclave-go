package micro

import "strings"

func SubtractString(original string, substring string) string {
	if strings.HasSuffix(original, substring) {
		return original[:len(original)-len(substring)]
	}
	return original
}
