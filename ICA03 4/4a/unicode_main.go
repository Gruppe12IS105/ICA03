package main

import (
  "./unicode"
  "fmt"
)

func main() {
  expression := "nord og sør"
  language := "jp"
  //language := "is"
  trans := unicode.Translate(expression, language)
  fmt.Printf("%s", trans)
}
