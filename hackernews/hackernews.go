package hackernews

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

type Hackernews struct {
	pw *playwright.Playwright
	br playwright.Browser
}

func New() (*Hackernews, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("could not start playwright: %v", err)
	}

	br, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})

	if err != nil {
		return nil, fmt.Errorf("could not launch browser: %v", err)
	}

	return &Hackernews{
		pw: pw,
		br: br,
	}, nil
}

func (n *Hackernews) Close() error {
	if err := n.pw.Stop(); err != nil {
		return fmt.Errorf("could not stop playwright: %v", err)
	}

	if err := n.br.Close(); err != nil {
		return fmt.Errorf("could not close browser: %v", err)
	}

	return nil
}

func (n *Hackernews) Title() ([]string, error) {
	page, err := n.br.NewPage()
	if err != nil {
		return []string{}, fmt.Errorf("could not create page: %v", err)
	}

	if _, err := page.Goto("https://news.ycombinator.com"); err != nil {
		return []string{}, fmt.Errorf("could not goto: %v", err)
	}

	entries, err := page.QuerySelectorAll(".athing") // class name
	if err != nil {
		return []string{}, fmt.Errorf("could not get entries: %v", err)
	}

	out := make([]string, 0)
	for _, e := range entries {
		a, err := e.QuerySelector("td.title > span > a")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}

		title, err := a.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}

		out = append(out, title)
	}

	return out, nil
}
