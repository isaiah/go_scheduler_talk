package main

import (
	"index/suffixarray"
	"math/rand"
)

func main() {
	const N = 6e5
	const M = 100
	const I = 20
	const J = 10
	const P = 2
	data := make([]byte, N)
	for i := 0; i < N; i++ {
		data[i] = byte(rand.Intn(255))
	}
	done := make(chan bool, P)
	for p := 0; p < P; p++ {
		go func() {
			for i := 0; i < I; i++ {
				suffix := suffixarray.New(data)
				for j := 0; j < J; j++ {
					str := make([]byte, M)
					for m := 0; m < M; m++ {
						str[m] = byte(rand.Intn(255))
					}
					_ = suffix.Lookup(str, 10)
				}
			}
			done <- true
		}()
	}
	for p := 0; p < P; p++ {
		<-done
	}
}
