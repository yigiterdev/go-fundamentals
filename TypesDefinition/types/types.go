package types

import "fmt"

type (
	integer        int
	sliceOfStrings []string
	distance       float64
	distanceKm     float64
)

func (miles distance) ToKm() distanceKm {
	return distanceKm(miles * 1.6)
}

func TestType() {
	d := distance(10)
	fmt.Println(d.ToKm())
}
