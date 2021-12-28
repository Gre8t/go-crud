package main

import(
"fmt"
"log"
"encoding/json"
"math/rand"
"net/http"
"github.com/gorilla/mux")


type Movie struct{
	ID string `json: "ID"`
	Isbn string `json: "Isbn"`
	Title string `json: "Title"`
	Director *Director `json: "Director"`

}
type Director struct{
	FirstName string `json: "FirstName"`
	LastName string `json: "LastName"`

}
 var movies []Movie

 func getMovies(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
 }

 func getMovie(){

 }
func createMovie(){

}
func updateMovie(){

}

func deleteMovie(){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
		break
	}


}
 func main(){
	 r := mux.NewRouter()
	 movies = append(movies, Movie{ID: "1", Isbn: "438227", Title : "movie one", Director : &Director{FirstName: "john", LastName: "doe"}})
	 movies = append(movies, Movie{ID: '2'})
	 r.HandleFunc("/movies",getMovies).Methods("GET")
	 r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	 r.HandleFunc("/movies",createMovie).Methods("POST")
	 r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	 r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	 fmt.Printf("starting server at port 8080\n")
	 log.Fatal(http.ListenAndServe(":8080", r))

 }