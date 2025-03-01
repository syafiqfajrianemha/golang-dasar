package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Syafiq", "Fajrian", "Emha"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	persons := map[string]string{
		"name":    "Syafiq",
		"address": "Purbalingga",
	}

	for _, person := range persons {
		fmt.Println(person)
	}
}
