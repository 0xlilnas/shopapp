package routes

import (
	"github.com/0xlilnas/shopapp/src/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Engine) {
	api := c.Group("api")

	//admin routes
	admin := api.Group("admin")
	admin.POST("/register", controllers.Register)
}
