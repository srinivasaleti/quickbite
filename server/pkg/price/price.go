package price

type Price float64
type Cent int

const Magitude = 100

func (p Price) ToCents() Cent {
	return Cent(p * Magitude)
}

func (c Cent) ToPrice() Price {
	return Price(float64(c) / Magitude)
}

func (c Cent) Multiply(quantity int) Cent {
	return Cent(c * Cent(quantity))
}

func (c Cent) Add(other Cent) Cent {
	return Cent(int(c) + int(other))
}

func (c Cent) Subtract(other Cent) Cent {
	return Cent(int(c) - int(other))
}

func (c Cent) Percentize(percent float64) Cent {
	return Cent((float64(c) * percent / 100.0))
}
