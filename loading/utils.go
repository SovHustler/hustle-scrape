package loading

import (
	"strings"

	"github.com/anaskhan96/soup"
	"golang.org/x/net/html"
)

func LoadPageRaw(url string) ([]string, error) {
	resp, err := soup.Get(url)
	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(resp)

	return getPageRaw(doc), nil
}

func getPageRaw(node soup.Root) []string {
	var result []string
	doGetPageRaw(node, &result)

	return result
}

func doGetPageRaw(node soup.Root, lines *[]string) {
	for _, c := range node.Children() {
		doGetPageRaw(c, lines)
	}

	text := getNodeText(node)
	if text == "" {
		return
	}

	*lines = append(*lines, strings.ToLower(text))
}

func getNodeText(node soup.Root) string {
	if node.Pointer.Type != html.TextNode {
		return ""
	}

	return strings.TrimSpace(strings.ReplaceAll(node.NodeValue, " ", " "))
}
