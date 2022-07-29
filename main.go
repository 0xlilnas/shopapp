package main

import (
	"github.com/0xlilnas/shopapp/src/initiliazers"
	"github.com/0xlilnas/shopapp/src/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initiliazers.LoadEnvVariables()
	initiliazers.ConnectDB()
	initiliazers.SyncDatabase()
}

func main() {

	router := gin.Default()

	routes.Setup(router)

	router.Run()
}
