package main

import "fmt"

func testingFunc(i int) string {
	switch i {
	case 1:
		return "one"
	case 2:
		return "two"
	default:
		return printString(i)
	}
}

func printString(i int) string {
	return fmt.Sprint(i)
}

func main() {
	fmt.Println("Hello from Goland!")
	fmt.Println(testingFunc(1))
	fmt.Println(1, 2, 3, "hello")
}
