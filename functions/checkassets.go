package ascii

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// compares local banner files to the one in cloud
func Checkfiles(s string) error {
	remoteURL := ""
	switch s {
	case "resources/standard.txt":
		remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"

	case "resources/thinkertoy.txt":
		remoteURL = "https://learn.zone01kisumu.ke/git/aadero/ascii-art-fs/raw/branch/wnjuguna/resources/thinkertoy.txt"
	case "resources/shadow.txt":
		remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/shadow.txt"
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

	str, err := os.Create(s)
	if err != nil {
		return err
	}
	defer str.Close()

	_, err = io.Copy(str, remoteFile.Body)
	if err != nil {
		return err
	}

	return nil
}
