package router

import (
	"postgres-gin-crud/server"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(s server.Server) *gin.Engine {
	r := gin.Default()

	r.POST("/book", s.CreateBook)
	r.POST("/author", s.CreateAuthor)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
