package main
import ("testing"
	     "time"
	    )

var testData = []time.Duration{0 * time.Second,
					       1 * time.Second,
					       10 * time.Second,
					       }

func TestSleep(t *testing.T) {
	for _, data := range testData {
	currentTime := time.Now().UTC() // Time when sleep function is about to execute
	awakeTime, status := Sleep(data) // Sleep  function will return time, at which it comes awake
	awakeTime = awakeTime.UTC()
	 if status == true && (awakeTime.Sub(currentTime)) < data  {
					t.Error(
					 "Sleep function was expected to sleep for " , data, "but slept for duration ",awakeTime.Sub(currentTime),
					)
				}
	
	}

}