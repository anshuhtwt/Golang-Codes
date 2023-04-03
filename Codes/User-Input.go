package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%v", &name)
	fmt.Println("Hey there,", name)
}
