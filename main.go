package main

import "fmt"

/*

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




	id := generateId()
	fmt.Println("Generated ID 1:", id())
	fmt.Println("Generated ID 2:", id())
	fmt.Println("Generated ID 3:", id())


	greaterThan := greaterThan(10)
	fmt.Println("Is 15 greater than 10?", greaterThan(15)) // true
	fmt.Println("Is 5 greater than 10?", greaterThan(5)) // false

//-----------------------------------------------------------------------------------

	multiplierBy2 := multiplier(2)
	fmt.Println("2 multiplied by 5 is:", multiplierBy2(5)) // 10
	fmt.Println("2 multiplied by 10 is:", multiplierBy2(10)) // 20

	multiplierBy3 := multiplier(3)
	fmt.Println("3 multiplied by 5 is:", multiplierBy3(5)) // 15
	fmt.Println("3 multiplied by 10 is:", multiplierBy3(10)) // 30
}

*/
type Vertext struct {
	lat, long float64
}

/* var m = map[string]Vertext{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
 */
//Map literals 

func main() {

	m := make(map[string]Vertext)

fmt.Println(m["Bell Labs"]) 
fmt.Println(m["Google"])

//Insert or update an element in map m:
m["New York"] = Vertext{40.7128, -74.0060}

//retrieve the value associated with the key "New York":
x := m["New York"]
fmt.Println(x) // Output: {40.7128 -74.006}

//Delete an element from map m:
delete(m, "New York")

//Check if a key exists in map m:
value, ok := m["New York"]
if ok {
	fmt.Println("Key exists with value:", value)
} else {
	fmt.Println("Key does not exist")


}
}
// ok funcion in map is used to check if a key exists in the map. When you retrieve a value from a map using a key, Go returns two values: the value associated with the key and a boolean indicating whether the key exists in the map. If the key exists, the boolean will be true, and if it does not exist, the boolean will be false. This allows you to safely check for the existence of a key before using its associated value, preventing potential runtime errors.

