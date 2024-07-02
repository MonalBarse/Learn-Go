package main

import (
	"fmt"
	"net/http"

	"github.com/MonalBarse/Learn-Go/tree/main/24-API-MongoDB/router"
)

func main() {
	fmt.Println("MongoDB API")

	router := router.Router()
	fmt.Println("Setting up Server")
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
