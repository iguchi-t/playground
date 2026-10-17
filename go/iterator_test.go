package main

import (
	"sync"
	"testing"
)

func Test_badIterator(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Log(i)
		}()
	}
	wg.Wait()
}

func Test_goodIterator(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			t.Log(n)
		}(i)
	}
	wg.Wait()
}
