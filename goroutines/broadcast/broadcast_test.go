package broadcast_test

import (
	"sync"
	"testing"

	"github.com/saad-karim/golearning/goroutines/broadcast"
)

func TestRunRoutines(t *testing.T) {
	l1 := broadcast.Listener{
		Name: "L1",
	}

	l2 := broadcast.Listener{
		Name: "L2",
	}

	p := broadcast.Producer{}

	b := broadcast.New()
	b.RegisterListener(l1)
	b.RegisterListener(l2)

	wg := &sync.WaitGroup{}

	ch := make(chan string)
	go b.Broadcast(ch, wg)

	messages := []string{"hi", "how", "are", "you"}
	for _, message := range messages {
		wg.Add(1)
		p.SendMessage(ch, message)
	}
	wg.Wait()
}
