package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/timepeace-api-gateway/pkg/api/handler"
	"github.com/sgokul961/timepeace-api-gateway/pkg/api/middleware"
	"github.com/sgokul961/timepeace-api-gateway/pkg/config"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(AdminHandler *handler.AdminHandler, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	userRoutes := engine.Group("/")
	userRoutes.POST("signup", userHandler.SignUp)
	userRoutes.POST("login", userHandler.Login)

	adminRoutes := engine.Group("/admin")
	adminRoutes.POST("login", AdminHandler.Login)

	productRoutes := engine.Group("/product")
	productRoutes.GET("/list", productHandler.ListProducts)
	productRoutes.Use(middleware.AdminAuthorizationMiddleware)
	productRoutes.POST("/add", productHandler.AddProduct)

	return &ServerHTTP{engine: engine}
}
func (sh *ServerHTTP) Start(c config.Config) {
	sh.engine.Run(c.Port)
}
