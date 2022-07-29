package main

import (
	"github.com/0xlilnas/shopapp/src/initiliazers"
	"github.com/0xlilnas/shopapp/src/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func init() {
	initiliazers.LoadEnvVariables()
	initiliazers.ConnectDB()
	initiliazers.SyncDatabase()
}

func main() {

	router := gin.Default()

	//router setup
	routes.Setup(router)

	//cors
	router.Use(cors.New(cors.Options{
		AllowCredentials: true,
	}))
	//port
	router.Run()
}
