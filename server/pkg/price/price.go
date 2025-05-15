package price

type Price float64

func (p Price) ToCent() int {
	return int(p * 100)
}

func ToPrice(cents int) Price {
	return Price(float64(cents) / 100)
}
