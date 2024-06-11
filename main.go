package main

import (
	"fmt"
	"os"

	asciiArt "ascii/functions"
)

func main() {
	args := len(os.Args)
	if args < 2 || args > 3{
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	//if only the string to be printed is provided
	if args == 2 {
		s, err := os.ReadFile("Resources/standard.txt")
		if err != nil {			
			fmt.Println("File not found")
			return
		}
		m := asciiArt.AsciiArt(string(s))
		res := asciiArt.Tab(os.Args[1])
		asciiArt.Paragraph(res, m)
	}
	
	//if the string to be printed is provided and also the bannerfile
	if args == 3 {
		file := os.Args[2]
		switch file {
		case "standard":
			file = "Resources/standard.txt"
		case "thinkertoy":
			file = "Resources/thinkertoy.txt"
		case "shadow":
			file = "Resources/shadow.txt"
		default:
			file = "Resources/standard.txt"
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

