package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	m "rentals/model"
)

func TestValidId(t *testing.T) {
	result, err := Validate(m.QueryParameter{Ids: "1"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestNegativeId(t *testing.T) {
	result, err := Validate(m.QueryParameter{Ids: "-1"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestInvalidId(t *testing.T) {
	result, err := Validate(m.QueryParameter{Ids: "spoonman"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestValidMinMaxPrice(t *testing.T) {
	result, err := Validate(m.QueryParameter{PriceMin: "1", PriceMax: "2"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInvalidMinMaxPrice(t *testing.T) {
	result, err := Validate(m.QueryParameter{PriceMin: "lima", PriceMax: "quito"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestValidMinPriceWithoutMax(t *testing.T) {
	result, err := Validate(m.QueryParameter{PriceMin: "1"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestValidMaxPriceWithoutMin(t *testing.T) {
	result, err := Validate(m.QueryParameter{PriceMax: "1"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestMinPriceLargerThanMax(t *testing.T) {
	result, err := Validate(m.QueryParameter{PriceMin: "2", PriceMax: "1"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestValidSort(t *testing.T) {
	result, err := Validate(m.QueryParameter{Sort: "price"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInalidSort(t *testing.T) {
	result, err := Validate(m.QueryParameter{Sort: "life goes down hill after preschool"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestValidLimit(t *testing.T) {
	result, err := Validate(m.QueryParameter{Limit: "1"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInvalidLimit(t *testing.T) {
	result, err := Validate(m.QueryParameter{Limit: "wonderkid, traffic law dr"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestValidOffset(t *testing.T) {
	result, err := Validate(m.QueryParameter{Offset: "1"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInvalidOffset(t *testing.T) {
	result, err := Validate(m.QueryParameter{Offset: "el hoyo"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestValidNear(t *testing.T) {
	result, err := Validate(m.QueryParameter{Near: "1,1"})

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestInvalidNearOnlyLatitude(t *testing.T) {
	result, err := Validate(m.QueryParameter{Near: "1,"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestInvalidNearLongitudeNotANumber(t *testing.T) {
	result, err := Validate(m.QueryParameter{Near: "1,put it in one of your codes"})

	assert.False(t, result)
	assert.NotNil(t, err)
}

func TestValidNoQueryParameters(t *testing.T) {
	result, err := Validate(m.QueryParameter{})

	assert.True(t, result)
	assert.Nil(t, err)
}
