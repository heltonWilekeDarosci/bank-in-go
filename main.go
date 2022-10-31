package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	account001 := accounts.CheckingAccount{Holder: "Helton", Balance: 500}
	account002 := accounts.CheckingAccount{Holder: "Mark", Balance: 200}

	fmt.Println(account001)
	fmt.Println(account002)
	status := account001.Transfer(100, &account002)

	fmt.Println(status)
	fmt.Println(account001)
	fmt.Println(account002)
}
