package price

type Price float64
type Cent int

func (p Price) ToCents() Cent {
	return Cent(p * 100)
}

func (c Cent) ToPrice() Price {
	return Price(float64(c) / 100)
}

func (c Cent) Multiply(quantity int) Cent {
	return Cent(c * Cent(quantity))
}
