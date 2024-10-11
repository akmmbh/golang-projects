package main
import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)
type Movie struct{
   ID string `json:"id"`//json:id this is only a way tp say id in json format
   Isbn string `json:"isbn"`
   Title string `json:"title"`
   Director *Director `json:"director"`
//we are using director pointer not by  value as it only stores memory address not the entire directorstruct

//we can write object like that
/*
m1:Movie{
ID:"1",
Title:"Inception",
Director:&Director{
Firstname:"Christopher",
lastname:"Nolan"},}

OR

m2 := Movie{
    ID: "2",
    Title: "Untitled Movie",
    Director: nil, // no director
}
*/
}
type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie
 func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")//tell the browser that content type of comming data

	json.NewEncoder(w).Encode(movies)
	//json.NewEncoder(w) create a encoder that directly writes to w 
	//.encode the data into json format and writes the json data to w

 }
 func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	//it checks the url of the incomming requests and looks for any variables define in routeor
	//extract dynaic variable(eg {id}) from url that has been routed with gorilla 
	//it return them as map[string]string
	for index,item := range movies{
		if item.ID== params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			//this is method of deleting something using append
			//in this we are saying
			//take movies[:index] this part and 
			//append movies[index+1:]... (all element of this array) to back of this movies[:index]
			break;
		}
	}
	json.NewEncoder(w).Encode(movies)
 }
 func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
 }
 func createMovie(w http.ResponseWriter,r *http.Request){
 w.Header().Set("Content-Type","application/json")
 //browers want to create a movie so it will
 //send us data for making we can utilize this by decodin
// decode sirf r ki body karni hai 

 var movie Movie 
 _ = json.NewDecoder(r.Body).Decode(&movie)
 movie.ID =strconv.Itoa(rand.Intn(100000000))
 //this is to convert into string
 movies = append(movies,movie)
 json.NewEncoder(w).Encode(movie)
 }
 func updateMovie(w http.ResponseWriter, r *http.Request){
 //set json content type
 w.Header().Set("Content-Type","application/json")
 //params
 params:= mux.Vars(r)
 //loop over the movies, range
 //delete the movie with the id that you have sent
 //add a new movie - that we send in the body of postman
 for index, item := range movies{
	if item.ID == params["id"]{
		movies= append(movies[:index],movies[index:]...)
		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)
 movie.ID =params["id"]
 //this is to convert into string
 movies = append(movies,movie)
 json.NewEncoder(w).Encode(movie)

	}
 }
 }
func main(){
	r:=mux.NewRouter()
	movies= append(movies,Movie{ID: "1", Isbn: "438227",Title: "Movie One" , Director : &Director{Firstname:"John",Lastname: "Doe"}})//this is method to appned movies in movies
	movies= append(movies,Movie{ID: "2",Isbn: "45455",Title: "Movie Two",Director: &Director{Firstname: "steve",Lastname:"Smith"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

 fmt.Printf("starting server at port 8000\n")
 log.Fatal(http.ListenAndServe(":8000",r))//starts http server and logs any error that occur when you are runnig the http server 

}