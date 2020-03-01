package broadcast

type Producer struct{}

func (p *Producer) SendMessage(ch chan string, message string) {
	ch <- message
}
