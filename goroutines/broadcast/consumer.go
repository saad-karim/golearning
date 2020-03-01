package broadcast

import (
	"fmt"
)

type Listener struct {
	Name string
}

func (l *Listener) RecieveMessage(message string) {
	fmt.Printf("%s received message: %s\n", l.Name, message)
}
