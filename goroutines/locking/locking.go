package locking

import (
	"fmt"
	"sync"
	"time"
)

type Shared struct {
	Data  map[string]int
	Mutex *sync.Mutex
}

func New() *Shared {
	return &Shared{
		Data:  map[string]int{},
		Mutex: &sync.Mutex{},
	}
}

func (s *Shared) StoreData(key string, value int) {
	fmt.Println("StoreData() -- Entered")
	s.Mutex.Lock()
	fmt.Println("StoreData() LOCKED")

	fmt.Printf("Storing key '%s' and value '%d'\n", key, value)
	s.Data[key] = (2 * s.Data[key]) + value

	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)

	fmt.Println("StoreData() UNLOCKED")
	s.Mutex.Unlock()
}

type Writer struct {
	Name   string
	Values map[string]int
	Shared *Shared
}

func (w *Writer) Write() {
	for key, value := range w.Values {
		fmt.Printf("'%s' is writing...\n", w.Name)
		w.Shared.StoreData(key, value)
		fmt.Printf("'%s' stored data '%d' \n", w.Name, value)
	}
}
