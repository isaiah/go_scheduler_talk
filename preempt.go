package main

import "sync/atomic"

func main() {
	// Test that goroutines can be preempted at function calls.
	const N = 5
	c := make(chan bool)
	var x uint32
	for g := 0; g < 2; g++ {
		go func(g int) {
			for i := 0; i < N; i++ {
				for atomic.LoadUint32(&x) != uint32(g) {
					preempt()
				}
				atomic.StoreUint32(&x, uint32(1-g))
			}
			c <- true
		}(g)
	}
	<-c
	<-c
}

// The function is used to test preemption at split stack checks.
var preempt = func() int {
	var a [128]int
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
