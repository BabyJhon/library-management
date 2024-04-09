package handlers

import (
	"github.com/BabyJhon/library-managment/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	admins := router.Group("/auth")
	{
		admins.POST("/sign-up", h.signUp)
		admins.GET("/sign-in", h.signIn)
	}

	api := router.Group("/api/v1", h.AdminIdentity)
	{

		users := api.Group("/users") 
		{
			users.POST("/new", h.createUser)          
			users.DELETE("/delete/:id", h.deleteUser) 
			users.GET("/get/:id", h.getUser)          
		}

		books := api.Group("/books")
		{
			books.GET("/all", h.getAllBooks)         
			books.GET("/user/:id", h.getBooksByUser) 
			
			books.GET("/get/:id", h.getBookById)          
			books.DELETE("/delete/:id", h.deleteBookById) 
			books.POST("/new", h.createBook)              

			giveAndReturnBook := books.Group("/:book_id/users/:user_id") 
			{
				giveAndReturnBook.POST("/", h.giveBook)     
				giveAndReturnBook.DELETE("/", h.returnBook) 
			}
		}
	}

	return router
}
