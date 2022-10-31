package accounts

import (
	"bank/clients"
)

type CheckingAccount struct {
	Holder                      clients.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
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

func (c *CheckingAccount) Transfer(transferValue float64, destinationAccount *CheckingAccount) bool {
	if transferValue < c.balance && transferValue > 0 {
		c.balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
