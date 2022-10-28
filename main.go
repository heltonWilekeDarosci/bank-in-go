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
	fmt.Println(account001.Deposit(100))
	fmt.Println(account001.balance)
}

func (c *CheckingAccount) Withdraw(withdrawValue float64) string {
	enabled := withdrawValue <= c.balance && withdrawValue > 0

	if enabled {
		c.balance -= withdrawValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Depósito realizado com sucesso", c.balance
	} else {
		return "Depósito não realizado, contate a gerência para mais infomações", c.balance
	}
}
