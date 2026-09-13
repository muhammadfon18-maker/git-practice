package main

import "fmt"

// 1) add function takes two integers as input and returns their sum
func add(a, b int) int {
	return a + b
}

// 2) divide function takes two integers as input and returns their quotient and an error if division by zero is attempted

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// 3) sum function takes a variable number of integers as input and returns their sum
// loops through the numbers using a for loop and adds each number to a total variable, which is returned at the end
// которым количество аргументов заранее неизвестно
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

// 4) anonymous function example
// anonymous function is a function without a name, which can be assigned to a variable and called later
// developer can use anonymous functions to create closures, which are functions that can capture and use variables from their surrounding scope and can be useful for creating callbacks or event handlers or for creating functions that can be passed as arguments to other functions.

var multiply = func(a, b int) int {
	return a * b
}

// 5) Closure is a function that captures and uses variables from its surrounding scope. In Go, closures are created using anonymous functions, which can be assigned to a variable and called later. Closures can be useful for creating callbacks or event handlers, or for creating functions that can be passed as arguments to other functions. They can also be used to create functions that maintain state between calls, allowing developers to create more complex and flexible programs.

//for example, a closure can be used to create a function that generates unique IDs by capturing a variable that keeps track of the last ID generated. Each time the closure is called, it can increment the variable and return a new unique ID. This allows developers to create functions that can maintain state between calls, which can be useful for creating more complex and flexible programs.

func uniqueIDGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

//6) passing functions as arguments to other functions is a powerful feature of Go that allows developers to create more flexible and reusable code. By passing functions as arguments, developers can create higher-order functions that can operate on different types of data or perform different operations based on the function passed in. This can help to reduce code duplication and make programs more modular and easier to maintain. Additionally, passing functions as arguments can also enable developers to create more complex and sophisticated programs that can adapt to changing requirements or user input.

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 7) returning

func multiplier(x int) func(int) int {
	return func(y int) int {
		return y * x
	}
}
