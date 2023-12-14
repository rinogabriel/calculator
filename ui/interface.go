package ui

import (
	"calculator_golang/calculator"
	"fmt"
)

func Interface() {
	for {
		var num1, num2 int

		fmt.Print("Input your first number: ")
		_, err := fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("Error occurred while reading input:", err)
			return
		}

		fmt.Print("Input your second number: ")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("Error occurred while reading input:", err)
			return
		}

		fmt.Println("Choose your method : ")
		fmt.Println("1. Addition ")
		fmt.Println("2. Subtraction ")
		fmt.Println("3. Multiply ")
		fmt.Println("4. Division ")
		fmt.Println("5. PowersOf ")
		fmt.Println("6. Exit")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "1":
			hasil := calculator.Addition(num1, num2)
			fmt.Println("Result:", hasil)
		case "2":
			hasil := calculator.Substarction(num1, num2)
			fmt.Println("Result:", hasil)
		case "3":
			hasil := calculator.Multiply(num1, num2)
			fmt.Println("Result:", hasil)
		case "4":
			hasil := calculator.Division(num1, num2)
			fmt.Println("Result:", hasil)
		case "5":
			hasil := calculator.PowerOf(num1, num2)
			fmt.Println("Result:", hasil)
		case "6":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid input. Please enter a valid option.")
		}
	}
}
