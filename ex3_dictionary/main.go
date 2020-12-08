package main

import (
	"fmt"
	"main/mydict"
)

func main() {
	dictionary := mydict.Dictionary{};
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)

	definition, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	word := "hii"
	definiton := "hii"
	err2 := dictionary.Add(word, definiton)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := dictionary.Update("hii", "dd")
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(dictionary)

	word, err4 := dictionary.Search("hii")
	dictionary.Delete("hii")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println("delete word -> " + word)
	}

	fmt.Println(dictionary)

}
