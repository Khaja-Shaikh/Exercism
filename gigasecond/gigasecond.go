

package gigasecond


import "time"

// AddGigasecond Returns the Time After Adding t duration of time .
func AddGigasecond(t time.Time) time.Time {

	t= t.Add(time.Second * 1000000000) 
	return t
}
