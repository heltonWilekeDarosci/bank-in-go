package accounts

type CheckingAccount struct {
	Holder        string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CheckingAccount) Withdraw(withdrawValue float64) string {
	enabled := withdrawValue <= c.Balance && withdrawValue > 0

	if enabled {
		c.Balance -= withdrawValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.Balance += depositValue
		return "Depósito realizado com sucesso", c.Balance
	} else {
		return "Depósito não realizado, contate a gerência para mais infomações", c.Balance
	}
}

func (c *CheckingAccount) Transfer(transferValue float64, destinationAccount *CheckingAccount) bool {
	if transferValue < c.Balance && transferValue > 0 {
		c.Balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}
