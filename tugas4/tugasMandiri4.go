package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var waitG sync.WaitGroup

	for _, number := range numbers {
		waitG.Add(1)
		go func(num int) {
			defer waitG.Done()
			func() {
				fmt.Println(num)
			}()
		}(number)
	}

	waitG.Wait()
}
