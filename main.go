package main

import (
	"crud_with_sqlite/config"
	"crud_with_sqlite/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database := config.NewDB()
	
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		handlers.CrateUsers(c, database)
	})
	router.Run(":8081")
}
