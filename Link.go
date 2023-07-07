package Link

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var list []Link

func GetLinks(node *html.Node) []Link {
	list = []Link{}
	findLink(node)
	return list
}

// finds the "a" tag through dfs
func findLink(node *html.Node) {
	if node.Data == "a" {
		href := strings.TrimSpace(node.Attr[0].Val)
		text := strings.TrimSpace(findText(node))
		list = append(list, Link{Href: href, Text: text})
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		findLink(child)
	}
}

// returns the text string in "a"
func findText(node *html.Node) string {
	if node == nil {
		return ""
	}

	if node.FirstChild == nil {
		if node.Type == html.TextNode {
			return node.Data
		}
		return ""
	}

	text := ""
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		text = text + " " + strings.TrimSpace(findText(child))
	}

	return text
}
