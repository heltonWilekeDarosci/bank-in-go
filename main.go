package main

import (
	"fmt"
	"os"
)

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {

	showIntro()

	for {
		showMenu()
		command := seeCommand()

		switch command {
		case 1:
		case 2:
		case 3:
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Command unknown")
			os.Exit(-1)
		}
	}

}

func PayBill(account verifyAccount, billValue float64) {
	account.Withdraw(billValue)
	fmt.Println("Seu boleto foi pago com sucesso")
}

func showIntro() {
	var name string
	version := 1.2
	fmt.Println("Hello", name, "Welcome!")
	fmt.Println("This program runs on version:", version)
}

func showMenu() {
	fmt.Println("")
	fmt.Println("1- Register a new client")
	fmt.Println("2- Deposit")
	fmt.Println("3- Withdraw")
	fmt.Println("0- Exit")
}

func seeCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("")

	return commandRead
}
