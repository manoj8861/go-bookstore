package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manoj8861/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.ResgisterBookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("server running on http://localhost:9010")
	if err := http.ListenAndServe("localhost:9010", r); err != nil {
		log.Fatal(err)
	}

}
