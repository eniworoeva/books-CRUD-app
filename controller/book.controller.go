package controller

import (
	"net/http"
	"time"

	"github.com/eniworoeva/books-CRUD-app/database"
	"github.com/eniworoeva/books-CRUD-app/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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


func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		bookId := c.Param("book_id")
		var book model.Book

		objectId, _ := primitive.ObjectIDFromHex(bookId)

		err := bookCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while fetching book."})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}


func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		bookId := c.Param("book_id")
		var book model.Book

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		objectId, _ := primitive.ObjectIDFromHex(bookId)
		filter := bson.M{"_id": objectId}

		var updateObj primitive.D

		if book.Author != nil {
			updateObj = append(updateObj, bson.E{Key: "author", Value: book.Author})
		}

		if book.Title != nil {
			updateObj = append(updateObj, bson.E{Key: "title", Value: book.Title})
		}

		if book.Description != nil {
			updateObj = append(updateObj, bson.E{Key: "description", Value: book.Description})
		}

		book.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: book.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		_ , err := bookCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Book update failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})

	}
}