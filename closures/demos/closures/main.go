package main

import (
	"fmt"
	"sync"
)

func main() {

	i := 10
	s := "Todd"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("%d %s\n", i, s)
	}()
	wg.Wait()
}
