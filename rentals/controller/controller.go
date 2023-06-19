package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	m "rentals/model"

	r "rentals/repository"

	q "rentals/repository/query"
)

func GetRentals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, r.FindRentals(q.QueryBuilder(m.QueryParameter{Ids: c.Query("ids"), PriceMin: c.Query("price_min"), PriceMax: c.Query("price_max"), Sort: c.Query("sort"), Limit: c.Query("limit"), Offset: c.Query("offset"), Near: c.Query("near")})))
}

func GetRental(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, r.FindRentals(q.IDsQuery(c.Param("id"))))
}
