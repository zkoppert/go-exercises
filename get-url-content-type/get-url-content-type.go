package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
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
	url := "https://github.com/"
	contentTypeValue, getError := contentType(url)
	if getError != nil {
		fmt.Printf("ERROR: Unable to GET URl\n")
	}
	fmt.Println(contentTypeValue)
}
