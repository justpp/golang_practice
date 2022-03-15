package practice

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := getSingleton()
	s2 := getSingleton()
	if s1 != s2 {
		t.Fatal("Instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	const parCount = 100
	wg.Add(parCount)
	instances := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		go func(i int) {
			instances[i] = getSingleton()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("Instance is not equal")
		}
	}
}
