package main

import "fmt"

type Car struct {
	make, model string
	year        int
	speed       int
	distance    int
	convertible bool
}

func (c *Car) accelerate(speed int) {
	c.speed = speed
}

func (c *Car) stop() {
	c.speed = 0
}

func (c *Car) move(distance int) {
	c.distance += distance
}

func (c *Car) print() {
	fmt.Printf("%#v\n", c)
}

func main() {
	myCar := Car{"Mercedes", "G", 2022, 0, 0, false}
	myCar.print()

	myCar.accelerate(100)
	myCar.move(200)
	myCar.accelerate(150)
	myCar.move(100)
	myCar.print()

	myCar.stop()
	myCar.print()
}

// output :

// &main.Car{make:"Mercedes", model:"G", year:2022, speed:0, distance:0, convertible:false}
// &main.Car{make:"Mercedes", model:"G", year:2022, speed:150, distance:300, convertible:false}
// &main.Car{make:"Mercedes", model:"G", year:2022, speed:0, distance:300, convertible:false}