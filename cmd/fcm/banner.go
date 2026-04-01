package main

import (
	"flag"
	"fmt"
	"os"
)

func supportsANSI() bool {
	term := os.Getenv("TERM")
	return term != "" && term != "dumb"
}

func printGradientText(url string) {
	reset := "\033[0m"

	runes := []rune(url)
	n := len(runes)
	if n == 0 {
		fmt.Println()
		return
	}

	mid := n / 2

	// Левая часть: cyan -> blue
	leftStart := [3]int{0, 255, 255}
	leftEnd := [3]int{0, 180, 255}

	// Правая часть: yellow -> orange
	rightStart := [3]int{255, 220, 0}
	rightEnd := [3]int{255, 140, 0}

	lerp := func(a, b, i, total int) int {
		if total <= 1 {
			return a
		}
		t := float64(i) / float64(total-1)
		return int(float64(a) + t*float64(b-a))
	}

	for i, ch := range runes {
		var r, g, b int

		if i < mid {
			r = lerp(leftStart[0], leftEnd[0], i, mid)
			g = lerp(leftStart[1], leftEnd[1], i, mid)
			b = lerp(leftStart[2], leftEnd[2], i, mid)
		} else {
			j := i - mid
			rightLen := n - mid
			r = lerp(rightStart[0], rightEnd[0], j, rightLen)
			g = lerp(rightStart[1], rightEnd[1], j, rightLen)
			b = lerp(rightStart[2], rightEnd[2], j, rightLen)
		}

		fmt.Printf("\033[38;2;%d;%d;%dm%c", r, g, b, ch)
	}

	fmt.Print(reset + "\n")
}

func printBanner() {
	const (
		reset  = "\033[0m"
		cyan   = "\033[38;5;51m"
		cyan2  = "\033[38;5;45m"
		yellow = "\033[38;5;220m"
		orange = "\033[38;5;208m"
	)

	if !supportsANSI() {
		fmt.Println("FCM CLI")
		fmt.Println("Firebase Cloud Messaging CLI")
		fmt.Println()
		return
	}

	fmt.Printf("%s███████  ██████ ███    ███%s", cyan, reset)
	fmt.Printf("      ")
	fmt.Printf("%s██████ ██      ██%s\n", orange, reset)

	fmt.Printf("%s██      ██      ████  ████%s", cyan, reset)
	fmt.Printf("     ")
	fmt.Printf("%s██      ██      ██%s\n", orange, reset)

	fmt.Printf("%s█████   ██      ██ ████ ██%s", cyan2, reset)
	fmt.Printf("     ")
	fmt.Printf("%s██      ██      ██%s\n", yellow, reset)

	fmt.Printf("%s██      ██      ██  ██  ██%s", cyan2, reset)
	fmt.Printf("     ")
	fmt.Printf("%s██      ██      ██%s\n", yellow, reset)

	fmt.Printf("%s██       ██████ ██      ██%s", cyan, reset)
	fmt.Printf("      ")
	fmt.Printf("%s██████ ███████ ██%s\n", orange, reset)

	fmt.Println()
	printGradientText("https://github.com/interdev7/fcm-cli")
	fmt.Println()
	fmt.Printf("%sFirebase Cloud Messaging CLI%s\n", cyan2, reset)
	fmt.Printf("%sFast • Simple • Scriptable%s\n\n", yellow, reset)
}

func setupUsage() {
	flag.Usage = func() {
		printBanner()

		fmt.Println("Usage: fcm [options]")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("  init                         Generate default fcm.yaml")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("  -k, --key <file>             Firebase key file")
		fmt.Println("  -t, --token <token>          Single FCM token")
		fmt.Println("  --tokens <t1,t2,t3>          Comma-separated token list")
		fmt.Println("  --tokens-file <file>         File with one token per line")
		fmt.Println("  -n, --notification <json>    Notification JSON")
		fmt.Println("  -d, --data <json>            Data JSON")
		fmt.Println("  -topic <topic>               Topic")
		fmt.Println("  -c, --condition <expr>       Condition")
		fmt.Println("  -f, --config <file>          YAML config file")
		fmt.Println("  --profile <name>             Profile inside config")
		fmt.Println("  --env-file <file>            Load additional .env file")
		fmt.Println("  -l, --log <level>            Log level: info|debug|json")
		fmt.Println("  --json                       Machine-readable output")
		fmt.Println("  -v, --version                Version")
		fmt.Println("  -h, --help                   Help")
	}
}
