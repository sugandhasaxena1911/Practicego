package main

//https://go.dev/tour/methods/4
import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

func main() {

	c1 := circle{
		radius: 5,
	}
	fmt.Println(c1)
	//calling infoshape(shape) when non pointer c circle implements interface , thefore circle iof type s shape
	infoShape(c1)
	infoShape(&c1)
	//modify area so that recevr is pointer :
	//func (c *circle) area() float64
	//then  infoShape(c1) : gives error
	//Cannot use 'c1' (type circle) as the type shape Type, does not implement 'shape' ,as the 'area' method has a pointer receiver

}
func (c circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

type shape interface {
	area() float64
}

func infoShape(s shape) {
	fmt.Println("The area of the shape is ", s.area())

}
