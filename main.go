package main

import (
	"net/http"

	"github.com/0xlilnas/shopapp/src/initiliazers"
	"github.com/gin-gonic/gin"
)

func init() {
	initiliazers.LoadEnvVariables()
	initiliazers.ConnectDB()
	initiliazers.SyncDatabase()
}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})

	router.Run()
}
