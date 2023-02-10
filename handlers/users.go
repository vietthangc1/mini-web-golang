package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/modules"
	"github.com/vietthangc1/mini-web-golang/tokens"
)

func (h *BaseHandler) HandlerAddUser(c *gin.Context) {
	var loginUser models.User

	if err := c.BindJSON(&loginUser); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	loginUser.Password, _ = modules.HasingPassword(loginUser.Password)

	err := modules.AddUser(h.db, &loginUser)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, loginUser)
}

func (h *BaseHandler) HandlerDeleteUser(c *gin.Context) {
	id := c.Param("id")
	var userDelete models.User

	_id, _ := strconv.ParseUint(id, 10, 32)
	
	err := modules.DeleteUser(h.db, &userDelete, uint(_id))
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

func (h *BaseHandler) HandlerLogin(c *gin.Context) {
	var loginUser models.User

	if err := c.BindJSON(&loginUser); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := loginUser.Email
	password := loginUser.Password

	err := modules.GetUserByEmail(h.db, &loginUser, email)
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

func (h *BaseHandler) HandlerGetUser(c *gin.Context) {
	var currentUser models.User
	user_email, err := tokens.ExtractTokenEmail(c)
	if (err != nil) {
		c.IndentedJSON(http.StatusNonAuthoritativeInfo, gin.H{"error": err.Error()})
	}
	err = modules.GetUserByEmail(h.db, &currentUser, user_email)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentUser.Password = ""
	c.IndentedJSON(http.StatusOK, currentUser)
}