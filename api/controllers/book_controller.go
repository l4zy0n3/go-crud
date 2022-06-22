package controllers

import (
	"net/http"

	"github.com/l4zy0n3/go-crud/api/models"
	"github.com/l4zy0n3/go-crud/api/responses"
)

func (server *Server) GetBooks(w http.ResponseWriter, r *http.Request) {

	book := models.Book{}
	found, err := book.GetAllBooks(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, found)
}
