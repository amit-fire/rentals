package service

import (
	"errors"
	m "rentals/model"
	r "rentals/repository"
	q "rentals/repository/query"
	u "rentals/utils"
	v "rentals/validator"
)

func GetRentals(qp m.QueryParameter) ([]m.Rental, error) {
	_, err := v.Validate(qp)
	if err == nil {
		return r.FindRentals(q.QueryBuilder(qp)), nil
	}

	return nil, err
}

func GetRental(id string) ([]m.Rental, error) {
	if u.IsPositiveInteger(id) {
		return r.FindRentals(q.IDsQuery(id)), nil
	}

	return nil, errors.New("id " + id + " is not a positive integer")
}
