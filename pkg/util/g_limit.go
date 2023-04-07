package util

import "sync"

type GLimit struct {
	n int
	c chan struct{}
}

func NewGLimit(limit int) *GLimit {
	return &GLimit{
		limit,
		make(chan struct{}, limit),
	}
}

func (l GLimit) Run(n int, f func(int)) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		l.run(i, f, &wg)
	}
	wg.Wait()
}

func (l *GLimit) run(i int, f func(int), wg *sync.WaitGroup) {
	l.c <- struct{}{}
	go func() {
		f(i)
		<-l.c
		wg.Done()
	}()
}
