package ascii

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	
)

// handles command line inputs checking if correctly formatted
// checks if the banner file is modified and is not as it should be
// it downloads the banner file if different from original
// checks flags usage from CLI if has correct format
func Input(input []string) {
	args := len(os.Args)
	if args < 2 || args > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	// preset the checksum values of the files
	standardCheckSum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowCheckSum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyCheckSum := "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52"

	
	if args == 2 {
		banner, err := os.Open("resources/standard.txt")
		if err != nil {
			fmt.Print("Unable to read file.")
			return
		}
		defer banner.Close()
		// a sha-256 hash object is created to store the contents of the banner file
		bannerTemp := sha256.New()
		if _, err := io.Copy(bannerTemp, banner); err != nil {
			log.Fatal(err)
		}
		// the final checksum is computed and converted to hex string
		checkSum := string(fmt.Sprintf("%x", bannerTemp.Sum(nil)))
		if checkSum != standardCheckSum && checkSum != thinkertoyCheckSum && checkSum != shadowCheckSum {
			fmt.Println("File contents have been corrupted. Redownloading the banner file")
			Checkfiles("resources/standard.txt")
			return
		}

		bannerFile, err := os.ReadFile("resources/standard.txt")
		if err != nil {
			fmt.Println("File not found")
			return
		}
		// the bannerfile is mapped
		mapped := AsciiArt(string(bannerFile))
		// handles \t controller aspect in the string
		result := Tab(os.Args[1])
		//the comvined ascii art is then printed
		AsciiCombine(result, mapped)
	}

	
	if args == 3 {
		file := os.Args[2]

		banners := []string{"standard", "thinkertoy", "shadow"}
		for i := range banners {
			if file != banners[i] && i == len(banners)-1 {
				fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
				return
			} else if file == banners[i] {
				break
			}
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

		banner, err := os.Open(file)
		if err != nil {
			fmt.Print("Unable to read file.")
			return
		}
		defer banner.Close()

		bannerTemp := sha256.New()
		if _, err := io.Copy(bannerTemp, banner); err != nil {
			log.Fatal(err)
		}
		checkSum := string(fmt.Sprintf("%x", bannerTemp.Sum(nil)))

		if checkSum != standardCheckSum && checkSum != thinkertoyCheckSum && checkSum != shadowCheckSum {
			fmt.Println("File contents have been corrupted. Redownloading the banner file")
			Checkfiles(file)
			return
		}

		bannerFile, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("File not found")
			return
		}

		mapped := AsciiArt(string(bannerFile))
		result := Tab(os.Args[1])
		AsciiCombine(result, mapped)

	}
}
