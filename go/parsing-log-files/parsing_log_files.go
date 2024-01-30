package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	a := regexp.MustCompile(`<[(~*=\-)>]+`)
	return a.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`".*(?is)password.*"`)
	for _, v := range lines {
		if re.FindStringSubmatch(v) != nil {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User\s+)+(\w*)`)

	for i, v := range lines {
		if match := re.FindStringSubmatch(v); match != nil {
			username := match[2]
			lines[i] = "[USR] " + username + " " + v
		}
	}

	return lines
}
