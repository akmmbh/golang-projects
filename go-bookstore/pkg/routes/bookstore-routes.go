package routes

import (
	"github.com/gorilla/mux"
	"github.com/akmmbh/go-bookstore/pkg/controlllers"
)
var RegisterBookStoreRoutes =func(router *mux.Router){
	router.HandleFunc("/book/",controlllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controlllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controlllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controlllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controlllers.DeleteBook).Methods("DELETE")

}