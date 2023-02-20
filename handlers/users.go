package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/repository"
	"github.com/vietthangc1/mini-web-golang/tokens"
)

func (h *Handler) HandlerAddUser(c *gin.Context) {
	var loginUser models.User
	if err := c.BindJSON(&loginUser); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	loginUser.Password, _ = repository.HasingPassword(loginUser.Password)

	loginUser, err := h.UserRepo.AddUser(loginUser)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, loginUser)
}

func (h *Handler) HandlerDeleteUser(c *gin.Context) {
	id := c.Param("id")
	var userDelete models.User

	_id, _ := strconv.ParseUint(id, 10, 32)

	userDelete, err := h.UserRepo.DeleteUser(uint(_id))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	err = h.CacheInstance.Delete(id)
	if err != nil {
		log.Println("Cannot delete from cache")
		log.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"user": userDelete, "message": "Deleted!"})
}

func (h *Handler) HandlerLogin(c *gin.Context) {
	var loginUser models.User

	if err := c.BindJSON(&loginUser); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := loginUser.Email
	password := loginUser.Password

	loginUser, err := h.UserRepo.GetUserByEmail(email)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	check := repository.ComparePassword(password, loginUser.Password)
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

func (h *Handler) HandlerGetUser(c *gin.Context) {
	var currentUser models.User
	user_email, err := tokens.ExtractTokenEmail(c)
	if err != nil {
		c.IndentedJSON(http.StatusNonAuthoritativeInfo, gin.H{"error": err.Error()})
	}
	currentUser, err = h.UserRepo.GetUserByEmail(user_email)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentUser.Password = ""
	c.IndentedJSON(http.StatusOK, currentUser)
}
