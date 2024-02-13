package main

import "fmt"


func Greet(name string) string {
  // fill in solution here
  return "Hello, "+name+" how are you doing today?"
}

func main() {
	name := "Osman"
	fmt.Println(Greet(name))
}
