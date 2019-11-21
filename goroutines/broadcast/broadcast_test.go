package broadcast_test

import (
	"goroutines/broadcast"
	"testing"
	"time"
)

func TestRunRoutines(t *testing.T) {
	b := broadcast.Broadcaster{}

	ch := make(chan string)
	go b.SendMessageOnChannel(ch)

	l := broadcast.Listner{
		Name: "L1",
	}

	go l.ReadMessageOnChannel(ch)
	time.Sleep(10 * time.Second)
}
