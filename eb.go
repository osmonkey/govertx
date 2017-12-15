package main

import (
	"./govtx"
	"log"
	"time"
)

var cs chan bool

func main() {

	cs = make(chan bool)
	go func() {
		eb := govtx.NewEventBus()

		eb.Consumer("rob", con)
		eb.Send("rob", []byte("Olaf"), res)
	}()
	<-cs
}

func con(m *govtx.Message) {

	log.Println("Start Res")
	time.Sleep(3 * time.Second)
	log.Println(string(m.SendBody))
	m.Reply([]byte(""))

}

func res(ar govtx.AsyncResult) {

	if ar.Succeeded() {
		log.Printf("Succeeded: %v\n", string(ar.Result))
	} else {
		log.Fatal(ar.Cause)
	}
	cs <- true
}
