package main
import ("fmt"
            "time"
	    )
	    
func main() {
	var input int
	fmt.Println("Enter the duration for which you want function to sleep: ")
	fmt.Scanf("%d",&input)
	fmt.Println("About to invoke sleep api, that will cause sleep for ",input,"seconds:", time.Now())
	awakeTime, status := Sleep(time.Duration(input) * time.Second)
		if status == false {
			fmt.Println("Please enter valid input for sleep function") 
			return //  If status false is returned from sleep function, this function will return from here
		}
	fmt.Println("Getting back from sleep api:", awakeTime)
	
}

/*
Sleep function using 'time.After' that will sleep for time greater or equal to input provided by the user
Current time and status is returned by this sleep function, status false is returned if invalid input is send to sleep function like some negative number.

*/
func Sleep(sleepTime  time.Duration) (time.Time, bool){

	if sleepTime< 0 {
	return time.Now(), false
	}
	select {
    		case <-time.After(sleepTime):
    		}
	return time.Now(), true
}