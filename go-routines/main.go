package main

import (
	"fmt"
	"sync"
)

func squareArray(nums []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range nums {
		square := num * num
		ch <- square
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, len(nums))
	var wg sync.WaitGroup

	// Launch a goroutine for each batch of array elements
	batchSize := 3
	for i := 0; i < len(nums); i += batchSize {
		end := i + batchSize
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go squareArray(nums[i:end], ch, &wg)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Collect results from the channel
	for square := range ch {
		fmt.Println(square)
	}
}
