package internal

import (
	"errors"
	"fmt"
)

const DEFAULT_DISCOUNT = 5000

type Customer struct {
	Name     string
	Age      int
	balance  int
	debt     int
	Discount bool
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		debt:     debt,
		Discount: discount,
	}
}

func (c *Customer) CalcDiscount() (int, error) {
	err := errors.New("discount is not avaliable")
	if !c.Discount {
		return 0, err
	}
	err = nil
	result := DEFAULT_DISCOUNT - c.debt
	if result < 0 {
		return 0, nil
	}
	return result, nil

}

func CalcPrice(discounter Discounter, price int) (int, error) {
	finalPrice := 0
	discount, err := discounter.CalcDiscount()
	if err == nil {
		finalPrice = price - discount
	}
	if finalPrice < 0 {
		fmt.Println("You have 100 percent sale")
		finalPrice = 0
	}
	return finalPrice, err
}

func (c *Customer) WrtOffDebt() error {
	if c.debt >= c.balance {
		return errors.New("not possible write off")
	}
	c.balance -= c.debt
	c.debt = 0

	return nil
}
