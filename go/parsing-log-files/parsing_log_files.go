package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	matched, _ := regexp.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]\s.+`, text)
	return matched
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		count += len(re.FindAllString(line, -1))
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (output []string) {
	re := regexp.MustCompile(`User\s+(?P<User>\w+)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			userIndex := re.SubexpIndex("User")
			tag := fmt.Sprintf("[USR] %s ", matches[userIndex])
			line = fmt.Sprint(tag, line)
		}
		output = append(output, line)
	}

	return
}
