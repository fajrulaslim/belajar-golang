package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	// str1 := "Fajrul"
	// str2 := "Aslim"
	// pool.Put(&str1)
	// pool.Put(&str2)
	pool.Put("Fajrul")
	pool.Put("Aslim")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
