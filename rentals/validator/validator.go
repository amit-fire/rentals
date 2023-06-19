package validator

import (
	"errors"
	"strconv"
	"strings"

	m "rentals/model"

	u "rentals/utils"
)

func Validate(qp m.QueryParameter) (bool, error) {
	var err error
	_, err = validateIds(qp.Ids)

	if err != nil {
		return false, err
	}

	_, err = validatePrice(qp.PriceMin, qp.PriceMax)

	if err != nil {
		return false, err
	}

	_, err = validateSort(qp.Sort)

	if err != nil {
		return false, err
	}

	_, err = validateValue(qp.Limit, "limit")

	if err != nil {
		return false, err
	}

	_, err = validateValue(qp.Offset, "offset")

	if err != nil {
		return false, err
	}

	_, err = validateNear(qp.Near)

	if err != nil {
		return false, err
	}

	return true, nil
}

func validateIds(ids string) (bool, error) {
	if ids != "" {
		s := strings.Split(ids, ",")
		for i := 0; i < len(s); i++ {
			if !u.IsPositiveInteger(s[i]) {
				return false, errors.New("id " + s[i] + " is not a positive integer")
			}
		}
	}

	return true, nil
}

func validatePrice(priceMin string, priceMax string) (bool, error) {
	var minPrice int
	var err error
	if priceMin != "" {
		minPrice, err = strconv.Atoi(priceMin)
		if err != nil {
			return false, errors.New("min price " + priceMin + " is not a number")
		}
	}

	var maxPrice int
	if priceMax != "" {
		maxPrice, err = strconv.Atoi(priceMax)
		if err != nil {
			return false, errors.New("max price " + priceMax + " is not a number")
		}
	}

	if priceMin != "" && priceMax != "" {
		if minPrice < maxPrice {
			return true, nil
		} else {
			return false, errors.New("max price " + priceMax + " is smaller than min price " + priceMin)
		}
	}

	return true, nil
}

func validateSort(sort string) (bool, error) {
	if sort == "" {
		return true, nil
	}

	_, key := u.SortMap()[sort]
	if key {
		return true, nil
	} else {
		return false, errors.New(sort + " is not a valid sort value")
	}
}

func validateValue(value string, msg string) (bool, error) {
	if value == "" {
		return true, nil
	}

	if !u.IsPositiveInteger(value) {
		return false, errors.New(msg + " " + value + " is not a number")
	}

	return true, nil
}

func validateNear(near string) (bool, error) {
	if near != "" {
		s := strings.Split(near, ",")
		if len(s) == 2 {
			if !u.IsDecimal(s[0]) {
				return false, errors.New(s[0] + " is not a valid latitude value")
			}

			if !u.IsDecimal(s[1]) {
				return false, errors.New(s[0] + " is not a valid longitude value")
			}
		} else {
			return false, errors.New("expected near structure x,y")
		}
	}

	return true, nil
}
