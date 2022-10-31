package main

import (
	"bank/accounts"
	"fmt"
)

func PayBill(account verifyAccount, billValue float64) {
	account.Withdraw(billValue)
	fmt.Println("Seu boleto foi pago com sucesso")
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	account001 := accounts.SavingsAccount{}
	account001.Deposit(100)
	PayBill(&account001, 30)

	fmt.Println(account001.GetBalance())
}
