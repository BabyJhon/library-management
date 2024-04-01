package handlers

import (
	"net/http"
	"strconv"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.services.Book.GetAllBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) getBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.Book.GetBook(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         book.Id,
		"title":      book.Title,
		"author":     book.Author,
		"in_library": book.InLibrary,
	})
}


func (h *Handler) deleteBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Book.DeleteBook(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) createBook(c *gin.Context) {
	var inPut entity.Book

	if err := c.BindJSON(&inPut); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Book.CreateBook(inPut)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
