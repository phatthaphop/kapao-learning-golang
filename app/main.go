package main

import (
	// "fmt"

	_service "kp_api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("kp/hello", func(c *gin.Context) { _service.HelloWeb(c) })

	r.Run("localhost" + ":" + "8080")
}
