package hackernews_test

import (
	"fmt"
	"log"

	"github.com/itsubaki/playwright-go-training/hackernews"
)

func ExampleHackernews_Title() {
	y, err := hackernews.New()
	if err != nil {
		log.Fatalf("new: %v", err)
	}
	defer y.Close()

	title, err := y.Title()
	if err != nil {
		log.Fatalf("title: %v", err)
	}

	for i, t := range title {
		fmt.Printf("%d: %v\n", i+1, t)
	}

	// Output:
	// 1: UPS and the Package Wars
	// 2: Just: A Command Runner
	// 3: Interactive California Reservoir Levels Dashboard
	// 4: How to store your app's entire state in the url
	// 5: VanadiumOS: Portable, multi-user Unix-like OS
	// 6: 3D in CSS
	// 7: Fyrox Game Engine 0.29
	// 8: Tomu – A family of devices which fit inside your USB port
	// 9: Explanations: Play Don't Show
	// 10: Jump Servers
	// 11: High level mods/patches for console video games
	// 12: Sequence8 – a music sequencing toy in PICO-8
	// 13: Microsoft eyes $10B bet on ChatGPT
	// 14: Fake it until you automate it
	// 15: Binance Is Bleeding Assets, $12B Gone in Less Than 60 Days
	// 16: Banks are reducing their exposure to crypto
	// 17: Amber Monitors (2020)
	// 18: Dilution of expertise in the rise and fall of collective innovation
	// 19: Skribilo: Document Programming Framework
	// 20: Show HN: HyperLogLog in Zig
	// 21: Mapping Python to LLVM
	// 22: Extending the Android SDK
	// 23: Three lessons from Threema: Analysis of a secure messenger
	// 24: Making Sense of Hexdump (2008)
	// 25: Rio de Janeiro Botanical Garden (2018)
	// 26: Values and objects in programming languages (1982) [pdf]
	// 27: A college student made an app to detect AI-written text
	// 28: Handling 100K Consumers with One Pulsar Topic
	// 29: A Wicked Problem – The Second Egress: Building a Code Change
	// 30: Kable (YC W22) Is Hiring Lead Engineer (Remote/US)
}
