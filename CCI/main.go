package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type roll struct {
	id int
}

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {

	rollNumbers := make(chan roll)

	//goroutine thread
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(1000)
		for i := 0; i < 1000; i++ {
			//wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				rollNumbers <- roll{
					id: result,
				}
			}()
		}
		wg.Wait()
		close(rollNumbers)
	}()

	// main thread
	for n := range rollNumbers {

		fmt.Println(n)
	}

}
