package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	startTransaction(cust)
	// partner := internal.NewPartner("Andrey", 25, 15000, 2500)
	// startTransaction(partner)
	// fmt.Printf("%+v\n", partner)
	price := 2222
	finalPrice, err := cust.CalcPrice(cust, price)
	fmt.Printf("Final price: %d; %v\n", finalPrice, err)
	fmt.Printf("%+v\n", cust)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrtOffDebt()
}
