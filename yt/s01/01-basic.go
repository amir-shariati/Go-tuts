package main

import "fmt"

func init() {
	fmt.Println("init function call before main function")
}

func main() {
	println("this is main function")
}
