package main

import(
"fmt"
"log"
"encoding/json"
"math/rand"
"net/http"
"github.com/gorilla/mux")


type movie struct{
	id string `json: "id"`
	isbn string `json: "isbn"`
	title string `json: "title"`
	director *director `json: "director"`

}
type director struct{
	firstName string `json: "firstName"`
	lastName string `json: "lastName"`

}
 var movies []movie

 func main(){
	 r := mux.NewRouter()
	 r.HandleFunc("/movies",getMovies).Methods("GET")
	 r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	 r.HandleFunc("/movies",createMovie).Methods("POST")
	 r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	 r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
 }