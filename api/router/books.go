package router

import (
	"github.com/gin-gonic/gin"
	"go-server/api/handler"
)

func BooksRouter(engine *gin.Engine) {
	// 统一管理
	v1 := engine.Group("/api/v1")
	v1.GET("/books", handler.GetBooks)
}
