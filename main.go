package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	account001 := accounts.CheckingAccount{}
	account001.Deposit(100)

	fmt.Println(account001.GetBalance())
}
