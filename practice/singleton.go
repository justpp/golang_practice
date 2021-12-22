package practice

import "sync"

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func getSingleton() *Singleton {

	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
