package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==START")

	c := make(chan string)
	people := [3]string{"A","B","C"}
	
	for _, person := range people {
		go print(person, c)
	}


	/*
	//time.Sleep(time.Second * 5)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // blocking operation "<-"
	*/

	// if channel count > variable count
	// all goroutines are asleep - deadlock occurs
	// so, the problem sloved example
	// channel init for loop
	for i:=0; i<len(people); i++{
		fmt.Println(<-c)
	}


	fmt.Println("==END")
}

func print(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + "!"
}
