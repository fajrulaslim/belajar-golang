package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronus(group *sync.WaitGroup, n int) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Process", n)
	time.Sleep(1 * time.Second)
}
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronus(group, i)
	}

	group.Wait()
	fmt.Println("Complete")
}
