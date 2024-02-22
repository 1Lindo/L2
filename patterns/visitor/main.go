package main

import "main/pkg/visitorPkg"

func main() {
	square := &visitorPkg.Square{Side: 2}
	circle := &visitorPkg.Circle{Radius: 5}
	rectangle := &visitorPkg.Rectangle{Side: 5}

	areaCalculator := &visitorPkg.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)
	//getPaymentShluse

}
