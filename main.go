package main

import (
	"fmt"
	"runtime"
)

func main() {
	i := 0
	for i < 3 {
		i++
		fmt.Println("hello,world")
	}
}

func Version() string {
	return runtime.Version()
}
