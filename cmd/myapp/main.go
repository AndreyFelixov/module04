package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	// price := 14000
	// finalPrice, err := internal.CalcPrice(cust, price)
	// fmt.Printf("Итоговая цена с учетом скидки: %d, %v", finalPrice, err)
	cust.WrtOffDebt()
	fmt.Printf("%+v\n", cust)
}
