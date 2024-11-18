package main

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func NewBook(Title string, Author string, Year int, Genre string) *Book {
	return &Book{Title: Title, Author: Author, Year: Year, Genre: Genre}
}

func main() {

}
