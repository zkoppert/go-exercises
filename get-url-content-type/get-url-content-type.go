package main

import (
	"fmt"
	"net/http"
)

const url = "https://github.com/"

func contentType(url string) (string, error) {
	resp, err := http.Get(url) //nolint:golint,G107
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()
	header := resp.Header.Get("Content-Type")
	if header == "" {
		return "", fmt.Errorf("Cannot get Content_Type header")
	}
	return header, nil
}

func main() {
	contentTypeValue, getError := contentType(url)
	if getError != nil {
		fmt.Printf("ERROR: Unable to GET URl\n")
	}
	fmt.Println(contentTypeValue)
}
