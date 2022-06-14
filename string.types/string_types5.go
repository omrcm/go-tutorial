package main

import "fmt"

func main() {
	/*
		substring a string in Go
	*/
	s := "Hello, World!"
	a := s[1:8]
	b := s[4:12]
	c := s[:3]
	d := s[5:]

	fmt.Println(a, b, c, d)
	fmt.Println(b)
}
