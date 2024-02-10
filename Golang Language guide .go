package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		// Display the menu
		fmt.Println("Welcome to GoLang Tutorial:")
		fmt.Println("1. Syntax")
		fmt.Println("2. History of Go Language")
		fmt.Println("3. Variables")
		fmt.Println("4. Functions")
		fmt.Println("5. Loops")
		fmt.Println("6. Arrays and Slices")
		fmt.Println("7. Pointers")
		fmt.Println("8. Structs")
		fmt.Println("9. Channels")
		fmt.Println("10. Multi-threading")
		fmt.Println("11. Backend Development")
		fmt.Println("12. Exit")

		// Ask for user input
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Perform the action based on user's choice
		switch choice {
		case 1:
			showSyntax()
		case 2:
			showHistory()
		case 3:
			showVariables()
		case 4:
			showFunctions()
		case 5:
			showLoops()
		case 6:
			showArraysAndSlices()
		case 7:
			showPointers()
		case 8:
			showStructs()
		case 9:
			showChannels()
		case 10:
			showMultiThreading()
		case 11:
			showBackendDevelopment()
		case 12:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please choose a number between 1 and 12.")
		}
	}
}

func showSyntax() {
	fmt.Println("Go syntax is designed to be simple and readable.")
	fmt.Println("It uses packages to organize code, and statements are terminated with a semicolon.")
	fmt.Println("Example:")
	fmt.Println("package main")
	fmt.Println("import \"fmt\"")
	fmt.Println("func main() {")
	fmt.Println("    fmt.Println(\"Hello, world!\")")
	fmt.Println("}")

	fmt.Println("\nIn Go, the syntax is straightforward and aims for readability.")
	fmt.Println("1. `package main`: Every Go program starts with a package declaration. The `main` package is the entry point for executable programs. Other packages are used for reusable code.")
	fmt.Println("2. `import \"fmt\"`: This line imports the 'fmt' package, which provides functions for formatting input and output.")
	fmt.Println("3. `func main() { ... }`: This defines the main function. It's the entry point of the program, and execution starts from here.")
	fmt.Println("4. `fmt.Println(\"Hello, world!\")`: Inside the main function, we call the 'Println' function from the 'fmt' package to print 'Hello, world!' to the console.")
	fmt.Println("5. Statements in Go do not require a semicolon at the end, but they can be used if you want to separate multiple statements on the same line.")
}

func showHistory() {
	fmt.Println("Go, also known as Golang, was created by Google engineers Robert Griesemer, Rob Pike, and Ken Thompson.")
	fmt.Println("It was first announced in November 2009 and reached version 1.0 in March 2012.")
	fmt.Println("Go was designed to be a statically typed, compiled language with a focus on simplicity, efficiency, and concurrency.")
}

func showVariables() {
	fmt.Println("Variables in Go are used to store data. They are declared using the 'var' keyword.")
	fmt.Println("Example:")
	fmt.Println("var x int = 5")

	fmt.Println("\nIn Go, variables are containers for storing data values.")
	fmt.Println("1. To declare a variable, you use the 'var' keyword followed by the variable name.")
	fmt.Println("2. Then you specify the variable's type using the colon-equals operator ':='. You can also explicitly specify the type using the 'var' keyword followed by the type.")
	fmt.Println("3. Finally, you assign a value to the variable using the assignment operator '='.")
	fmt.Println("4. In the example provided, we declare a variable 'x' of type 'int' and initialize it with the value 5.")
}

