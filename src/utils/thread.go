package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractIssueNumberFromThreadTitle(title string) (int, error) {
	sepIndex := strings.Index(title, ")")
	if sepIndex == -1 {
		return 0, fmt.Errorf("invalid format: missing closing parenthesis")
	}
	beforeParenthesis := title[:sepIndex]
	beforeParenthesis = strings.TrimSpace(beforeParenthesis)
	x, err := strconv.Atoi(beforeParenthesis)
	if err != nil {
		return 0, fmt.Errorf("invalid format: %v", err)
	}

	return x, nil
}
