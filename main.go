package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	account001 := CheckingAccount{}
	account001.holder = "Helton"
	account001.agencyNumber = 123
	account001.accountNumber = 123456
	account001.balance = 500

	fmt.Println(account001.balance)
	fmt.Println(account001.withdraw(200))
	fmt.Println(account001.balance)
}

func (c *CheckingAccount) withdraw(withdrawValue float64) string {
	enabled := withdrawValue <= c.balance && withdrawValue > 0

	if enabled {
		c.balance -= withdrawValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}
