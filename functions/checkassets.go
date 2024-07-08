package ascii

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

//compares local banner files to the one in cloud
func Checkfiles(s string) error {
	remoteURL:=""
	switch s{
	case "resources/text/standard.txt":
	remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"
	
case "resources/text/thinkertoy.txt" :
		remoteURL= "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/thinkertoy.txt"
case "resources/text/shadow.txt"  :
		remoteURL="https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/shadow.txt"
default:
	remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"
	
	}	
	// fetch the content of the remote file
	remoteFile, err := http.Get(remoteURL)
	if err != nil {
		fmt.Println("Error fetching remote file:=", err)
		return err
	}
	defer remoteFile.Body.Close()

	// create a new scanner to read from the remote file
	remoteScanner := bufio.NewScanner(remoteFile.Body)

	// read the content of the local file into a buffer
	var localBuffer bytes.Buffer
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("Error reading ASCII art file:", err)
		os.Exit(0)
	}

	localBuffer.WriteString(string(file))

	// create a scanner to read from the local buffer
	localScanner := bufio.NewScanner(&localBuffer)

	// compare the content of the remote file with the content of the local buffer
	for remoteScanner.Scan() {
		if !localScanner.Scan() || remoteScanner.Text() != localScanner.Text() {
			fmt.Println("Files do not match: Remote file content is different from the local content")
			var Stop string
			fmt.Println("Do you still want to continue:Y or N")
			fmt.Scan(&Stop)
			if Stop == "N" {
				os.Exit(0)
			} else if Stop == "Y" {
				return nil
			}

		}
	}

	return nil
}
