package book

import (
	"fmt"
	"regexp"
	"strings"
)

var idIndex int = -1

func nextId() string {
	idIndex++
	return fmt.Sprintf("afield%d", idIndex)
}

var rxText = regexp.MustCompile("{{text}}")
var rxWord = regexp.MustCompile("{{word}}")

type placeholder struct {
	pattern     string
	replacement string
}

var placeholders = []placeholder{
	{
		"{{text}}",
		"<textarea id=\"%d\"></textarea>",
	},
	{
		"{{word}}",
		"<input type=\"text\" id=\"%d\">",
	},
	{
		"{{picture}}",
		"<input type=\"file\" id=\"%d\">",
	},
}

func replaceByRegex(text string, rx regexp.Regexp, ns string) string {
	start := strings.Index(text, "{{")
	if start == -1 {
		return text
	}

	return ""
}

func preprocessLine(line string) string {
	return ""
}
