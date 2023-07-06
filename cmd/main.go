package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var List []Link

func main() {
	testCase := 4
	for i := 1; i <= testCase; i++ {
		filePath := fmt.Sprintf("./Testcase/ex%v.html", i)
		fmt.Println(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		r := io.Reader(file)

		// parse the reader
		doc, _ := html.Parse(r)

		// check for the a tag
		findLink(doc)

		fmt.Printf("output for test case %v \n", i)
		fmt.Println(List)

		// clear the list
		List = []Link{}
	}
}

// finds the "a" tag through dfs
func findLink(node *html.Node) {
	if node.Data == "a" {
		href := strings.TrimSpace(node.Attr[0].Val)
		text := strings.TrimSpace(findText(node))
		List = append(List, Link{Href: href, Text: text})
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
