package main

import (
	"fmt"
)

func main() {
	content, _ := fetchPage()
	fmt.Println("Running from main.go")
	fmt.Printf(string(content))
}
