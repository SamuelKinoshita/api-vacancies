package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	// Specify the port in the Run function argument
	err := router.Run(":9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
