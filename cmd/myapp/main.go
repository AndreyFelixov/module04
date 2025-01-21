package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {

			return 0, errors.New("discount not avaliable")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	discount, _ := cust.CalcDiscount()
	fmt.Printf("%+v", cust)
	fmt.Printf("\n%d", discount)
}
