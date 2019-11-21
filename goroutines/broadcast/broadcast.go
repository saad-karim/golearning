package broadcast

import (
	"fmt"
	"time"
)

type Broadcaster struct{}

func (b *Broadcaster) SendMessageOnChannel(ch chan string) {
	counter := 0
	for {
		counter++
		ch <- fmt.Sprintf("message - %d", counter)
		time.Sleep(2 * time.Second)
	}
}

type Listner struct {
	Name string
}

func (l *Listner) ReadMessageOnChannel(ch chan string) {
	for {
		message := <-ch
		fmt.Printf("%s received message: %s\n", l.Name, message)
	}
}
