package micro

import (
	"os"
	"strings"
)

func SubtractString(original string, substring string) string {
	if strings.HasSuffix(original, substring) {
		return original[:len(original)-len(substring)]
	}
	return original
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
