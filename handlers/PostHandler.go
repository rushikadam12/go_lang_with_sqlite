package handlers

import (
	"crud_with_sqlite/config"
	"crud_with_sqlite/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrateUsers(c *gin.Context, db *config.DB) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.CreateUser(&user)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func GetAllUsers(c *gin.Context, db *config.DB) {
	users, err := db.GetUsers()

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
