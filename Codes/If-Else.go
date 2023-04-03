package main

import "fmt"

func main() {
	salary := 600
	if salary > 10000 {
		salary += 1000
	} else if salary > 5000 {
		salary += 500
	} else {
		salary = salary
	}
	fmt.Println(salary)
}
