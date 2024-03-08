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

	//Tambem posso fazer assim nativamente
	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "pong")
	//   })

	//   err := http.ListenAndServe(":9000", nil)
	//   if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// 	return
	//   }

}
