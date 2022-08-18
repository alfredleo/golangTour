package main

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	group := &sync.WaitGroup{}
	group.Add(2)
	go check(1, group)
	go check(2, group)
	group.Wait()
}

func check(n int, group *sync.WaitGroup) {
	i := 0
	for i < 300 {
		fmt.Print(n)
		// create just a heavy calculation
		_ = new(big.Int).Exp(big.NewInt(19000), big.NewInt(90000), nil)
		i++
	}
	defer group.Done()
}
