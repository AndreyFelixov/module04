package internal

type Overduer struct {
	balance int
	debt    int
}

func NewOverduer(balance int, debt int) *Overduer {
	return &Overduer{
		balance: balance,
		debt:    debt,
	}
}
func (ov *Overduer) GetBalance() int {
	return ov.balance
}
func (ov *Overduer) GetDebt() int {
	return ov.debt
}
