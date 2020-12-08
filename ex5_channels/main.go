package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==START")

	c := make(chan bool)
	people := [2]string{"A","B"}
	
	for _, person := range people {
		go print(person, c)
	}
	time.Sleep(time.Second * 5)
	result := <-c
	fmt.Println(result)
	fmt.Println("==END")
}

func print(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
