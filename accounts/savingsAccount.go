package accounts

import "bank/clients"

type SavingsAccount struct {
	Holder                          clients.Holder
	AgencyNumber, AccountNumber, Op int
	balance                         float64
}

func (c *SavingsAccount) Withdraw(withdrawValue float64) string {
	enabled := withdrawValue <= c.balance && withdrawValue > 0

	if enabled {
		c.balance -= withdrawValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *SavingsAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Depósito realizado com sucesso", c.balance
	} else {
		return "Depósito não realizado, contate a gerência para mais infomações", c.balance
	}
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
