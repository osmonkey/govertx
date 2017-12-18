package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("GO %d\n", runtime.NumCPU())
	sl := make(chan bool, 5)
	go l2("Rob", 8*time.Second, &sl)
	go l2("Olaf", 7*time.Second, &sl)
	go l2("Claus", 2*time.Second, &sl)
	go l2("Marc", 4*time.Second, &sl)
	go l2("Hel", 3*time.Second, &sl)

	for i := 0; i < 5; i++ {
		<-sl
	}
	log.Println("Done")
}

func l2(s string, d time.Duration, c *chan bool) {
	time.Sleep(d)
	log.Printf("%s with %s", s, d.String())
	*c <- true
}
