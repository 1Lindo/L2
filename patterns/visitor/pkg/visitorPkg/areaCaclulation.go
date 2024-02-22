package visitorPkg

import (
	"fmt"
	"math"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	sideSq := s.Side
	a.area = sideSq * 2
	fmt.Printf("Square area is [%+v] \n", a.area)
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	r := c.Radius
	pi := math.Pi
	a.area = int(math.Pow(float64(r), 2) * pi)
	fmt.Printf("Square area is [%+v] \n", a.area)
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	sideRec := float64(r.Side)
	a.area = int(math.Pow(sideRec, 2)*math.Sqrt(3)) / 2
	fmt.Printf("Square area is [%+v] \n", a.area)
}
