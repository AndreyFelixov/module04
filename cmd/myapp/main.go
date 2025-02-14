package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	customer := internal.NewCustomer("Dmitry", 23, true, 10000, 1000)
	fmt.Println("Customer:", customer.Name, customer.Age, customer.Discount)
	fmt.Println("Overduer: Balance =", customer.GetBalance(), ", Debt =", customer.GetDebt())
	startTransaction(customer)
	fmt.Println("Overduer: Balance =", customer.GetBalance(), ", Debt =", customer.GetDebt())
}

func startTransaction(debtor internal.Debtor) error {

	return debtor.WrtOffDebt()
}
