package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-server/db"
	"go-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	db.Connect()
	defer func() {
		if err := db.Client.Close(context.Background()); err != nil {
			panic(err)
		}
	}()
	collection := db.Mongo.Collection("books")

	ctx := context.Background()
	var (
		books    []models.Books
		page, _  = strconv.ParseInt(c.Query("list_page"), 10, 64)
		limit, _ = strconv.ParseInt(c.Query("list_limit"), 10, 64)
		bookName = c.DefaultQuery("book_name", "")
	)

	if err := collection.Find(ctx, bson.M{
		"bookName": bson.M{
			"$regex": bookName,
		},
	}).Skip((page - 1) * limit).Limit(limit).All(&books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data": gin.H{
			"total": len(books),
			"books": books,
		},
	})
}
