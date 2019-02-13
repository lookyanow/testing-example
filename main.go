package main

import "fmt"

func main(){
	fmt.Println("Testing example")
	avg := []float64{1,2,3,4}
	fmt.Println(Average(avg))
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs)) 
}

