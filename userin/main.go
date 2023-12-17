package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	//  Comma ok syntax || Comma error
	input ,_  := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}
