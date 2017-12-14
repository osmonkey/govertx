package main

import "log"

type foo interface {
	Log()
}

type Woo struct {
}

func (w *Woo) Log() {
	log.Println("Woo")
}

func Too(f foo) {
	f.Log()
}

func main() {
	w := Woo{}
	Too(&w)
}
