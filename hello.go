package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello," + name)
}

func main() {
	sayHello("Bobby")
	number :=TestFunction(5)
	fmt.Println(number)
}
