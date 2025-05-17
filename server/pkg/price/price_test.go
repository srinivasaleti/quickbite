package price

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriceConversion(t *testing.T) {
	assert.Equal(t, Cent(1234), Dollar(12.34).ToCents())
	assert.Equal(t, Dollar(12.34), Cent(1234).ToDollar())

	assert.Equal(t, Cent(1200), Dollar(12).ToCents())
	assert.Equal(t, Dollar(12), Cent(1200).ToDollar())

	assert.Equal(t, Cent(0), Dollar(0).ToCents())
	assert.Equal(t, Dollar(0.0), Cent.ToDollar(0))

	assert.Equal(t, Cent(100).Multiply(10), Cent(1000))
	assert.Equal(t, Cent(10).Multiply(5), Cent(50))
}

func TestCentOperations(t *testing.T) {
	price := Cent(1500)

	assert.Equal(t, Cent(4500), price.Multiply(3))

	assert.Equal(t, Cent(1550), price.Add(Cent(50)))
	assert.Equal(t, Cent(1450), price.Subtract(Cent(50)))

	assert.Equal(t, Cent(270), price.Percentize(18.0))
	assert.Equal(t, Cent(277), price.Percentize(18.5))

	assert.Equal(t, Cent(0), price.Percentize(0))
	assert.Equal(t, price, price.Percentize(100))
}
