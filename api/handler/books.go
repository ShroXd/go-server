package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-server/db"
	"go-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetBooks(c *gin.Context) {
	mongo := db.Mongo

	ctx := context.Background()
	var books []models.Books
	collection := mongo.Collection("books")

	if err := collection.Find(ctx, bson.M{}).Limit(10).All(&books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg": "Success",
		"data": books,
	})
}
