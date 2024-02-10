package main

import (
	"bufio"
	"fmt"
	"os"
)

const todoFile = "todos.txt"

func main() {
	for {
		showMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTodo()
		case 2:
			listTodos()
		case 3:
			deleteTodo()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\nTodo App Menu:")
	fmt.Println("1. Add Todo")
	fmt.Println("2. List Todos")
	fmt.Println("3. Delete Todo")
	fmt.Println("4. Exit")
}

func addTodo() {
	fmt.Print("Enter todo item: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	todo := scanner.Text()

	file, err := os.OpenFile(todoFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "%s\n", todo); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("Todo added successfully.")
}

func listTodos() {
	file, err := os.Open(todoFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("\nTodo List:")
	for i := 1; scanner.Scan(); i++ {
		fmt.Printf("%d. %s\n", i, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func deleteTodo() {
	listTodos()
	fmt.Print("Enter the number of todo to delete: ")
	var num int
	fmt.Scanln(&num)

	file, err := os.Open(todoFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if num <= 0 || num > len(lines) {
		fmt.Println("Invalid todo number.")
		return
	}

	lines = append(lines[:num-1], lines[num:]...)

	file, err = os.Create(todoFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := fmt.Fprintln(file, line); err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
	fmt.Println("Todo deleted successfully.")
}
