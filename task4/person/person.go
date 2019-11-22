package person

import (
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	BirthDay  time.Time
}
