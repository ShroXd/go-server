package main

import (
	"github.com/gin-gonic/gin"
	"go-server/api"
	"go-server/db"
)

func main() {
	db.Connect()
	r := gin.Default()
	api.RunApiHandlers(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
