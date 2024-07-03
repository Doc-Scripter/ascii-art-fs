package main

import (
	"fmt"
	"os"

	asciiArt "ascii/functions"
)

func main() {
	args := len(os.Args)
	if args < 2 || args > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	// if only the string to be printed is provided
	if args == 2 {
		s, err := os.ReadFile("resources/standard.txt")
		if err != nil {
			fmt.Println("File not found")
			return
		}
		m := asciiArt.AsciiArt(string(s))
		res := asciiArt.Tab(os.Args[1])
		asciiArt.Paragraph(res, m)
	}

	// if the string to be printed is provided and also the bannerfile
	if args == 3 {
		file := os.Args[2]

		if !(file == "standard" || file == "thinkertoy" || file == "shadow" || file == "ac") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		}
		switch file {
		case "standard":
			file = "resources/standard.txt"
		case "thinkertoy":
			file = "resources/thinkertoy.txt"
		case "shadow":
			file = "resources/shadow.txt"
		case "ac":
			file = "resources/ac.txt"
		default:
			file = "resources/standard.txt"
		}

		s, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("File not found")
			return
		}

		m := asciiArt.AsciiArt(string(s))
		res := asciiArt.Tab(os.Args[1])
		asciiArt.Paragraph(res, m)

	}
}
