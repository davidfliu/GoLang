// Practicing go structures.

package main 

import "fmt"

type Books struct {

	title string
	author string
	subject string
}

type Movies struct {

	releaseYear int
	title string
	director string
}

func main() {

	var Book Books
	var Movie Movies

	Book.title = "Little Red Riding Hood"
	Book.author = "Charlie Brown"
	Book.subject = "Fiction"

	Movie.releaseYear = 2000
	Movie.title = "Interstellar"
	Movie.director = "James Cameron"

	fmt.Printf("Book title is: %s\n", Book.title)
	fmt.Printf("Book subject is: %s\n\n", Book.subject)

	fmt.Printf("Movie release year is: %d\n", Movie.releaseYear)
	fmt.Printf("Movie title is: %s\n", Movie.title)

}