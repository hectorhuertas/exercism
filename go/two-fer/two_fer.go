package twofer

import "fmt"

func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", s)
}
