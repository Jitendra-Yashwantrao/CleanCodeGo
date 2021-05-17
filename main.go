package main

import (
	"cleancode/movierental"
	"fmt"
)

func main() {

	fmt.Println("Welcome to clean code refactoring")
	customer := movierental.Customer{}
	statement := customer.Statement()

	fmt.Println("Customer Statement", statement)

}
