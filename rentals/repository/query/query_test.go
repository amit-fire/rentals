package query

import (
	m "rentals/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base = `SELECT r.id, name, description, type, vehicle_make, vehicle_model, vehicle_year, vehicle_length, sleeps, price_per_day, primary_image_url, home_city, home_state, home_zip, home_country, lat, lng, u.id, first_name, last_name FROM rentals r INNER JOIN users u ON user_id = u.id`
var idIsOne = base + ` WHERE r.id in (1)`
var orderById = ` ORDER BY r.id ASC`

func TestIDsQuery(t *testing.T) {
	assert.Equal(t, idIsOne, IDsQuery("1"))
}

func TestIDs(t *testing.T) {
	var expected = idIsOne + orderById
	assert.Equal(t, expected, QueryBuilder(m.QueryParameter{Ids: "1"}))
}

func TestMinMaxPrice(t *testing.T) {
	var expected = base + ` WHERE price_per_day > 1 AND price_per_day < 2` + orderById
	assert.Equal(t, expected, QueryBuilder(m.QueryParameter{PriceMin: "1", PriceMax: "2"}))
}

func TestSort(t *testing.T) {
	var expected = base + ` ORDER BY vehicle_year ASC`
	assert.Equal(t, expected, QueryBuilder(m.QueryParameter{Sort: "year"}))
}

func TestLimit(t *testing.T) {
	var expected = base + orderById + ` LIMIT 1`
	assert.Equal(t, expected, QueryBuilder(m.QueryParameter{Limit: "1"}))
}

func TestOffset(t *testing.T) {
	var expected = base + orderById + ` OFFSET 1`
	assert.Equal(t, expected, QueryBuilder(m.QueryParameter{Offset: "1"}))
}

func TestNear(t *testing.T) {
	var expected = base + ` WHERE lat > 1-100 AND lat < 1+100 AND lng > 2-100 AND lng < 2+100` + orderById
	assert.Equal(t, expected, QueryBuilder(m.QueryParameter{Near: "1,2"}))
}

func TestAll(t *testing.T) {
	var expected = base + ` WHERE r.id in (1,2,3) AND price_per_day > 4 AND price_per_day < 5 AND lat > 8-100 AND lat < 8+100 AND lng > 9-100 AND lng < 9+100 ORDER BY vehicle_length ASC LIMIT 6 OFFSET 7`
	assert.Equal(t, expected, QueryBuilder(m.QueryParameter{Ids: "1,2,3", PriceMin: "4", PriceMax: "5", Sort: "length", Limit: "6", Offset: "7", Near: "8,9"}))
}
