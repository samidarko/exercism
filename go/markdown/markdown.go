package markdown

// implementation to refactor

import (
	"fmt"
	"regexp"
	"strings"
)

var rules = [][2]string{
	//header rules
	//{`#{7}\s?([^\n]+)`, "<h7>$1</h7>"},
	//{`#{6}\s?([^\n]+)`, "<h6>$1</h6>"},
	//{`#{5}\s?([^\n]+)`, "<h5>$1</h5>"},
	//{`#{4}\s?([^\n]+)`, "<h4>$1</h4>"},
	//{`#{3}\s?([^\n]+)`, "<h3>$1</h3>"},
	//{`#{2}\s?([^\n]+)`, "<h2>$1</h2>"},
	//{`#{1}\s?([^\n]+)`, "<h1>$1</h1>"},
	//bold, italics and paragraph rules
	{`__([^_]+)__`, "<strong>$1</strong>"},
	{`_([^_]+)_`, "<em>$1</em>"},
	//{`([^\n]+\n?)`, "<p>$1</p>"},
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
	for _, line := range strings.Split(markdown, "\n") {
		switch line[0] {
		case '*':
			element := line[2:]
			list = append(list, element)
		case '#':
			if len(list) > 0 {
				html.WriteString(createList(list))
				list = make([]string, 0)
			}
			headerSize := len(regexp.MustCompile(`#+`).FindString(line))
			if headerSize < 7 {
				html.WriteString(fmt.Sprintf("<h%d>%s</h%d>", headerSize, line[headerSize+1:], headerSize))
			} else {
				html.WriteString(fmt.Sprintf("<p>%s</p>", line))
			}
		default:
			if len(list) > 0 {
				html.WriteString(createList(list))
				list = make([]string, 0)
			}
			line = applyRules(line)
			html.WriteString(fmt.Sprintf("<p>%s</p>", line))
		}
	}

	if len(list) > 0 {
		html.WriteString(createList(list))
	}

	return html.String()

}
