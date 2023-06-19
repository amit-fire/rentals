package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositiveInteger(t *testing.T) {
	assert.True(t, IsPositiveInteger("1"))
}

func TestNegativeInteger(t *testing.T) {
	assert.False(t, IsPositiveInteger("-1"))
}

func TestNotANumber(t *testing.T) { //TODO: provide better name
	assert.False(t, IsPositiveInteger("i am number 6"))
}

func TestNotAnInteger(t *testing.T) {
	assert.False(t, IsPositiveInteger("7.8"))
}

func TestDecimal(t *testing.T) {
	assert.True(t, IsDecimal("1.2"))
}

func TestNotANumber2(t *testing.T) { //TODO: provide better name
	assert.False(t, IsDecimal("you are number 2"))
}
func TestSortMap(t *testing.T) {
	m := SortMap()

	assert.Equal(t, 4, len(m))

	var value string

	value = m["price"]
	assert.Equal(t, "price_per_day", value)

	value = m["sleeps"]
	assert.Equal(t, "sleeps", value)

	value = m["year"]
	assert.Equal(t, "vehicle_year", value)

	value = m["length"]
	assert.Equal(t, "vehicle_length", value)
}
