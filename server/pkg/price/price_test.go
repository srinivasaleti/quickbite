package price

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriceConversion(t *testing.T) {
	assert.Equal(t, 1234, Price(12.34).ToCent())
	assert.InDelta(t, 12.34, float64(ToPrice(1234)), 0.001)
	assert.Equal(t, 0, Price(0).ToCent())
	assert.Equal(t, 0.0, float64(ToPrice(0)))
}
