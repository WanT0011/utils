package strings

import (
	"strings"
)

func Join(splitStr string, items ...string) string {
	return strings.Join(items, splitStr)
}
