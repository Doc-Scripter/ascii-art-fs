package main

import (
	"bytes"
	"os"
	"testing"

	a "ascii/functions"
)

var testCases = []struct {
	name          string
	expectedLines string
}{
	{"hello", read()},
}

func TestParagraph(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			s, _ := os.ReadFile("standard.txt")

			m := a.AsciiArt(string(s))
			a.Art("hello", m)
			w.Close()
			os.Stdout = old
			var got bytes.Buffer
			_, err := got.ReadFrom(r)
			if err != nil {
				t.Fatal("error reading from os out")
			}
			expected := tc.expectedLines
			if expected != got.String() {
				t.Errorf("expected\n%v\n got\n%v", expected, got.String())
			}
		})
	}
}

func read() string {
	word, err := os.ReadFile("test.txt")
	if err != nil {
		panic("error reading file")
	}
	return string(word)
}
