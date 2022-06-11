package main

import "fmt"

func main() {
	var i int8 = 20
	var f float32 = 5.6
	fmt.Println(i + int8(f+1.9))

	var a int8 = 20
	var b int32 = 40
	fmt.Println(a + int8(b))
}
