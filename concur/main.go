package main

import (
	"fmt"
)

// Go Routine ->

// func main() {
// 	// Fork join model
// 	go func() {
// 		fmt.Println("Ayuhs")
// 	}()
// 	time.Sleep(time.Second *2 );

// 	fmt.Println(")=>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
// }

// Channels are used to communicate between the go-routines
// -=====================================>> routine 1
// -=====================================>> routine 2
// -=====================================>> routine 3
// -=====================================>> routine 4

// So channels are the FIFO queue in which one channels enters and one channels read

// func main() {
// 	myChannel := make(chan string)
// 	myChannel2 := make(chan string)
// 	go func() {
// 		myChannel2 <- "Cow "

// 	}()
// 	go func() {
// 		myChannel <- "Hello from the Ano fn "
// 	}()
// 	select {
// 	case f := <-myChannel:
// 		fmt.Println(f)
// 	case g := <-myChannel2:
// 		fmt.Println("aada-> ", g)
// 	}
// }

// Pattern 1 -> for select loop
//  We have to insert a number of char in the channel

func main() {
	cnl := make(chan string, 3)
	cnl2 := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case cnl <- s:
			fmt.Println("Hello World in the for loop -> <-")

		case cnl2 <- s:
			panic("fdf")
		}

	}
	close(cnl)
	for res := range cnl {
		fmt.Println(res)
	}

}
