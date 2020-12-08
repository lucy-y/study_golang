package main

import (
	"fmt"
	"log"
	"main/accounts"
)

func main() {

	account := accounts.NewAccount("dd")

	account.Deposit(3000)
	fmt.Println(account)
	fmt.Println(account.Banlance())

	err := account.Withdraw(2000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
	fmt.Println(account.Banlance(), account.Owner())

    account.ChangeOwner("bb")
    owner := account.Owner()
    fmt.Println(account.Banlance(), owner)

}

