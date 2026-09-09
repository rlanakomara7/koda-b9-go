package channel

import "fmt"

type Message struct {
	Sender  string
	Content string
}

func SendMessage(ch chan Message, sender string, content string) {
	message := Message{Sender: sender, Content: content}
	ch <- message
}

func WhiteBoard(ch chan Message) {
	for message := range ch {
		fmt.Println(message.Sender, message.Content)
	}
}
