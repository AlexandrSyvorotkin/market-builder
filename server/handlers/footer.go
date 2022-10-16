package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/schemas"
)

func PostFooter(c *gin.Context) {
	var newData schemas.FooterSchema
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error occured while post header": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newData)
}
