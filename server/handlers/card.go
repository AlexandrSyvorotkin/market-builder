package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/schemas"
)

func PostCard(c *gin.Context) {
	var newData schemas.CardSchema
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error occured while post body": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newData)
}
