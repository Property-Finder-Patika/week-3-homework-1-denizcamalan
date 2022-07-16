package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y     float64
	distance float64
}

func (p *Point) calculateDistance() float64 {
	p.distance = math.Sqrt(p.x*p.x + p.y*p.y)
	return p.distance
}

func (p *Point) calculateDistanceTo(q Point) float64 {
	distance := math.Sqrt(math.Pow((p.x-q.x), 2) + math.Pow((p.y-q.y), 2))
	return distance
}

func (p *Point) move(direction rune, amount float64) {
	if direction == 'x' {
		p.x += amount
	} else if direction == 'y' {
		p.y += amount
	}
	p.calculateDistance()
}

func (p *Point) print() {
	fmt.Printf("%#v\n", p)
}

func main() {
	processPoint()
	//processCircle()
}

func processPoint() {
	p1 := new(Point)
	p1.print()

	p1.move('x', 5)
	p1.move('y', 11)
	p1.print()

	p1.move('x', 4)
	p1.move('y', 1)
	p1.print()

	p2 := new(Point)
	fmt.Printf("Distance is %f\n", p1.calculateDistanceTo(*p2))
	fmt.Printf("Distance is %f\n", p2.calculateDistanceTo(*p1))
}

func processCircle() {
	p1 := Point{x: 3, y: 4}
	p1.calculateDistance()
	c1 := Circle{p1, 10}
	c1.print()

	c1.move('x', -3)
	c1.move('y', -4)
	c1.print()
}

type Circle struct {
	Point
	r int
}

func (c *Circle) calculateDistance1() float64 {
	c.distance = math.Sqrt(c.x*c.x + c.y*c.y)
	return c.distance
}

func (c *Circle) calculateDistance2() float64 {
	c.distance = math.Sqrt(c.Point.x*c.Point.x + c.Point.y*c.Point.y)
	return c.distance
}

func (c *Circle) calculateDistance3() float64 {
	return c.calculateDistance()
}

func (c *Circle) calculateDistance4() float64 {
	return c.Point.calculateDistance()
}

func (c *Circle) move(direction rune, amount float64) {
	c.Point.move(direction, amount)
}

func (c *Circle) print() {
	fmt.Printf("%#v\n", c)
}

// output :

// &main.Point{x:0, y:0, distance:0}
// &main.Point{x:5, y:11, distance:12.083045973594572}
// &main.Point{x:9, y:12, distance:15}
// Distance is 15.000000
// Distance is 15.000000