package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/modules"
	"github.com/vietthangc1/mini-web-golang/tokens"
)

func (a *App) HandlerAddUser(c *gin.Context) {
	var loginUser models.User
	if err := c.BindJSON(&loginUser); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	loginUser.Password, _ = modules.HasingPassword(loginUser.Password)

	loginUser, err := a.Handler.UserRepo.AddUser(loginUser)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, loginUser)
}

func (a *App) HandlerDeleteUser(c *gin.Context) {
	id := c.Param("id")
	var userDelete models.User

	_id, _ := strconv.ParseUint(id, 10, 32)

	userDelete, err := a.Handler.UserRepo.DeleteUser(uint(_id))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	err = a.CacheInstance.Delete(id)
	if err != nil {
		log.Println("Cannot delete from cache")
		log.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"user": userDelete, "message": "Deleted!"})
}

func (a *App) HandlerLogin(c *gin.Context) {
	var loginUser models.User

	if err := c.BindJSON(&loginUser); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := loginUser.Email
	password := loginUser.Password

	loginUser, err := a.Handler.UserRepo.GetUserByEmail(email)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	check := modules.ComparePassword(password, loginUser.Password)
	if !check {
		log.Println("Wrong Email or Password")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Wrong Email or Password"})
		return
	}
	token, err := tokens.GenerateToken(loginUser.Email)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func (a *App) HandlerGetUser(c *gin.Context) {
	var currentUser models.User
	user_email, err := tokens.ExtractTokenEmail(c)
	if err != nil {
		c.IndentedJSON(http.StatusNonAuthoritativeInfo, gin.H{"error": err.Error()})
	}
	currentUser, err = a.Handler.UserRepo.GetUserByEmail(user_email)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentUser.Password = ""
	c.IndentedJSON(http.StatusOK, currentUser)
}
