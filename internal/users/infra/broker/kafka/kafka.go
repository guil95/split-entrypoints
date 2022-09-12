package kafka

import (
	"fmt"
	"os"
	"time"
)

func Listen(quit chan os.Signal) {
	message := make(chan string)
	go listenMessage(message)

	for {
		select {
		case m := <-message:
			fmt.Println(m)
		}
	}
}

func listenMessage(message chan<- string) {
	for {
		time.Sleep(time.Second * 2)

		message <- "New message"
	}
}
