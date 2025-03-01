package main

import "fmt"

// defer, panic, & recover itu seperti try catch

func endAppTo() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runAppTo(error bool) {
	defer endAppTo()
	if error {
		panic("Ups Error")
	}
}

func main() {
	runAppTo(true)
	fmt.Println("Syafiq Fajrian Emha")
}
