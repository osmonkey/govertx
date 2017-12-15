package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("GO")
	sl := []chan bool{}
	for i := 0; i < 5; i++ {
		sl = append(sl, make(chan bool))
	}
	go l("Rob", 8*time.Second, sl[0])
	go l("Olaf", 7*time.Second, sl[1])
	go l("Claus", 2*time.Second, sl[2])
	go l("Marc", 4*time.Second, sl[3])
	go l("Hel", 3*time.Second, sl[4])

	for _, v := range sl {
		<-v
	}
	log.Println("Done")
}

func l(s string, d time.Duration, c chan bool) {
	time.Sleep(d)
	log.Printf("%s with %s", s, d.String())
	c <- true
}
