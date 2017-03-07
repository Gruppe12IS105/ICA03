package main

import (
	"fmt"
	"os"

	"./unicode"
)

func main() {
	//Skriv jp eller is etter 'go run unicode_main.go'
	language := os.Args[1]
	trans := unicode.Translate(language)
	fmt.Printf("%s", trans)
}
