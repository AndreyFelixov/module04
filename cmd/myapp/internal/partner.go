package internal

type Partner struct {
	Name    string
	Age     int
	balance int
	debt    int
}

func (p *Partner) WrtOffDebt() error {
	p.debt = 0

	return nil
}

func NewPartner(name string, age int, balance int, debt int) *Partner {
	return &Partner{
		Name:    name,
		Age:     age,
		balance: balance,
		debt:    debt,
	}
}
