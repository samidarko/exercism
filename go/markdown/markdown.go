package markdown

// implementation to refactor

import (
	"fmt"
	"regexp"
	"strings"
)

var rules = [][2]string{
	//bold, italics and paragraph rules
	{`__([^_]+)__`, "<strong>$1</strong>"},
	{`_([^_]+)_`, "<em>$1</em>"},
}

func applyRules(line string) string {
	for _, rule := range rules {
		line = regexp.MustCompile(rule[0]).ReplaceAllString(line, rule[1])
	}
	return line
}

func createList(list []string) string {
	var output strings.Builder
	output.WriteString("<ul>")
	for _, element := range list {
		output.WriteString(fmt.Sprintf("<li>%s</li>", element))
	}
	output.WriteString("</ul>")
	return applyRules(output.String())
}

// Render translates markdown to HTML
func Render(markdown string) string {

	var html strings.Builder
	list := make([]string, 0)
	lines := strings.Split(markdown, "\n")
	for i, line := range lines {

		if line[0] == '*' {
			// if line is a list item, we append to the list
			element := line[2:]
			list = append(list, element)

			// if a list item was the last line
			if i == len(lines)-1 {
				html.WriteString(createList(list))
			}
			continue
		}

		if len(list) > 0 {
			// this line is not a list item	anymore, we can process the list
			html.WriteString(createList(list))
			list = make([]string, 0)
		}

		switch line[0] {
		case '#':
			headerSize := len(regexp.MustCompile(`#+`).FindString(line))
			if headerSize < 7 {
				html.WriteString(fmt.Sprintf("<h%d>%s</h%d>", headerSize, line[headerSize+1:], headerSize))
			} else {
				html.WriteString(fmt.Sprintf("<p>%s</p>", line))
			}
		default:
			line = applyRules(line)
			html.WriteString(fmt.Sprintf("<p>%s</p>", line))
		}
	}

	return html.String()

}
