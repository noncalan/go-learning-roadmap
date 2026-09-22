package main

import "fmt"

func main() {
	for c := 'a'; c <= 'z'; c++ {
		fmt.Printf("%c", c)
	}
	println()
	for c := 'z'; c >= 'a'; c-- {
		fmt.Printf("%c", c)
	}
}
