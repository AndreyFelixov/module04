package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	// price := 14000
	// finalPrice, err := cust.CalcPrice(price)
	// fmt.Printf("Итоговая цена с учетом скидки: %d, %v\n", finalPrice, err)
	cust.WrtOffDebt()
	fmt.Printf("%+v\n", cust)
}
