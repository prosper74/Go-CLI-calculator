package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//Checking if the user has provided any input args
	if len(os.Args) < 4 {
		fmt.Println("Please provide input args for calculation")
		os.Exit(1)
	}

	//Checking if the user has provided valid input strings
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Please provide valid input strings")
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Please provide valid input strings")
		os.Exit(1)
	}

	// Calculating the result
	switch os.Args[2] {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Please provide valid operator")
		os.Exit(1)
	}
}
