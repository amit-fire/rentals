package utils

var PricePerDay = "price_per_day"
var Sleeps = "sleeps"
var VehicleYear = "vehicle_year"
var VehicleLength = "vehicle_length"
var Lat = "lat"
var Lng = "lng"

func SortMap() map[string]string {
	m := make(map[string]string)
	m["price"] = PricePerDay
	m["sleeps"] = Sleeps
	m["year"] = VehicleYear
	m["length"] = VehicleLength
	return m
}
