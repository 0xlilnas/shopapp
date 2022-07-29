package controllers

import (
	"net/http"

	"github.com/0xlilnas/shopapp/src/initiliazers"
	"github.com/0xlilnas/shopapp/src/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var data map[string]string

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if data["password"] != data["confirm_password"] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password do not match!",
		})

		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password!",
		})

		return
	}

	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		Password:     string(password),
		IsAmbassador: false,
	}

	result := initiliazers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user. User already existed!",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {

}
