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

package main

import "fmt"

// HelloWorld returns a greeting message.
func HelloWorld() string {
	return "Hello, World!"
}

// HelloEnoc returns a greeting message specifically for Enoc.
func HelloEnoc() string {
	return "Hello, Enoc!"
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(HelloEnoc())
}
