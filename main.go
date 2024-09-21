package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("EndPoint hit success")
	fmt.Fprintf(w, "Welcome to homepage")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Below is the list of all the Products \n")
	log.Println("EndPoint Hit: returnAllproducts")
	json.NewEncoder(w).Encode(Products)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	log.Println(r.URL.Path)
	// key := r.URL.Path[len("/product/"):]
	intkey, _ := strconv.Atoi(key)
	for _, product := range Products {
		if product.Id == intkey {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getProductById)
	myRouter.HandleFunc("/", homepage)
	error := http.ListenAndServe("localhost:8080", myRouter)
	if error != nil {
		log.Panic(error)
	}
}

func main() {

	Products = []Product{
		{Id: 1, Name: "Table", Quantity: 10, Price: 200.00},
		{Id: 2, Name: "Chair", Quantity: 15, Price: 100.00},
	}
	handleRequests()
}
