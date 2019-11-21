package broadcast

import (
	"fmt"
	"time"
)

type Broadcaster struct {
	Producer  Producer
	Listeners []Listener
}

func (b *Broadcaster) Broadcast(ch chan string) {
	go b.Producer.Message(ch)
	for {
		message := <-ch
		for _, l := range b.Listeners {
			l.ReadMessageOnChannel(message)
		}
	}
}

type Producer struct{}

func (p *Producer) Message(ch chan string) {
	counter := 0
	for {
		counter++
		ch <- fmt.Sprintf("message - %d", counter)
		time.Sleep(2 * time.Second)
	}
}

type Listener struct {
	Name string
}

func (l *Listener) ReadMessageOnChannel(message string) {
	fmt.Printf("%s received message: %s\n", l.Name, message)
}
