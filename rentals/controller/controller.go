package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	r "rentals/repository"

	q "rentals/repository/query"
)

func GetRentals(c *gin.Context) {
	ids := c.Query("ids")

	if ids == "" {
		c.IndentedJSON(http.StatusOK, r.FindRentals(q.DefaultQuery()))
	} else {
		c.IndentedJSON(http.StatusOK, r.FindRentals(q.IDQuery(ids)))
	}
}

func GetRental(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, r.FindRentals(q.IDQuery(c.Param("id"))))
}
