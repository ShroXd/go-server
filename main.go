package main

import (
	"github.com/gin-gonic/gin"
	"go-server/api"
	"go-server/middleware"
)

func main() {
	gin.ForceConsoleColor()

	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	api.RunApiHandlers(r)

	r.Run(":4040") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
