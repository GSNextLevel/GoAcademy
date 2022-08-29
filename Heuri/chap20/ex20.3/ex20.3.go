package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender{}
	SendBook("Mastering Go", sender)
	SendBook("Mastering Rust", sender)
}
