package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/ynnekF/gin-boilerplate/models"
	"net/http"
)


func GetHealth(c *gin.Context) {
	response := models.Response{
		Status:         "success",
		Message:        "Service is healthy",
		Error:          "",
		Date:           "",
		Transaction_Id: "",
	}

	c.IndentedJSON(http.StatusOK, response)
}
