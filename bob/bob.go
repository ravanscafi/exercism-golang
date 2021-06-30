// Package bob simulates this lackadaisical teenager conversations.
package bob

import (
	"regexp"
	"strings"
)

// Hey responds like Bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	var r = regexp.MustCompile(`[^A-Za-z?]`)

	if r.ReplaceAllString(remark, "") == "?" {
		return "Sure."
	}

	if r.ReplaceAllString(remark, "") == "" {
		return "Whatever."
	}

	var isQuestion = remark[len(remark)-1] == '?'

	if remark == strings.ToUpper(remark) && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	if remark == strings.ToUpper(remark) {
		return "Whoa, chill out!"
	}

	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