func showFunctions() {
	fmt.Println("Functions in Go are blocks of code that perform a specific task.")
	fmt.Println("Example:")
	fmt.Println("func add(a, b int) int {")
	fmt.Println("    return a + b")
	fmt.Println("}")

	fmt.Println("\nIn Go, functions are defined using the 'func' keyword followed by the function name.")
	fmt.Println("1. The function name is followed by a list of parameters, enclosed in parentheses. Each parameter consists of a name and a type.")
	fmt.Println("2. If a function takes multiple parameters of the same type, you can specify the type only once followed by the parameter names.")
	fmt.Println("3. After the parameter list, you specify the return type of the function. If the function does not return any value, you can omit the return type.")
	fmt.Println("4. Inside the function body, you write the code that performs the desired task.")
	fmt.Println("5. You use the 'return' statement to return a value from the function.")

	fmt.Println("\nExamples of different types of functions:")
	fmt.Println("- Function with parameters and return value:")
	fmt.Println("func add(a, b int) int {")
	fmt.Println("    return a + b")
	fmt.Println("}")

	fmt.Println("\n- Function with multiple return values:")
	fmt.Println("func swap(x, y string) (string, string) {")
	fmt.Println("    return y, x")
	fmt.Println("}")

	fmt.Println("\n- Function with variadic parameters:")
	fmt.Println("func sum(nums ...int) int {")
	fmt.Println("    total := 0")
	fmt.Println("    for _, num := range nums {")
	fmt.Println("        total += num")
	fmt.Println("    }")
	fmt.Println("    return total")
	fmt.Println("}")

	fmt.Println("\n- Function with named return values:")
	fmt.Println("func divide(x, y float64) (result float64, err error) {")
	fmt.Println("    if y == 0 {")
	fmt.Println("        return 0, errors.New(\"division by zero\")")
	fmt.Println("    }")
	fmt.Println("    return x / y, nil")
	fmt.Println("}")

	fmt.Println("\nThese examples demonstrate different ways of defining functions in Go, including functions with parameters and return values, functions with multiple return values, variadic functions, and functions with named return values.")
}
func showLoops() {
	fmt.Println("Loops in Go are used to execute a block of code repeatedly.")
	fmt.Println("Example:")
	fmt.Println("for i := 0; i < 5; i++ {")
	fmt.Println("    fmt.Println(i)")
	fmt.Println("}")

	fmt.Println("\nIn Go, loops are controlled using three primary loop constructs: 'for', 'while', and 'range'.")
	fmt.Println("1. The 'for' loop is the most common loop in Go, and it's similar to the 'for' loop in many other programming languages.")
	fmt.Println("2. The 'for' loop consists of three parts: initialization, condition, and post statement.")
	fmt.Println("3. The initialization statement is executed before the loop starts. It typically initializes a loop control variable.")
	fmt.Println("4. The condition is evaluated before each iteration. If it evaluates to true, the loop continues; otherwise, the loop terminates.")
	fmt.Println("5. The post statement is executed after each iteration. It's commonly used to increment or decrement loop control variables.")
	fmt.Println("6. Inside the loop body, you write the code that you want to execute repeatedly.")

	fmt.Println("\nExamples of different types of loops:")
	fmt.Println("- Standard 'for' loop:")
	fmt.Println("for i := 0; i < 5; i++ {")
	fmt.Println("    fmt.Println(i)")
	fmt.Println("}")

	fmt.Println("\n- 'while' loop (using 'for' without initialization and post statement):")
	fmt.Println("i := 0")
	fmt.Println("for i < 5 {")
	fmt.Println("    fmt.Println(i)")
	fmt.Println("    i++")
	fmt.Println("}")

	fmt.Println("\n- Looping over collections using 'range':")
	fmt.Println("nums := []int{1, 2, 3, 4, 5}")
	fmt.Println("for index, value := range nums {")
	fmt.Println("    fmt.Printf(\"Index: %d, Value: %d\\n\", index, value)")
	fmt.Println("}")

	fmt.Println("\nThese examples demonstrate different ways of using loops in Go, including standard 'for' loops, 'while' loops, and looping over collections using 'range'.")
}

