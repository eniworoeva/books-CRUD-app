package controller

import (
	"net/http"
	"time"

	"github.com/eniworoeva/books-CRUD-app/database"
	"github.com/eniworoeva/books-CRUD-app/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var bookCollection *mongo.Collection = database.OpenCollection(database.Client, "book")
var validate = validator.New()

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var book model.Book

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validatorErr := validate.Struct(book)
		if validatorErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validatorErr.Error()})
			return
		} 

		book.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		book.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		book.ID = primitive.NewObjectID()

		result, insertErr :=  bookCollection.InsertOne(ctx, book)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Book item was created."})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}