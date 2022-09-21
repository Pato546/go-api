package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users = []user{{"Patrick", "Boateng", "pato@gmail.com"}, {"Gabriel", "Nkansah", "gab@gmail.com"}}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:8000")
}