func showArraysAndSlices() {
	fmt.Println("Arrays and slices are used to store ordered collections of items in Go.")
	fmt.Println("Example:")
	fmt.Println("var arr [3]int = [3]int{1, 2, 3}")
	fmt.Println("slice := []int{1, 2, 3}")

	fmt.Println("\nIn Go, arrays and slices are used to store ordered collections of items.")
	fmt.Println("1. An array is a fixed-size collection of elements of the same type. Once an array is declared, its size cannot be changed.")
	fmt.Println("2. To declare an array, you specify the type of its elements and the number of elements enclosed in square brackets [].")
	fmt.Println("3. You can initialize the array with values using the curly braces {}.")
	fmt.Println("4. A slice, on the other hand, is a flexible, dynamic view of an array. It's like a window that specifies a range of elements from an array.")
	fmt.Println("5. Slices are declared using square brackets [] without specifying a size. They are created using the 'make' function or by slicing an existing array or slice.")

	fmt.Println("\nExamples of arrays and slices:")
	fmt.Println("- Array:")
	fmt.Println("var arr [3]int = [3]int{1, 2, 3}")

	fmt.Println("\n- Slice:")
	fmt.Println("slice := []int{1, 2, 3}")
	fmt.Println("or")
	fmt.Println("arr := [5]int{1, 2, 3, 4, 5}")
	fmt.Println("slice := arr[1:4] // creates a slice from index 1 to 3 of the array")

	fmt.Println("\nIn the example provided, we declare an array 'arr' of type '[3]int' with three elements, and a slice 'slice' of type '[]int' with three elements.")
	fmt.Println("The slice 'slice' is created using a slice literal and does not specify a size.")
	fmt.Println("Alternatively, a slice can be created from an existing array or slice by specifying the start and end index.")
}

func showPointers() {
	fmt.Println("Pointers in Go are variables that store the memory address of another variable.")
	fmt.Println("Example:")
	fmt.Println("var x int = 5")
	fmt.Println("var ptr *int = &x")

	fmt.Println("\nIn Go, pointers are variables that store the memory address of another variable.")
	fmt.Println("1. To declare a pointer variable, you use the asterisk '*' followed by the type of the variable it points to.")
	fmt.Println("2. You can obtain the memory address of a variable using the address-of operator '&'.")
	fmt.Println("3. When you declare a pointer variable, you initialize it with the address of the variable it points to.")

	fmt.Println("\nExamples of pointers:")
	fmt.Println("- Declaration and initialization:")
	fmt.Println("var x int = 5")
	fmt.Println("var ptr *int = &x")

	fmt.Println("\n- Dereferencing a pointer:")
	fmt.Println("fmt.Println(\"Value of x: \", x) // prints the value of x")
	fmt.Println("fmt.Println(\"Address of x: \", &x) // prints the address of x")
	fmt.Println("fmt.Println(\"Value pointed by ptr: \", *ptr) // prints the value pointed by ptr")

	fmt.Println("\nIn the example provided, we declare an integer variable 'x' with the value 5.")
	fmt.Println("Then, we declare a pointer variable 'ptr' of type '*int' and initialize it with the address of 'x' using the address-of operator '&'.")
	fmt.Println("The 'ptr' variable now holds the memory address of 'x', allowing us to indirectly access and modify the value of 'x' through 'ptr'.")
}

func showStructs() {
	fmt.Println("Structs in Go are used to create user-defined types.")
	fmt.Println("Example:")
	fmt.Println("type Person struct {")
	fmt.Println("    Name string")
	fmt.Println("    Age  int")
	fmt.Println("}")

	fmt.Println("\nIn Go, structs are used to create user-defined types that group together different fields of data.")
	fmt.Println("1. To define a struct, you use the 'type' keyword followed by the name of the struct and the keyword 'struct'.")
	fmt.Println("2. Inside the curly braces {}, you list the fields of the struct, each with a name and a type.")
	fmt.Println("3. You can then create variables of the struct type and access or modify its fields.")

	fmt.Println("\nExample of a struct:")
	fmt.Println("type Person struct {")
	fmt.Println("    Name string")
	fmt.Println("    Age  int")
	fmt.Println("}")

	fmt.Println("\nUsage of the struct:")
	fmt.Println("var p Person")
	fmt.Println("p.Name = \"John\"")
	fmt.Println("p.Age = 30")
	fmt.Println("fmt.Println(\"Name: \", p.Name)")
	fmt.Println("fmt.Println(\"Age: \", p.Age)")

	fmt.Println("\nIn the example provided, we define a struct 'Person' with two fields: 'Name' of type 'string' and 'Age' of type 'int'.")
	fmt.Println("Then, we create a variable 'p' of type 'Person' and assign values to its fields.")
	fmt.Println("Finally, we access and print the values of the fields using dot notation.")
}

