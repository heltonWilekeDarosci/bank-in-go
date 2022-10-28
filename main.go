package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	Account001 := CheckingAccount{"Helton", 2423, 123456, 150.43}

	fmt.Println(Account001)
}
