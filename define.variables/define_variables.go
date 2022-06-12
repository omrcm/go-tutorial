package main

import "fmt"

func main() {

	/*
		There are a few important issues we need to know before moving on to the definitions.
		1) When we define a variable we have to use it.
		2) Default all parameters will be assiged 0

		Declaring with var keyword.
		We don't have to define the type of the variable.
		After the first assign Go will automatically detect the type
	*/
	var i = 20
	fmt.Println(i)

	/*
		Declaring with ":=" operator
	*/
	j := 30
	fmt.Println(j)

	/*
		Ok! Why we need the "var" keyword if we can define it easly with ":=".
		There is some reasons for that. But the most 2 reason is;

		1)  You cannot use a shorthand syntax outside of a function.
			So if you declare a package level variable, you must use the VAR keyword,
			although you can leave out the type if the type can be inferred from the right hand side.

		2)  We don't always want to use the type the compiler differs from the right hand side of the equal sign

		This is a example for why we use the "var" keyword to define a variable
	*/
	var c int
	fmt.Println(c)
	c = 10
	fmt.Println(c)

	/*
		When we run above code, it'll print first 0 and then 10.

		If you declare a variable, you must use it,any variables that are declared but never actually used
		will trigger a compilation error
	*/

	var s = 10

	/*
		if we remove the code on line 54 Go will give us compilation error
	*/
	fmt.Println(s)

}
