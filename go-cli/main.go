// package main

// import (
// 	"fmt"
	
// )

// func main() {
//    var x int
//    var y int
//    var operator string
//    fmt.Print("Enter a number: ")
//    _, err := fmt.Scan(&x)
   
//    if err != nil {
// 	   fmt.Println("Invalid input:" , err)
// 	   return
//    }

//    fmt.Print("Enter an operator (+, -, *, /): ")
//    _, err = fmt.Scan(&operator)
//    if err != nil {
// 	   fmt.Println("Invalid input:" , err)
// 	   return
//    }

//    fmt.Print("Enter another number: ")
//    _, err = fmt.Scan(&y)
//    if err != nil {
// 	   fmt.Println("Invalid input:" , err)
// 	   return
//    }

//    result := calculateFigures(x, y, operator)
//    fmt.Printf("The result is: %d\n", result)
// }

// func add(x int, y int) int {
//     return x + y
// }

// func subtract(x int, y int) int {
// 	return x - y
// }

// func multiply(x int, y int) int {
// 	return x * y
// }

// func divide(x int, y int) (int, error) {
// 	if y == 0 {
// 		return 0, fmt.Errorf("division by zero is not allowed")
// 	}
//       return x/y , nil
// }

// func calculateFigures(x int, y int, operator string) int{
// 	switch operator {
// 	case "+":
// 		return add(x, y)
// 	case "-":
// 		return subtract(x, y)
// 	case "*":
// 		return multiply(x, y)
// 	case "/":
// 		result, err := divide(x, y)
// 		if err != nil {
// 			fmt.Println(err)
// 			return 0
// 		}
// 		return result
// 	default:
// 		fmt.Println("Invalid operator")
// 		return 0
// 	}
// }


package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
)

func main() {
   numberPrompt := bufio.NewReader(os.Stdin)

   fmt.Println("Enter a number: ")
   numberInput, err := numberPrompt.ReadString('\n')
 
   number1, err := strconv.ParseFloat(strings.TrimSpace(numberInput), 64)
   if err != nil {
	   fmt.Println("Invalid input:" , err)
	   return
   }
   fmt.Println("You entered number:", number1)
   


   fmt.Println("Enter an operator (+, -, *, /): ")
   operator, err := numberPrompt.ReadString('\n')
   operator = strings.TrimSpace(operator)
   if err != nil {
	   fmt.Println("Invalid input:" , err)
	   return
   }	
   
   fmt.Println("Enter another number: ")
   numberInput2, err := numberPrompt.ReadString('\n')

   number2, err := strconv.ParseFloat(strings.TrimSpace(numberInput2), 64)
   if err != nil {
	   fmt.Println("Invalid input:" , err)
	   return
   }
   fmt.Println("You entered number:", number2)

   result, err := calculateResult(number1, number2, operator)
   if err != nil {
	   fmt.Println("Error:", err)
	   return
   }
   fmt.Printf("The result is: %v\n", result)
}

func calculateResult(x float64, y float64 ,operator string) (float64, error) {
	switch operator{
	   case "+":
		return x+y, nil
       case "-":
		return x-y, nil
	   case "*":
		return x*y, nil
	   case "/":
		if y == 0{
			return 0, fmt.Errorf("Division by zero is not allowed")
		}
		return x/y, nil
	default:
		return 0, fmt.Errorf("Invalid operator")
	}
}
