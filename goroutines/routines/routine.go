package routines

import (
	"fmt"
	"time"
)

func RunRoutines() {
	go Routine1()
	go Routine2()
	go Routine3()

	time.Sleep(10 * time.Second)
}

func Routine1() {
	for {
		fmt.Println("1 - running")
		time.Sleep(time.Second)
	}
}

func Routine2() {
	for {
		fmt.Println("2 - running")
		time.Sleep(2 * time.Second)
	}
}

func Routine3() {
	for {
		fmt.Println("3 - running")
		time.Sleep(3 * time.Second)
	}
}
