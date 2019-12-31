package main

import "fmt"

func main() {
  var publisher string
  var writer string
  var artist string
  var title string
  var year int
  var pageNumber int
  var grade float32
  
  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "and is", pageNumber, "pages long with a condition of", grade, "out of 10")
  
}