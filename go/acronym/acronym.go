package acronym

import "strings"

func Abbreviate(s string) string {
	s2 := strings.Replace(s, "-", " ", -1)
	words := strings.Split(s2, " ")

	var acronym string
	for _, word := range words {
		acronym += strings.ToUpper(word[0:1])
	}

	return acronym
}
