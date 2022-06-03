package router

import (
	"github.com/gorilla/mux"
	"github.com/hiteshpattanayak-tw/golangtraining/new_books_api/internal/app/handlers"
	"github.com/hiteshpattanayak-tw/golangtraining/new_books_api/internal/app/services"
)

func ProvideRouter(bookService services.BooksService) *mux.Router {
	r := mux.NewRouter()

	healthHandler := handlers.HealthHandler{}
	booksHandler := handlers.GetNewBooksHandler(bookService)

	r.HandleFunc("/health", healthHandler.HandlerFunc).Methods("GET")
	r.HandleFunc("/books/all", booksHandler.BooksHandler).Methods("GET")
	r.HandleFunc("/books", booksHandler.UpsertBookHandler).Methods("POST", "PUT")
	r.HandleFunc("/books/{isbn:[0-9]+}", booksHandler.AddOrRemoveBookHandler).Methods("GET", "DELETE")

	return r
}