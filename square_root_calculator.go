package main

import (
	"fmt"
	"math"
)

func squareRootCalculator() {
	fmt.Printf("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	var output float64
	output = math.Sqrt(input)

	if output == float64(int64(output)) {
		fmt.Println("Your answer is a whole number")
	} else {
		fmt.Println("Your answer did not return a whole number")
	}

}
