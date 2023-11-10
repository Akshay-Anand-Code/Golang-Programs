package main


import "time"

type Book struct {
	title string
	author string
	numPages int

	isSaved bool
	savedAT time.Time
}

func saveBook(book *Book) {
   book.isSaved = true
   book.savedAT = time.Now()
}

func main() {
	
}