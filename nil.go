package main

import "fmt"

// nil hanya untuk interface, function, map, slice, pointer, dan channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println("Hallo", data["name"])
	}
}
