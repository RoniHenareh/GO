package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res
func sum(array []int, res chan<- int) {

	tot := 0
	for _, i := range array {
		tot += i
	}
	res <- tot // ty kan inte retunera när concurrent
}

// concurrently sum the array a
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	n1 := <-ch
	n2 := <-ch
	summa := n1 + n2

	return summa
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
