package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// var users = []user{{"Patrick", "Boateng", "pato@gmail.com"}, {"Gabriel", "Nkansah", "gab@gmail.com"}}

var db *sql.DB
var err error

func getUsers(c *gin.Context) {
	sql_stmt := "SELECT * FROM users;"
	rows, err := db.Query(sql_stmt)

	if err != nil {
		panic(err)
	}

	var users []user

	for rows.Next() {
		var user1 user

		err = rows.Scan(&user1.ID, &user1.FirstName, &user1.LastName, &user1.Email)

		if err != nil {
			fmt.Println("An error occured")
			panic(err)
		}

		users = append(users, user1)

	}
	fmt.Println(users)
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	env, er := godotenv.Read(".env")

	if er != nil {
		fmt.Println("Error loading env file")
		panic(er)
	}

	username := env["USERNAME"]
	passwd := env["PASSWORD"]
	connection_var := fmt.Sprintf("%s:%s@tcp(localhost:3306)/users", username, passwd)

	db, err = sql.Open("mysql", connection_var)

	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/users", getUsers)

	defer db.Close()

	router.Run("localhost:8000")
}
