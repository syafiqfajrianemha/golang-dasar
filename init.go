package main

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal" // hnya ingin menjalankan function init (blank identifier)
)

func main() {
	fmt.Println(database.GetDatabase())
}
