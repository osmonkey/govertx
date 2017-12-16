package main

import (
	n "./netutil"
	"log"
	"net"
)

func main() {
	u := n.NetUtil{}
	p, e := u.GetNextFree()
	log.Println(p)
	log.Println(e)

	l, _ := net.Listen("tcp", "0.0.0.0:55000")
	defer l.Close()

	log.Println(u.GetUsedPorts())
	n.P()
}
