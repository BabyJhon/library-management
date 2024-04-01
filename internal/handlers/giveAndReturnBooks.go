package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) giveBook(c *gin.Context) {
	book_id, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid book_id param")
		return
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id param")
		return
	}

	id, err := h.services.Book.GiveBookToUser(user_id, book_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) returnBook(c *gin.Context) {
	book_id, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid book_id param")
		return
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id param")
		return
	}

	err = h.services.Book.ReturnBookFromUser(user_id, book_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getBooksByUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	books, err := h.services.Book.GetBooksByUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}
