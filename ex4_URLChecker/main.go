package main

import (
	"fmt"
	"errors"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// var results = map[string]string{}
	var results = make(map[string]string)

	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.daum.net/",
	}

	
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		} 	
		results[url] = result
	}

	for url, result := range results {
		fmt.Println("url: ", url)
		fmt.Println("status: ", result)
	}
	
}

func hitURL(url string) error{
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	} 
	return nil
}
