package main

import "fmt"

var name string

func init() {
	fmt.Println("This will get called on main initialization")
	name = "Elliot"
}

func main() {
	fmt.Println("My Wonderful Go Program")
	fmt.Printf("Name: %s\n", name)
}
