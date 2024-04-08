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

func (h *Handler) InitRoutes() *gin.Engine { //Engine is the framework's instance, it contains the muxer, middleware and configuration settings
	router := gin.New()

	admins := router.Group("/auth")
	{
		admins.POST("/sign-up", h.signUp)
		admins.GET("/sign-in", h.signIn)
	}

	api := router.Group("/api/v1", h.AdminIdentity)
	{

		users := api.Group("/users") //для таблицы юзеров
		{
			users.POST("/new", h.createUser)          //добавление пользователя +
			users.DELETE("/delete/:id", h.deleteUser) //удаление пользователя +
			users.GET("/get/:id", h.getUser)          //инфа а пользователе +
		}

		books := api.Group("/books") //для таблицы книг
		{
			//получение книг
			books.GET("/all", h.getAllBooks)         //все книги +
			books.GET("/user/:id", h.getBooksByUser) //все книги одного пользователя +

			//операции с одной книгой
			books.GET("/get/:id", h.getBookById)          //инфа о книге +
			books.DELETE("/delete/:id", h.deleteBookById) //удаление +
			books.POST("/new", h.createBook)              //добавление новой книги +

			//для связующей таблицы юзер-книга
			//функционал выдачи книги пользователю и возврата книги в библиотеку
			giveAndReturnBook := books.Group("/:book_id/users/:user_id") //сначала book_id, потом user_id
			{
				giveAndReturnBook.POST("/", h.giveBook)     // пока просто добавляет без логики
				giveAndReturnBook.DELETE("/", h.returnBook) // пока просто возвращает без логики
			}
		}
	}

	return router
}
