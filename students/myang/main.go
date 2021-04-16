package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func main() {
	filenames := []string{"./ex1.html", "./ex2.html", "./ex3.html", "./ex4.html"}
	for _, filename := range filenames {
		tokenizer, err := getHtmlTokenizer(filename)
		if err != nil {
		panic(err)
		}

		links := getLinks(tokenizer)
		fmt.Println(filename, links)
	}
}

func getHtmlTokenizer(filename string) (*html.Tokenizer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	tokenizer := html.NewTokenizer(file)

	return tokenizer, nil
}

func getLinks(tokenizer *html.Tokenizer) []Link {
	var links []Link
	var link Link
	depth := 0
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links
		case html.TextToken:
			if depth > 0 {
				if link.Text != "" {
					link.Text += " "
				}
				text := strings.TrimSpace(string(tokenizer.Text()))
				link.Text += text
			}
		case html.StartTagToken, html.EndTagToken:
			token, _ := tokenizer.TagName()

			if len(token) == 1 && token[0] == 'a' {
				if tokenType == html.StartTagToken {
					key, value, _ := tokenizer.TagAttr()
					if string(key) == "href" {
						link.Href = string(value)
					}
					depth++
				} else {
					depth--
					links = append(links, link)
					link = Link{}
				}
			}
		}
	}
}