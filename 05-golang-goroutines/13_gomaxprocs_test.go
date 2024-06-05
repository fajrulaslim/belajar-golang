package golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	// runtime.GOMAXPROCS(20) // ubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

	group.Wait()
}
