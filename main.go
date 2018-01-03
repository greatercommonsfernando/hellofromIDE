package main

import "fmt"

func testFunc(input int) string {
	return fmt.Sprint(input)
}

func main() {
	fmt.Println("Hello from Goland!")
	fmt.Println(testFunc(1))
}
