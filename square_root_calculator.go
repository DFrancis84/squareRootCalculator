package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	var output float64
	output = math.Sqrt(input)

	fmt.Println("They Square Root of", input, "is", output)

	if output == float64(int64(output)) {
		fmt.Println("Your answer returned a Perfect Square!")
	} else {
		fmt.Println("Your answer did not return a Perfect Square.")
	}

}
