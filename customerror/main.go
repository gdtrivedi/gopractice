package main

import (
	"fmt"
	"math"

	"github.com/gdtrivedi/gopractice/customerror/myerror"
	"github.com/gdtrivedi/gopractice/customerror/myerror1"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, myerror.New("Area calculation failed, radius is less than zero", "MYERR001")
		// return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func circleArea1(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &myerror1.MyError1{"Area calculation failed, radius is less than zero", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	testMyError()
	testMyError1()
	here()
}

func testMyError() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err, ", Radius: ", radius)
		return
	}
	fmt.Println("Area of circle %0.2f", area)
}

func testMyError1() {
	radius := -20.0
	area, err := circleArea1(radius)
	if err != nil {
		if err, ok := err.(*myerror1.MyError1); ok {
			fmt.Printf("Radius %0.2f is less than zero \n", err.GetRadius())
			fmt.Printf("Radius %0.2f is less than zero \n", err.Radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("Area of circle %0.2f", area)
}

func here() {
	fmt.Println("Here")
}
