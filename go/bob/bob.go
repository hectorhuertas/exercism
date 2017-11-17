package bob

import "strings"
import "regexp"
import "unicode"

func removeWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func Hey(remark string) string {
	remark = removeWhitespace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}

	reg, _ := regexp.Compile("[^a-zA-Z]+")
	r := reg.ReplaceAllString(remark, "")
	isShouting := r == strings.ToUpper(r) && r != ""
	if isShouting {
		return "Whoa, chill out!"
	}

	isQuestion := remark[len(remark)-1:] == "?"
	if isQuestion {
		return "Sure."
	}
	return "Whatever."
}
