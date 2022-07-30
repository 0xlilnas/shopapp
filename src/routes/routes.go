package routes

import (
	"github.com/0xlilnas/shopapp/src/controllers"
	"github.com/0xlilnas/shopapp/src/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Engine) {

	public := c.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	admin := c.Group("/api/admin")
	admin.Use(middleware.JwtAuthMiddleware())
	admin.GET("/user", controllers.User)
}
