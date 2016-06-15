package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := os.Create("/tmp/plik.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(file, "Hello world. I can see you.")

	file.Close()

	newFile, err := os.Open("/tmp/plik.txt") //Otwieramy ten sam plik który właśnie stworzyliśmy.
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(newFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	newFile.Close()
}
