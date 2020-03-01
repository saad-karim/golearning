package broadcast

import (
	"fmt"
	"sync"
)

type Broadcaster struct {
	Listeners []Listener
}

func New() *Broadcaster {
	return &Broadcaster{
		Listeners: []Listener{},
	}
}

func (b *Broadcaster) RegisterListener(listener Listener) {
	b.Listeners = append(b.Listeners, listener)
}

func (b *Broadcaster) Broadcast(ch chan string, wg *sync.WaitGroup) {
	for {
		fmt.Println("Waiting to broadcast message...")
		message := <-ch
		fmt.Println("Broadcast message: ", message)
		for _, l := range b.Listeners {
			l.RecieveMessage(message)
		}
		wg.Done()
	}
}
