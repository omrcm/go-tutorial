package main

import "fmt"

func main() {
	/*
		uint8       unsigned  8-bit integers (0 to 255)
		uint16      unsigned 16-bit integers (0 to 65535)
		uint32      unsigned 32-bit integers (0 to 4294967295)
		uint64      unsigned 64-bit integers (0 to 18446744073709551615)
		int8        signed  8-bit integers (-128 to 127)
		int16       signed 16-bit integers (-32768 to 32767)
		int32       signed 32-bit integers (-2147483648 to 2147483647)
		int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

		float32     IEEE-754 32-bit floating-point numbers
		float64     IEEE-754 64-bit floating-point numbers
		complex64   complex numbers with float32 real and imaginary parts
		complex128  complex numbers with float64 real and imaginary parts

		byte        alias for uint8
		rune        alias for int32

		Please check the images uploaded in this folder.
		The most imported thing to note in Go is;
		* You can't use variables of different numeric types together with out explicitly converting
		  their types to match. IF we want to sum a int8 with float 32 Go will give us;
		  "mismatch types int8 and float32". We have to fix it with type conversion
	*/
	var i int8 = 8
	var j int8 = 10
	fmt.Println(i + j)

	var a int32 = 100
	var b int32 = 300
	fmt.Println(a + b)

	var c float32 = 10.3
	var d float32 = 0.8
	fmt.Println(c + d)

	var e uint = 1
	var f uint = 2
	fmt.Println(e + f)
}
