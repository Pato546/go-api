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
	Password  string `json:"-"`
}

// var users = []user{{"Patrick", "Boateng", "pato@gmail.com"}, {"Gabriel", "Nkansah", "gab@gmail.com"}}

var db *sql.DB
var err error

func signUp(c *gin.Context) {
	var newUser user

	err := c.BindJSON(&newUser)

	if err != nil {
		return
	}

	// PASSWORD ecryption will occur Here

	r, err := db.Exec("insert into users(first_name, last_name, email, password) values(?, ?, ?, ?)", newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)

	fmt.Println(r)

	if err != nil {
		fmt.Println("An error occured")
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func getUsers(c *gin.Context) {
	sql_stmt := "SELECT * FROM users;"
	rows, err := db.Query(sql_stmt)

	if err != nil {
		panic(err)
	}

	var users []user

	for rows.Next() {
		var user1 user

		err = rows.Scan(&user1.ID, &user1.FirstName, &user1.LastName, &user1.Email, &user1.Password)

		if err != nil {
			fmt.Println("An error occured")
			panic(err)
		}

		users = append(users, user1)

	}
	fmt.Println(users)
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	var newUser user
	var foundUser bool = false

	id := c.Param("id")
	sql_stmt := "select * from users where id=?"
	rows, err := db.Query(sql_stmt, id)

	if err != nil {
		fmt.Println("An error occured")
		c.JSON(http.StatusNotFound, nil)
	}

	for rows.Next() {
		err = rows.Scan(&newUser.ID, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.Password)

		if err != nil {
			fmt.Println("An error occured")
			panic(err)
		}

		foundUser = true
	}

	if !foundUser {
		fmt.Println("An error occured")
		c.JSON(http.StatusNotFound, nil)
		panic(err)
	}

	c.JSON(http.StatusOK, newUser)
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
	router.GET("users/:id", getUserByID)
	router.POST("/signup", signUp)

	defer db.Close()

	router.Run("localhost:8000")
}
