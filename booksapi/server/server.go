package server

import (
	"fmt"
	"github.com/gdtrivedi/gopractice/booksapi"
	"github.com/gdtrivedi/gopractice/booksapi/dto"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//slice books
var books []dto.Book

func NewMuxServer() {
	// Init router
	r := mux.NewRouter()

	// set books mock data
	books = append(books, dto.Book{ID: "1", Isbn: "111", Title: "Book One", Author: dto.Author{Firstname: "G", Lastname: "T"}})
	books = append(books, dto.Book{ID: "2", Isbn: "222", Title: "Book Two", Author: dto.Author{Firstname: "P", Lastname: "T"}})
	books = append(books, dto.Book{ID: "3", Isbn: "333", Title: "Book Three", Author: dto.Author{Firstname: "D", Lastname: "T"}})

	// Init books api
	booksapi.InitAPI(books)

	// Route handlers
	r.HandleFunc("/api/books", booksapi.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", booksapi.GetBook).Methods("GET")
	r.HandleFunc("/api/books", booksapi.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", booksapi.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", booksapi.DeleteBook).Methods("DELETE")

	fmt.Println("Listening on port 7000")
	log.Fatal(http.ListenAndServe(":7000", r))
}
