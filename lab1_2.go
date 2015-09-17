package main
import ("fmt";"math")

type Shape interface {
	Perimeter() float64 
}

type Circle struct {
	radius float64
}

func (circle *Circle) Perimeter() float64 {
	fmt.Println("Perimeter of circle wih radius ",circle.radius)
	perimeter := 2* math.Pi * circle.radius
	return perimeter
}
	
type Rectangle struct {
	length float64
	width float64
}

func (rectangle *Rectangle)  Perimeter() float64{
	fmt.Println("Perimeter of rectangle with length ",rectangle.length," and width ",rectangle.width)
	return 2*(rectangle.length+rectangle.width)
}

func main() {
	circle := Circle{radius:10}
	rectangle := Rectangle{length:12, width:5}
	shape := []Shape{&circle,&rectangle}

	for _,selectedShape:= range shape {
		fmt.Println(int(selectedShape.Perimeter()))
	}
}