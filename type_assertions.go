package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt = result.(int) // panic
	// fmt.Println(resultInt)

	// baiknya gunakan switch agar tidak terjadi panic saat tidak tahu tipe datanya saat mau di konversi
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
