package main

import "fmt"

func main() {
	// Testing the add function
	result := add(5, 10)
	fmt.Println("The result is:", result)



	// Testing the divide function

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The result is:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The result is:", result)
	}

	//Variadic function

	fmt.Println("Sum of 1, 2, 3, 4, 5 is:", sum(1, 2, 3, 4, 5))

	// Testing the anonymous function
	result = multiply(5, 10)
	fmt.Println("The result is:", result)

	//go mod is a dependency management system for Go programming language. It allows developers to manage the dependencies of their Go projects in a more efficient and organized way. With go mod, developers can specify the required versions of external packages and libraries, and the tool will automatically download and manage those dependencies for them. This helps to ensure that the project is using the correct versions of its dependencies, which can help to prevent compatibility issues and other problems that can arise when working with external packages. Additionally, go mod also provides features for versioning and publishing modules, making it easier for developers to share their code with others.

	// go mod and go sum difference is that go mod is a tool for managing dependencies in Go projects, while go.sum is a file that contains cryptographic hashes of the dependencies used in a project. The go.sum file is automatically generated and updated by the go mod tool whenever a new dependency is added or an existing dependency is updated. It serves as a way to verify the integrity of the dependencies used in a project, ensuring that they have not been tampered with or modified in any way. In other words, go mod manages the dependencies, while go.sum verifies their integrity.

	closure := uniqueIDGenerator()
	fmt.Println("Unique ID 1:", closure())
	fmt.Println("Unique ID 2:", closure())
	fmt.Println("Unique ID 3:", closure())

	result = applyOperation(5, 10, add)
	fmt.Println("The result of applying the add operation is:", result)

	result = applyOperation(5, 10, multiply)
	fmt.Println("The result of applying the multiply operation is:", result)


}
