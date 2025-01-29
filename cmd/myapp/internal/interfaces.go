package internal

type Debtor interface {
	WrtOffDebt() error
}

type Discounter interface {
	CalcDiscount() (int, error)
}
