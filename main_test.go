package main_test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
)

func Example() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	defer func() {
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
	}()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer func() {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
	}()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err := page.Goto("https://news.ycombinator.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	entries, err := page.QuerySelectorAll(".athing")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}

	for i, entry := range entries {
		titleElement, err := entry.QuerySelector("td.title > span > a")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}

		title, err := titleElement.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}

		fmt.Printf("%d: %s\n", i+1, title)
	}

	// Output:
	// 1: Microsoft is preparing to add ChatGPT to Bing
	// 2: Meta Quest Pro – Bad AR Passthrough
	// 3: French startup unveils new residential thermo-acoustic heat pump
	// 4: The faker's guide to reading x86 assembly language
	// 5: Infinite AI Array
	// 6: Scientists have discovered the first virovore – an organism that eats viruses
	// 7: Thanks to DALL-E, the race to make artificial protein drugs is on
	// 8: Google wants RISC-V to be a “tier-1” Android architecture
	// 9: Show HN: A device that only lets you type lol if you've truly laughed out loud
	// 10: Names that decreased in popularity from 2020 to 2021
	// 11: Reviving proteins from billions of years ago to fight diseases in human cells
	// 12: Breaking RSA with a quantum computer?
	// 13: The Beauty of Bézier Curves (2021) [video]
	// 14: Japan’s business owners can’t find successors – one man is giving his away
	// 15: A brief rant on converging compliance regimes
	// 16: “Unexplainable” core dump (2011)
	// 17: Even More Bay Area House Party
	// 18: South Korea’s online security dead end
	// 19: Migrating from AWS to Fly.io
	// 20: Google researcher, long out of math, cracks devilish problem about sets
	// 21: Web hackers vs. the auto industry
	// 22: Hnefatafl
	// 23: The MOS 6502 is (mostly) Turing-complete without registers
	// 24: Measuring an engineering organization
	// 25: Underappreciated challenges with Python packaging
	// 26: Interesting things about the Lua interpreter (2020)
	// 27: Show HN: Ov – feature rich terminal pager
	// 28: Modules, not microservices
	// 29: Frederic Tudor in 1806 brings cocktails and ice cream to the rest of the world
	// 30: AT&T's predictions of the future (1993)
}

func Example_fileupload() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	defer func() {
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
	}()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer func() {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
	}()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err := page.Goto("https://google.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	if err := page.Click("img[alt='カメラ検索']"); err != nil {
		log.Fatalf("could not click camera icon: %v", err)
	}

	filename := "kotowaza_neko_koban.png"
	buf, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("read file(%v): %v", filename, err)
	}

	if err := page.SetInputFiles("input[type='file']", []playwright.InputFile{
		{
			Name:     filepath.Base(filename),
			MimeType: "image/png",
			Buffer:   buf,
		},
	}); err != nil {
		log.Fatalf("could not set input files: %v", err)
	}
	if _, err := page.WaitForSelector("div:is(:text(\"この画像を検索\"))"); err != nil {
		log.Fatalf("wait for selector: %v", err)
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
