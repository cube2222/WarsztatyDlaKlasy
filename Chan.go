package main

import "fmt"

func main() {
	myChannel := make(chan int)

	myChannel <- 5 // Odczyt
	x := <- myChannel // Zapis
	fmt.Println(x)
}
