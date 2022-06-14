package main

import "fmt"

func main() {
	s := "Hello, World!"
	a := s[0]
	b := s[3]

	fmt.Println(s, a, string(a), b, string(b))
}
