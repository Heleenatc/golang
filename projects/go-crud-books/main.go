/* 					Routes 					Functions 				EndPoints 				Methods
----------------------------------------------------------------------------------------------------
		   		|--> GET ALL				getBooks				/books					GET 	<--|
DataBase   		|																					   |
--------   		|--> GET BY ID				getBook					/books/id				GET		<--|
	|	   		|																					   |
	|	   		|--> CREATE					createBook				/books					POST	<--|
    |      		|																					   |
    \/       	|--> UPDATE					updateBook				/books/id				PUT		<--|
BooksServer-----|																					   |
	Gorilla MUX	|--> DELETE					deleteBook				/books/id				DELETE	<--|
----------------| 																					   |
Localhost:8000	|																					   |

*/

//Note

/*
	Tried many ways to install gorilla/mux

	Finally this ran without errors
	Haven't yet checked it in code
	go mod init github.com/gorilla/mux
	go get github.com/gorilla/mux

	***Couldn't run. The following error in import
	***could not import github.com/gorilla/mux (missing metadata for import of "github.com/gorilla/mux")

	#Final solution
		D:\heleena\golang\projects\go-crud-books> go mod init main.go
		go: creating new go.mod: module main.go
		go: to add module requirements and sums:
        go mod tidy
		D:\heleena\golang\projects\go-crud-books> go mod tidy
		go: finding module for package github.com/gorilla/mux
		go: found github.com/gorilla/mux in github.com/gorilla/mux v1.8.1

	### use "go mod init" with our current module name - not with the name of the package we are importing
	### after this run "go mod tidy", which will include all the missing dependencies required for our project

	-------------------------------------
	only getBooks and getBook are working while checking with postman
	405 - method not allowed error in case of Create and Update
	Delete option - data appears to be deleted but its still there in the list
	***
	May be because of the version updates in go or gorilla/mux
	should revisit these issues once I get familiar with go

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id,omitempty"`
	Isbn   string  `json:"isbn,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author *Author `json:"author,omitempty"`
}

type Author struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break

		}
	}
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "12345", Title: "Indian Polity", Author: &Author{Firstname: "Lakshmi", Lastname: "Kant"}})
	books = append(books, Book{ID: "2", Isbn: "23456", Title: "The Story of My Experiments with Truth", Author: &Author{Firstname: "Mahatma", Lastname: "Gandhi"}})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books/{id}", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Starting server at port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))

}
