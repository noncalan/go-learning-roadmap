package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)
	var age int
	fmt.Scan(&age)
	fmt.Printf("Hello, %s!\n", name)
	fmt.Print("You are ", age, " years old.")

}
