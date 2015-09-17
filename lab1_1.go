package main
import "fmt"

/*
This method  push the series in slice and return it to main method, 
 along with boolean true if series is successfully generated and false otherwise.
*/
func Fibonacci(number  int)  ([]int,bool){
if number >=0 {
	seriesSlice := make([]int, number+1)
		for count:=0; count<= number; count++ {	
			returnedValue := FibonacciRecursion(count)	
			seriesSlice[count] = returnedValue
			}
	return seriesSlice, true
	}
return []int{-1}, false
}

func FibonacciRecursion(number int) int{
	 if number == 0 {
	  return 0
	  }
	 if number == 1 {
	  return 1
	  }
	 return (FibonacciRecursion(number -1) + FibonacciRecursion(number -2))
}

func main() {
	fmt.Println("Enter the no you want to get fibonacci series:")
	var input int
	fmt.Scanf("%d", &input)
	resultSeries, bool := Fibonacci(input)
	if bool == true {
	fmt.Print("fibonacci series for f(",input,")")
	fmt.Println(resultSeries)
	}
	if bool == false {
	fmt.Println("Please enter valid input")
	}
}

