package store

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Handler functions for the controllers
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Controller is ...
type Controller struct {
	Repository Repository
}

//get all books
func (c *Controller) getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.Repository.GetBooks()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

//get single book
func (c *Controller) getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	b, err := c.Repository.GetBook(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, b)
}

//create book
func (c *Controller) createBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	id, err := c.Repository.CreateBook(b)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	b.ID = id
	respondWithJSON(w, http.StatusCreated, b)
}

//update book
func (c *Controller) updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var b Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	b.ID = id

	if err := c.Repository.UpdateBook(b); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, b)
}

//delete book
func (c *Controller) deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := c.Repository.DeleteBook(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
