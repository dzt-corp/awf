package utils

import (
	"fmt"
	"github.com/h2non/filetype"
	"github.com/melbahja/got"
	"io/ioutil"
	"os"
	"strings"
)

// extFromUrl determines the file extension by inferring it from the URL.
func extFromUrl(url string) (string, error) {
	bits := strings.Split(url, "/")
	bit := bits[len(bits)-1]
	if strings.Index(bit, ".") >= 0 {
		bits = strings.Split(bit, ".")
		bit = bits[len(bits)-1]
		return bit, nil
	}
	return "", fmt.Errorf("no extension in URL %s", url)
}

// extFromFile determines the file extension by inferring it from its MIME type.
func extFromFile(filepath string) (string, error) {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	kind, err := filetype.Match(buf)
	if err != nil {
		return "", err
	}
	if kind == filetype.Unknown {
		return "", fmt.Errorf("could not determine ext for %s", filepath)
	}

	return kind.Extension, nil
}

// DownloadFile downloads the file with the given URL to a temporary file with
// the identifier as the name. It tries to determine the extension using the
// URL and the MIME type of the downloaded file.
// Returns the path to the downloaded file.
func DownloadFile(identifier string, url string) (string, error) {
	var err error = nil

	TEMP_DIR := "/tmp"
	filepath := fmt.Sprintf("%s/%s", TEMP_DIR, identifier)

	client := got.New()
	err = client.Download(url, filepath)
	if err != nil {
		return "", err
	}

	ext, err := extFromUrl(url)
	if err != nil {
		ext, err = extFromFile(filepath)
		if err != nil {
			return "", err
		}
	}
	finalPath := fmt.Sprintf("%s.%s", filepath, ext)
	err = os.Rename(filepath, finalPath)
	if err != nil {
		return "", err
	}

	return finalPath, err
}
