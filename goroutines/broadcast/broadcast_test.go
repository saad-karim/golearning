package broadcast_test

import (
	"testing"
	"time"

	"github.com/golearning/goroutines/broadcast"
)

func TestRunRoutines(t *testing.T) {
	l := broadcast.Listener{
		Name: "L1",
	}

	l2 := broadcast.Listener{
		Name: "L2",
	}

	p := broadcast.Producer{}
	b := broadcast.Broadcaster{
		Producer:  p,
		Listeners: []broadcast.Listener{l, l2},
	}

	ch := make(chan string)
	go b.Broadcast(ch)

	time.Sleep(10 * time.Second)
}
