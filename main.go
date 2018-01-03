package main

import "fmt"

func testFunc(input int) string {
	switch input {
	case 1:
		return "one"
	default:
		return fmt.Sprint(input)
	}
}

func main() {
	fmt.Println("Hello from Goland!")
	fmt.Println(testFunc(1))
}
