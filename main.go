package main

import "fmt"

func HelloWorld() string {
	return "Hello, World!"
}

func HelloMary() string {
	return "Hello, Mary!"
}


func main() {
	fmt.Println(HelloWorld())
	fmt.Println(HelloMary())
}
