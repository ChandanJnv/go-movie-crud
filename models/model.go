package models

// Movie represents a movie with an ID, ISBN, title, and director information.
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director represents a movie director with a first and last name.
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// movies is a slice that holds a list of movies.
var MovieList []Movie

// Initialize dummy data for movies.
func InitializeMovies() {
	MovieList = append(MovieList, Movie{ID: "1", Isbn: "342434", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Snow"}})
	MovieList = append(MovieList, Movie{ID: "2", Isbn: "432433", Title: "Movie Two", Director: &Director{FirstName: "Bill", LastName: "Smith"}})
	MovieList = append(MovieList, Movie{ID: "3", Isbn: "523423", Title: "Movie Three", Director: &Director{FirstName: "Anna", LastName: "Taylor"}})
	MovieList = append(MovieList, Movie{ID: "4", Isbn: "634234", Title: "Movie Four", Director: &Director{FirstName: "Tom", LastName: "Hanks"}})
	MovieList = append(MovieList, Movie{ID: "5", Isbn: "734535", Title: "Movie Five", Director: &Director{FirstName: "Emma", LastName: "Stone"}})
}
