package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond Given a moment, determine the moment that would be after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
