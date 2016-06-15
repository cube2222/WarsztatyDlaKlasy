package main

import "fmt"

type wheeler interface {
	printWheelCount()
}

type car struct {
	color int
}

func (c car) printWheelCount() {
	fmt.Println("I have 4 wheels and my color is ", c.color)
}

type bike struct {
}

func (b bike) printWheelCount() {
	fmt.Println("I have 2 wheels and no color.")
}

func main() {
	var myBike wheeler
	var myCar wheeler

	myBike = bike{}
	myCar = car{color: 1}

	myBike.printWheelCount()
	myCar.printWheelCount()

}
