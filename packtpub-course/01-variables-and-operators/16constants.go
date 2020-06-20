package main

import (
	"fmt"
)

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit

// OR group by relationship:
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

//Constants are like variables, but you can't change their initial value
func main() {
	//Initialize our cache by creating a map:
	cache = make(map[string]string)
	//Add a book to the cache:
	SetBook("1234-5678", "Get Ready To Go")
	//Add a CD to the cache:
	SetCD("1234-5678", "Get Ready To Go Audio Book")
	//Get and print that Book from the cache:
	fmt.Println("Book :", GetBook("1234-5678"))
	//Get and print that CD from the cache:
	fmt.Println("CD :", GetCD("1234-5678"))
}
