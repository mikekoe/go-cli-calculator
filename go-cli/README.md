# CLI Calculator Application (Go)

A simple **Command-Line Calculator** built with Go that performs basic arithmetic operations based on user input.

This project demonstrates **user input handling**, **string parsing**, **error handling**, **functions**, and **control flow** in Go.

---

## Features

- Addition
- Subtraction
- Multiplication
- Division (with zero-division protection)
- Reads input from standard input (CLI)
- Graceful error handling for invalid inputs

---

## How It Works

1. The program prompts the user to enter the **first number**
2. The user selects an **operator** (`+`, `-`, `*`, `/`)
3. The user enters the **second number**
4. The program calculates and prints the result
5. Invalid inputs or operations are handled safely

---

## Project Structure

```text
cli-calculator/
├── main.go
└── README.md
```
---

## Requirements
- Go installed (Go 1.18+ recommended)
### Verify the installation
```bash
go version
```

---

## Run the program on your PC
### From the project directory, run:
```bash
go run main.go
```
---

## Example Usage
### Input the following on your CLI
```text
Enter a number:
10
Enter an operator (+, -, *, /):
*
Enter another number:
5
```
---

## Output
```text
You entered number: 10
You entered number: 5
The result is: 50
```
