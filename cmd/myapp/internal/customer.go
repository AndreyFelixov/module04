package internal

import (
	"errors"
)

const DEFAULT_DISCOUNT = 5000

type Customer struct {
	Name         string
	Age          int
	balance      int
	debt         int
	Discount     bool
	CalcDiscount func() (int, error)
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

func CalcPrice(cust *Customer, price int) (int, error) {
	err := errors.New("discount is not avaliable")
	cust.CalcDiscount = func() (int, error) {

		if !cust.Discount {
			return 0, err
		}
		err = nil
		result := DEFAULT_DISCOUNT - cust.debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	finalPrice := 0
	discount, _ := cust.CalcDiscount()
	if _, err := cust.CalcDiscount(); err == nil {
		finalPrice = price - discount
		if cust.balance+discount < price {
			err = errors.New("not enough balance")
			return finalPrice, err
		} else if price == 0 {
			err = errors.New("price can't be 0, its not a razdacha na spavne")
			finalPrice = 0
			return finalPrice, err
		} else if finalPrice < 0 {
			err = errors.New("you have a 100 percent sale")
			finalPrice = 0
			return finalPrice, err
		}
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
