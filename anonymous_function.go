package main

import "fmt"

// type declaration
type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Syafiq", blackList)

	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
