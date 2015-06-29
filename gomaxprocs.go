package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Number of CPU: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(-1))
}
