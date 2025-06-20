package service

import (
	"fmt"
	_model "kp_api/models"

	// "kp_api/port"

	"github.com/gin-gonic/gin"
)

func HelloWeb(c *gin.Context) {

	var request _model.HelloModelRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, _model.HelloModelReponse{Message: "Invalid request"})
		return
	}

	if request.Name == "" || request.LastName == "" {
		c.JSON(400, _model.HelloModelReponse{Message: "Name and LastName cannot be empty"})
		return
	}

	resp := fmt.Sprintf("Hello %s %s!", request.Name, request.LastName)
	c.JSON(200, _model.HelloModelReponse{Message: resp})
}
