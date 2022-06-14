package main

import "fmt"

func main() {
	a := "Hello \"World\" with backslashes!"
	fmt.Println(a)

	var s1 string = "Hello"
	var s2 string = "World!"
	s3 := s1 + " " + s2
	fmt.Println(s3)
}
