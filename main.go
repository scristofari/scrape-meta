package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var url = flag.String("url", "https://www.google.fr/", "Specify the url")

type meta struct {
	Title, Description, Image string
}

func main() {
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	m := &meta{}

	title, ok := scrape.Find(root, scrape.ByTag(atom.Title))
	if ok {
		m.Title = scrape.Text(title)
	}

	metas := scrape.FindAll(root, scrape.ByTag(atom.Meta))
	for _, meta := range metas {
		if scrape.Attr(meta, "name") == "description" {
			m.Description = scrape.Attr(meta, "content")
		}
		if scrape.Attr(meta, "property") == "og:image" {
			m.Image = scrape.Attr(meta, "content")
		}
	}

	fmt.Printf("Title : %s\n", m.Title)
	fmt.Printf("Description : %s\n", m.Description)
	fmt.Printf("Image : %s\n", m.Image)
}
