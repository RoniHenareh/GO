package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := float64(1) // startgissning
	z_old := 0.0
	// stoppvillkor för liten skillnad

	for math.Abs(z_old-z) > 1e-8 {

		z_old = z // intressant, inte :=
		z = z - (z*z-x)/(2*z)
		fmt.Println("z", z)

	}

	return z
}

func main() {

	fmt.Println("vi får", Sqrt(2))
}
