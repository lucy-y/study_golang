package accounts

import (
    "errors"
    "fmt"
)

var errNoMoney = errors.New("Can't withdraw")

// Account struct
type Account struct {
	owner		string
	banlance	int
}

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, banlance: 0 }
	return &account
}

// Deposit x amount on you account
func (account *Account) Deposit(amount int) {
	account.banlance += amount
}

// Withdraw x amount from you acoount
func (account *Account) Withdraw(amount int) error {
	if account.banlance < amount {
		return errNoMoney
	}
	account.banlance -= amount
	return nil
}

// Blance of your account
func (account Account) Banlance() int {
	return account.banlance
}

// ChangeOwner of the account 
func (account *Account) ChangeOwner(newOwner string) {
    account.owner = newOwner
}

// Owner of the account
func (account Account) Owner() string {
    return account.owner
}

func (account Account) String() string {
    // 구조 정의 후 호출
    return fmt.Sprint(account.Owner(), ": ", account.Banlance())
}


