package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
}

type Result struct {
	R []interface{}
}

func (r *Result) Add(i interface{}) {
	r.R = append(r.R, i)
}

func main() {

	res := Result{}
	str := []byte(`{"name":"Rob"}`)

	pptr := person{}

	err := json.Unmarshal(str, &pptr)
	if err != nil {
		fmt.Println(err.Error())
	}
	res.Add(&pptr)

	fmt.Println(res.R[0].(*person).Name)

	res.R[0].(*person).Name = "Lisa"

	fmt.Println(res.R[0].(*person).Name)

	prs := res.R[0].(*person)
	prs.Name = "Lorbeer"
	fmt.Println(res.R[0].(*person).Name)

	ps := person{Name: "Olaf"}
	res.Add(&ps)

	for i := range res.R {
		fmt.Println(res.R[i].(*person).Name)
	}
}

func Np() {
	str := []byte(`{"name":"Rob"}`)

	pptr := person{}

	err := json.Unmarshal(str, &pptr)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func Wp() {
	str := []byte(`{"name":"Rob"}`)

	pptr := new(person)

	err := json.Unmarshal(str, pptr)
	if err != nil {
		fmt.Println(err.Error())
	}

}
