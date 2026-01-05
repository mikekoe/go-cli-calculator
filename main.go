package main

import (
	"fmt"
)

func main() {
   var x int
   var y int
   var operator string
   fmt.Print("Enter a number: ")
   _, err := fmt.Scan(&x)
   if err != nil {
	   fmt.Println("Invalid input:" , err)
	   return
   }

   fmt.Print("Enter an operator (+, -, *, /): ")
   _, err = fmt.Scan(&operator)
   if err != nil {
	   fmt.Println("Invalid input:" , err)
	   return
   }

   fmt.Print("Enter another number: ")
   _, err = fmt.Scan(&y)
   if err != nil {
	   fmt.Println("Invalid input:" , err)
	   return
   }

   result := calculateFigures(x, y, operator)
   fmt.Printf("The result is: %d\n", result)
}

func add(x int, y int) int {
    return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
      return x/y , nil
}

func calculateFigures(x int, y int, operator string) int{
	switch operator {
	case "+":
		return add(x, y)
	case "-":
		return subtract(x, y)
	case "*":
		return multiply(x, y)
	case "/":
		result, err := divide(x, y)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return result
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}