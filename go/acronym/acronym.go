package acronym

import "strings"

func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)

	words := strings.Fields(s)

	var acronym string
	for _, w := range words {
		acronym += strings.ToUpper(w[0:1])
	}

	return acronym
}
