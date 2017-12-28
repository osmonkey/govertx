package main

import (
	"log"
	"time"
)

type voo struct {
	V []*A
}

type A struct {
	Id string
}

type B struct {
	C int
}

func one(c *chan int) {
	time.Sleep(3 * time.Second)
	log.Println("one")
	*c <- 1
}

func two(c *chan int) {
	time.Sleep(3 * time.Second)
	log.Println("two")
	*c <- 2
}

func main() {

	v := voo{}
	a := A{Id: "111"}
	ap := &a
	v.V = append(v.V, ap)
	v.V = append(v.V, &A{Id: "ddddd"})
	log.Println(*v.V[0])
	log.Println(*v.V[1])

	c := make(chan int, 3)

	go one(&c)
	go one(&c)
	go one(&c)
	go two(&c)
	go one(&c)
	go one(&c)

	for {
		log.Println(<-c)
	}

}
