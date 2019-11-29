package locking_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golearning/goroutines/locking"
)

func TestLocking(t *testing.T) {
	shared := locking.New()

	w1 := locking.Writer{
		Name:   "w1",
		Values: map[string]int{"v1": 2, "v2": 5},
		Shared: shared,
	}

	w2 := locking.Writer{
		Name:   "w2",
		Values: map[string]int{"v1": 3, "v2": 6},
		Shared: shared,
	}

	go w1.Write()
	time.Sleep(time.Second)
	go w2.Write()

	for i := 0; i < 8; i++ {
		fmt.Println("shared data: ", shared.Data)
		time.Sleep(time.Second)
	}

	if shared.Data["v1"] != 7 {
		t.Error("wrong value, expecting 7 ---", shared.Data["v1"])
	}

	if shared.Data["v2"] != 16 {
		t.Error("wrong value, expecting 16 --- ", shared.Data["v2"])
	}
}
