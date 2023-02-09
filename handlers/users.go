package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/modules"
	"github.com/vietthangc1/mini-web-golang/tokens"
)

func HandlerAddUser(c *gin.Context) {
	var loginUser models.User

	if err := c.BindJSON(&loginUser); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}
	id := time.Now().UnixMilli()
	loginUser.ID = strconv.Itoa(int(id))
	loginUser.Password, _ = modules.HasingPassword(loginUser.Password)

	query := "INSERT INTO users (id, email, password) VALUES (?, ?, ?)"

	loginUser, err := modules.QueryAddUser(query, loginUser)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, loginUser)
}

func HandlerDeleteUser(c *gin.Context) {
	id := c.Param("id")

	query := `
	DELETE FROM users 
	WHERE id = ?
	`

	err := modules.QueryDeleteUser(query, id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	err = cacheInstance.Delete(id)
	if err != nil{
		fmt.Println("Cannot delete from cache")
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Deleted!"})
}

func HandlerLogin(c *gin.Context) {
	var loginUser models.User

	if err := c.BindJSON(&loginUser); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := loginUser.Email
	password := loginUser.Password

	query := "SELECT id, email, password FROM users WHERE email = ?"

	loginUser, err := modules.QueryGetUserByEmail(query, email)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	check := modules.ComparePassword(password, loginUser.Password)

	if !check {
		fmt.Println("Wrong Email or Password")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Wrong Email or Password"})
		return
	}

	token, err := tokens.GenerateToken(loginUser.Email)
	if (err != nil) {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func HandlerGetUser(c *gin.Context) {
	user_email, err := tokens.ExtractTokenEmail(c)
	if (err != nil) {
		c.IndentedJSON(http.StatusNonAuthoritativeInfo, gin.H{"error": err.Error()})
	}

	query := "SELECT id, email, password FROM users WHERE email = ?"

	currentUser, err := modules.QueryGetUserByEmail(query, user_email)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentUser.Password = ""

	c.IndentedJSON(http.StatusOK, currentUser)
}