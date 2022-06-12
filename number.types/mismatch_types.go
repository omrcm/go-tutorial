package main

import "fmt"

func main() {
	/*
		The most imported thing to note in Go is;
			* You can't use variables of different numeric types together with out explicitly converting
			their types to match. IF we want to sum a int8 with float 32 Go will give us;
			"mismatch types int8 and float32". We have to fix it with type conversion
	*/

	i := 10.3
	var j float32 = 2.1
	fmt.Println(float32(i) + j)

	var a int32 = 8
	var b int64 = 12
	fmt.Println(a + int32(b))

	var c float64 = 19.54
	var d int32 = 11
	fmt.Println(c + float64(d))
	//it will not round the number!
	fmt.Println(int32(c) + d)
}
