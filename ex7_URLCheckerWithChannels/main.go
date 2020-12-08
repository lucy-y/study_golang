package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url		string
	status	string
}

func main() {

	var results = map[string]string{}
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.daum.net/",
	}

	c := make(chan requestResult) // channel

	for _, url := range urls {		
		go hitURL(url, c)
	}

	for i:=0; i<len(urls); i++{
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}

// channel Type
// send-only : (c chan<- type)
// send, write: (c chan type)
func hitURL(url string, c chan<- requestResult){
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} 
	c <- requestResult{url: url, status: status}
}


