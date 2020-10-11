package api

import (
	"github.com/gin-gonic/gin"
	"go-server/api/router"
)

func RunApiHandlers(engine *gin.Engine) {
	router.BooksRouter(engine)
}