func showChannels() {
	fmt.Println("Channels in Go are used for communication between goroutines.")
	fmt.Println("Example:")
	fmt.Println("ch := make(chan int)")

	fmt.Println("\nIn Go, channels are a powerful feature for synchronizing and communicating between goroutines.")
	fmt.Println("1. Channels allow goroutines to send and receive values to and from each other.")
	fmt.Println("2. Channels are typed, meaning you specify the type of data that can be sent and received through the channel.")
	fmt.Println("3. Channels can be created using the 'make' function, specifying the type of data the channel will carry.")

	fmt.Println("\nExample of creating a channel:")
	fmt.Println("ch := make(chan int)")

	fmt.Println("\nUsage of the channel:")
	fmt.Println("ch <- 42 // send value 42 to the channel")
	fmt.Println("value := <-ch // receive value from the channel")

	fmt.Println("\nIn the example provided, we create a channel 'ch' that can carry values of type 'int' using the 'make' function.")
	fmt.Println("Then, we demonstrate sending and receiving values through the channel.")
}

func showMultiThreading() {
	fmt.Println("Go supports multi-threading through goroutines.")
	fmt.Println("Goroutines are lightweight threads managed by the Go runtime.")
	fmt.Println("They enable concurrent execution of functions or methods.")

	fmt.Println("\nExample:")
	fmt.Println("func main() {")
	fmt.Println("    // Start a goroutine to execute a function concurrently")
	fmt.Println("    go sayHello()")
	fmt.Println("    // This line will be executed immediately after starting the goroutine")
	fmt.Println("    fmt.Println(\"Hello from main goroutine\")")
	fmt.Println("    // Wait for user input before exiting the program")
	fmt.Println("    fmt.Scanln()")
	fmt.Println("}")

	fmt.Println("\n// Function to be executed as a goroutine")
	fmt.Println("func sayHello() {")
	fmt.Println("    fmt.Println(\"Hello from goroutine\")")
	fmt.Println("}")

	fmt.Println("\nIn this example, when the main function is executed, it starts a goroutine using the 'go' keyword to execute the 'sayHello' function concurrently.")
	fmt.Println("While the 'sayHello' function is running in its own goroutine, the main function continues to execute.")
	fmt.Println("This allows 'Hello from main goroutine' to be printed immediately after starting the goroutine, without waiting for the 'sayHello' function to finish.")
	fmt.Println("As a result, 'Hello from goroutine' may be printed after or between the lines printed by the main function, depending on the scheduling by the Go runtime.")
	fmt.Println("The program then waits for user input before exiting, allowing time for both the main goroutine and the 'sayHello' goroutine to complete.")
}
func showBackendDevelopment() {
	fmt.Println("Go is often used for backend development due to its simplicity, performance, and built-in support for concurrency.\n")
	fmt.Println("It is commonly used to build web servers, APIs, and microservices.\n")
	fmt.Println("Frameworks like Gin, Echo, and Beego are popular choices for building web applications in Go.\n")

	fmt.Println("In Go, backend development refers to building server-side applications and services.")
	fmt.Println("1. Go's simplicity and readability make it well-suited for backend development.")
	fmt.Println("2. Its compiled nature leads to high performance and efficient resource utilization.")
	fmt.Println("3. Go's built-in support for concurrency through goroutines and channels makes it ideal for handling concurrent requests.")
	fmt.Println("4. Go is commonly used to build web servers to serve HTTP requests, APIs for inter-service communication, and microservices for building scalable and modular systems.")
	fmt.Println("5. There are several popular web frameworks and libraries in Go, such as Gin, Echo, and Beego, which provide features and utilities to streamline backend development.\n")

	fmt.Println("These frameworks offer features like routing, middleware support, request/response handling, and more, making it easier to develop robust and scalable web applications in Go.")
}

