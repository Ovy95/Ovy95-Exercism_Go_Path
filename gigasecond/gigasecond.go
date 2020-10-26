package gigasecond

import "time"

//Time package

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	/*In the package has a add function checked example can add different measurements
	of time to current time chose this to add the Gigasecond */
	return t
}
