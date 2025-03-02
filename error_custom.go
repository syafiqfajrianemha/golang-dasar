package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "Syafiq" {
		return &notFoundError{"Data not found"}
	}

	// OK

	return nil
}

func main() {
	err := saveData("Syafiq", nil)
	if err != nil {
		// terjadi error
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error", finalError.Error())
		default:
			fmt.Println("unknown", err.Error())
		}
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
