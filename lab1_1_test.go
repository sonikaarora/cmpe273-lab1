package main
import "testing"

/*
This structure will hold input for which fiboncacci series is required,
Second Parameter specifies expected series stored in slice.
*/
type testStructure struct {
	input int
	expectedSeries []int
}

/*
Test data 
*/
var testData = []testStructure {
{10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
{0, []int{0}},
{1, []int{0,1}},
{2, []int{0,1,1}},
{-1, []int{-1}}, /* This test is testing  for invalid input entered*/
}

/*
This method will compare series generated from Fibonacci method and expected series and return true if  both are equal and false otherwise.
*/
func compareSeries(seriesGenerated []int, expectedSeries []int) bool{
	for count:=0; count< len(seriesGenerated); count++ {
		  if seriesGenerated[count] != expectedSeries[count] {
		   return false
		  }
	}
	return true
}

/*
This method is calling Fibonacci method and executing it by providing it data from test data in for loop.
*/
func TestFibonacci(t *testing.T) {
	
	for _, data := range testData {
	no := data.input
	series, _:= Fibonacci(no)
		if series != nil {
		  statusCode := compareSeries(series,data.expectedSeries)
			  if statusCode == false {
					t.Error(
					 "Fibonacci series for" , data.input, "expected", data.expectedSeries, "got", series,
					)
				}	
		}

	}
	
}
