package main

import (
	"log"
)

type boo struct {
	Pstr *string
}

type zoo struct {
	Ptr interface{}
}

func main() {

	str := "Rob"
	pstr := &str
	log.Println(str)
	*pstr = "Roy"
	log.Println(str)

	b := boo{}
	b.Pstr = &str
	log.Println(*b.Pstr)

	*b.Pstr = "blah"

	log.Println(str)

	call(&str)
	show(&str)

	s := "ooo"
	sptr := &s
	i := 99

	z1 := zoo{}
	z1.Ptr = sptr

	z2 := zoo{}
	z2.Ptr = &i

	log.Println(i)
	log.Println(*z2.Ptr)

}

func call(p *string) {
	*p = "xxx"
}

func show(p *string) {
	log.Println(*p)
}
