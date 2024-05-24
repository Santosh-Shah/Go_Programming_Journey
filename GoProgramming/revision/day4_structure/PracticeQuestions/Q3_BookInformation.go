package main

import "fmt"

type BookInfo struct {
	title, author string
	price         int64
}

func printBookInfo(book BookInfo) {
	fmt.Println(book.title)
	fmt.Println(book.author)
	fmt.Println(book.price)
}

func main() {
	book := BookInfo{"Rich Dad Poor Dad", "Santosh", 500}
	printBookInfo(book)
}
