package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	// yang huruf depannya besar/kapital bisa diakses dari luar package
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
}
