package utils

import "strconv"

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

func IsPositiveInteger(value string) bool {
	val, _ := strconv.Atoi(value)
	return val > 0
}

func IsDecimal(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}
