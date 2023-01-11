package google

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
)

type Google struct {
	pw *playwright.Playwright
	br playwright.Browser
}

func New() (*Google, error) {
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

	return &Google{
		pw: pw,
		br: br,
	}, nil
}

func (g *Google) Close() error {
	if err := g.pw.Stop(); err != nil {
		return fmt.Errorf("could not stop playwright: %v", err)
	}

	if err := g.br.Close(); err != nil {
		return fmt.Errorf("could not close browser: %v", err)
	}

	return nil
}

func (g *Google) Text(text string) (playwright.Page, error) {
	page, err := g.br.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %v", err)
	}

	if _, err := page.Goto("https://google.com"); err != nil {
		return nil, fmt.Errorf("could not goto: %v", err)
	}

	input, err := page.Locator("input")
	if err != nil {
		return nil, fmt.Errorf("could not locate: %v", err)
	}

	if err := input.Fill(text); err != nil {
		return nil, fmt.Errorf("fill: %v", err)
	}

	if err := input.Press("Enter"); err != nil {
		return nil, fmt.Errorf("enter: %v", err)

	}

	return page, nil
}

func (g *Google) Image(filename string) (playwright.Page, error) {
	page, err := g.br.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %v", err)
	}

	if _, err := page.Goto("https://google.com"); err != nil {
		return nil, fmt.Errorf("could not goto: %v", err)
	}

	if err := page.Click("img[alt='カメラ検索']"); err != nil {
		return nil, fmt.Errorf("could not click camera icon: %v", err)
	}

	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file(%v): %v", filename, err)
	}

	if err := page.SetInputFiles("input[type='file']", []playwright.InputFile{
		{
			Name:     filepath.Base(filename),
			MimeType: "image/png",
			Buffer:   buf,
		},
	}); err != nil {
		return nil, fmt.Errorf("could not set input files: %v", err)
	}

	if _, err := page.WaitForSelector("div:is(:text(\"この画像を検索\"))"); err != nil {
		return nil, fmt.Errorf("wait for selector: %v", err)
	}

	return page, nil
}
