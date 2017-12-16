package main

import "log"

type woo struct {
	Str string
	B   bool
}

func (w *woo) P() {
	log.Println(w.Str)
}
func main() {

	w := []woo{}
	w = append(w, woo{"A", false})
	w = append(w, woo{"C", false})
	w = append(w, woo{"T", true})
	w = append(w, woo{"X", false})

	s := []int{}
	for i := range w {
		if w[i].B == false {
			s = append(s, i)
		}
	}

	for _, v := range s {
		w[v].P()
	}
}
