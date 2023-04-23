package spam

import (
	"strings"
)

type SpamChecker struct{}

func (s *SpamChecker) CheckText(newText, oldText string) bool {
	counter := 0.0
	check := make([]string, 0, len(oldText))
	if len(oldText) < len(newText) {
		return false
	}
	for _, v := range strings.Split(oldText, " ") {
		check = append(check, strings.ToLower(v))
	}
	for i, v := range strings.Split(newText, " ") {
		if check[i] == strings.ToLower(v) {
			counter += 1.0
		}
	}
	result := float64(counter / float64(len(check)) * 100.0)
	if result > 50 {
		return true
	}
	return false
}
