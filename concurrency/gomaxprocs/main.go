package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Default GOMAXPROCS:\t", runtime.GOMAXPROCS(-1))
	fmt.Println("Number of CPUs:\t\t", runtime.NumCPU())
	runtime.GOMAXPROCS(20)
	fmt.Println("Adjusted GOMAXPROCS:", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(300)
	fmt.Println("Maximum GOMAXPROCS:\t", runtime.GOMAXPROCS(-1))
}
