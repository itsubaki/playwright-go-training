package google_test

import (
	"fmt"
	"log"
	"net/url"

	"github.com/itsubaki/playwright-go-training/google"
)

func ExampleGoogle_Text() {
	g, err := google.New()
	if err != nil {
		log.Fatalf("new: %v", err)
	}
	defer g.Close()

	page, err := g.Text("foobar")
	if err != nil {
		log.Fatalf("text search: %v", err)
	}

	pageURL := page.URL()
	url, err := url.Parse(pageURL)
	if err != nil {
		log.Fatalf("parse %v: %v", pageURL, err)
	}

	fmt.Printf("%v%v\n", url.Host, url.Path)

	// Output:
	// www.google.com/search

}

// go test -timeout 30s -tags -v -run ^ExampleGoogle_Image$ github.com/itsubaki/playwright-go-training/google
func ExampleGoogle_Image() {
	g, err := google.New()
	if err != nil {
		log.Fatalf("new: %v", err)
	}
	defer g.Close()

	page, err := g.Image("../kotowaza_neko_koban.png")
	if err != nil {
		log.Fatalf("image search: %v", err)
	}

	pageURL := page.URL()
	url, err := url.Parse(pageURL)
	if err != nil {
		log.Fatalf("parse %v: %v", pageURL, err)
	}
	fmt.Println(url.Host)

	// Output:
	// lens.google.com

}
