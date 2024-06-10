package main

import (
	"fmt"
	"os"

	a "ascii/functions"
)

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Println("Enter message to be printed")
		return
	}

	s, err := os.ReadFile("standard.txt")
	// confirms if the file used is available
	if err != nil {
		// confirms if all the lines on the ascii art file are there

		fmt.Println("File not found")
		return
	}
	

	m := a.AsciiArt(string(s))
	res := a.Tab(os.Args[1])
	a.Paragraph(res, m)
}
