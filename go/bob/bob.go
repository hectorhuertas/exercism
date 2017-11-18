package bob

import (
	"regexp"
	"strings"
)

var hasLetters = regexp.MustCompile("[a-zA-Z]+")

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	isShouting := remark == strings.ToUpper(remark) && hasLetters.MatchString(remark)
	if isShouting {
		return "Whoa, chill out!"
	}

	isQuestion := strings.HasSuffix(remark, "?")
	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
