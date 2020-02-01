package main

import "fmt"

func getShootingPercentage(goals int, attempts int) float64 {
	percentage := 0.0

	if goals != 0 {
		percentage = float64(goals) / float64(attempts)
	}
	return percentage * 100
}

func main() {
	fmt.Println("Hockey Stats!")
	myPercentage := getShootingPercentage(100, 10)
	fmt.Println(myPercentage)
}
