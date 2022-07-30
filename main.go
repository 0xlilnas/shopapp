package main

import (
	"github.com/0xlilnas/shopapp/src/initiliazers"
	"github.com/0xlilnas/shopapp/src/models"
	"github.com/0xlilnas/shopapp/src/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func init() {
	initiliazers.LoadEnvVariables()
	models.ConnectDataBase()
}

func main() {

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	//router setup
	routes.Setup(router)

	//cors
	router.Use(cors.New(cors.Options{
		AllowCredentials: true,
	}))
	//port
	router.Run()
}
