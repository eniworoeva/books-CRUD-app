package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/eniworoeva/books-CRUD-app/controller"
)



func BookRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("books/create", controllers.CreateBook())
	incomingRoutes.GET("books/:book_id", controllers.GetBook())
	//incomingRoutes.PATCH("books/:book_id", controllers.Updatebook())
	//incomingRoutes.DELETE("books/:book_id", controllers.DeleteBook())
	//incomingRoutes.GET("books", controllers.GetAllBook())
}