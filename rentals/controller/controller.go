package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	m "rentals/model"

	s "rentals/service"
)

func GetRentals(c *gin.Context) {
	qp := m.QueryParameter{Ids: c.Query("ids"), PriceMin: c.Query("price_min"), PriceMax: c.Query("price_max"), Sort: c.Query("sort"), Limit: c.Query("limit"), Offset: c.Query("offset"), Near: c.Query("near")}

	results, err := s.GetRentals(qp)

	if err == nil {
		c.IndentedJSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func GetRental(c *gin.Context) {
	result, err := s.GetRental(c.Param("id"))

	if err == nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
