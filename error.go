package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi Dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, error := pembagian(100, 10)
	if error == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", error.Error())
	}
}
