package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := a + b
	fmt.Println(c)

	i := 10
	fmt.Println(i)
	i += 10 // i = i + 10
	fmt.Println(i)
	i += 5 // i = i + 5
	fmt.Println(i)

	j := 1
	j++
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
}
