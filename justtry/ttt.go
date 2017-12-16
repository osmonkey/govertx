package main

import (
	"flag"
	"log"
	"net"
)

const (
	appid = "2eD234dfVA221!PSX"
)

func main() {

	s := flag.Bool("s", false, "shutdown")
	flag.Parse()
	c := make(chan error)
	if *s {
		go shutdown(&c)
	} else {
		go startShutdownServer(&c)
	}

	err := <-c
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown(c *chan error) {
	conn, err := net.Dial("tcp", "localhost:4444")
	if err != nil {
		*c <- err
		return
	}
	defer conn.Close()
	msg := []byte(appid)
	_, err = conn.Write(msg)
	if err != nil {
		*c <- err
		return
	}
	*c <- nil
}

func startShutdownServer(c *chan error) {
	l, err := net.Listen("tcp", "0.0.0.0:4444")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer l.Close()
	log.Println("Shutdown server listen")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleRequest(conn, c)
	}
}

func handleRequest(conn net.Conn, c *chan error) {
	defer conn.Close()
	buf := make([]byte, len(appid))
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err.Error())
	} else if string(buf) == appid {
		*c <- nil
	}
}
