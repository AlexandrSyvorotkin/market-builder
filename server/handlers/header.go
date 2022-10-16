package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/schemas"
)

func PostHeader(c *gin.Context) {
	var newData schemas.HeaderSchema
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error occured while post header": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newData)
}
