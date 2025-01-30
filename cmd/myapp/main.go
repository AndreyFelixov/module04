package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	partner := internal.NewPartner("Andrey", 25, 15000, 2500)
	k := 95
	fmt.Println(startTransactionDynamic(cust))
	fmt.Println(startTransactionDynamic(partner))
	fmt.Println(startTransactionDynamic(k))
	fmt.Printf("%+v\n", cust)
	fmt.Printf("%+v\n", partner)
}

func startTransactionDynamic(i interface{}) error {
	debtor, ok := i.(internal.Debtor)
	if !ok {
		return errors.New("incorrect type")
	}
	return debtor.WrtOffDebt()
}
