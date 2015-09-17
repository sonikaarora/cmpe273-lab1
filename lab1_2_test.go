package main
import "testing"
import "fmt"

/*
Test data 
*/
func prepareData() []Shape{
	circle := Circle{radius:2}
	rectangle := Rectangle{length:6, width:5}
	shape := []Shape{&circle,&rectangle}
	return shape
}

var expectedParameters = []int {12,22}

/*
This method is calling Perimeter method for both shapes Circle and Rectangel and executing it by providing it data in for loop.
*/
func TestPerimeter(t *testing.T) {
        testData := prepareData()
	for count,shape := range testData {
		perimeter := int(shape.Perimeter())
		fmt.Println("permiter: ",perimeter)
		  if perimeter !=expectedParameters[count]{
					t.Error(
					 "Permieter for" , shape, "expected", expectedParameters[count], "got", perimeter,
					)
				}	
	}
}