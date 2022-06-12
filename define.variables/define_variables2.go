package main

import "fmt"

func main() {
	/*
		In this code we have defined a variable named j and have set a default value.
		But in line 12 we have set a new valueto the same variable.
		When we print the "j" variable it will print the latest value
	*/
	j := 30
	j = 20
	fmt.Println(j)
}
