package router

import (
	"github.com/arfan21/hacktiv8-golang-jwt/controller"
	"github.com/arfan21/hacktiv8-golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controller.CreateProduct)

		productRouter.PUT("/:productId", middleware.ProductAuthorization(), controller.UpdateProduct)
	}
	return r
}
