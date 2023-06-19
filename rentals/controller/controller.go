package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	r "rentals/repository"
)

func GetRentals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, r.FindRentals())
}
