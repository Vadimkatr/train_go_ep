package person

import (
	"time"
	"strconv"
)

type Person struct {
	FirstName string
	LastName  string
	BirthDay  time.Time
}

func (p Person) String() string {
	return p.FirstName + " " + p.LastName + " " + p.BirthDay.Format("2006-01-02")
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	personIBDay, personJBDay := p[i].BirthDay.Unix(), p[j].BirthDay.Unix()
	personIFName, personJFName := p[i].FirstName, p[j].FirstName
	personILName, personJLName := p[i].LastName, p[j].LastName

	if (personIBDay > personJBDay) || 
		(personIBDay == personJBDay && personIFName < personJFName) ||	
		(personIBDay == personJBDay && personIFName == personJFName && personILName < personJLName){
		return true
	}
	return false
}

func (p People) String() (res string) {
	for i, per := range p {
		res += strconv.Itoa(i + 1) + ") " + per.String() + ".\n"
	}
	return
}
