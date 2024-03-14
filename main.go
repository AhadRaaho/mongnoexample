// main.go
package main

import (
	"refactor/api"
	"refactor/dao"
)

func main() {
	// Initialize database connection
	err := dao.InitDB()
	if err != nil {
		// Handle the error
		panic(err)
	}

	// Run the API server
	api.StartServer()
}
