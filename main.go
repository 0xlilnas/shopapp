package main

import (
	"github.com/0xlilnas/shopapp/src/initiliazers"
	"github.com/gin-gonic/gin"
)

func init() {
	initiliazers.LoadEnvVariables()
	initiliazers.ConnectDB()
}

func main() {

	router := gin.Default()

	router.Run()
}
