package main

import (
	"fmt"
	"runtime"
	"time"
)

const loop = 2

func main() {
	runtime.GOMAXPROCS(2)
	for i := 0; i < loop; i++ {
		go func() {
			for {
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Hello, 世界")
}
