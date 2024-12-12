package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = 5
	var d = 2
	var result = a + b - c*d
	fmt.Println(result)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = 1 + 5
	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}
