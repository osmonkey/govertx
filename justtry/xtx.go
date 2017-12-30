package main

import (
	"encoding/json"
	"fmt"
)

type base struct {
	Id string
}

type person struct {
	base
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
	res.R[0].(*person).Id = "123"

	fmt.Println(res.R[0].(*person))

	idstr := do(res.R[0])

	fmt.Println("123" == *idstr)

	prs := res.R[0].(*person)
	prs.Name = "Lorbeer"
	fmt.Println(res.R[0].(*person).Name)

	ps := person{Name: "Olaf"}
	res.Add(&ps)

	for i := range res.R {
		fmt.Println(res.R[i].(*person).Name)
	}

	asByte, err := json.Marshal(res.R)
	s := string(asByte)
	fmt.Println(s)
}

func do(i interface{}) *string {
	switch i.(type) {
	case *person:
		return &i.(*person).Id
	case *base:
		return &i.(*base).Id
	default:
		panic("type not found")
	}
	//return new(string)
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
