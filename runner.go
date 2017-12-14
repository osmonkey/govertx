package main

import (
	"./govtx"
	"log"
)

func result(dr govtx.DeployResult) {
	log.Println("Done")
}

func main() {

	govtx := govtx.NewGoVertx()

	s1 := Service1{}
	s2 := Service2{}
	s3 := Service3{}
	s4 := Service4{}
	s5 := Service5{}

	govtx.Add(5, &s1)
	govtx.Add(5, &s3)
	govtx.Add(5, &s5)
	govtx.Add(2, &s4)
	govtx.Add(2, &s2)

	govtx.Deploy(result)

	govtx.Close()

}

type Service1 struct {
	govtx.ServiceType
}

func (ns *Service1) Start() {
	log.Println("Start Service1")
}

func (ns *Service1) Stop() {
	log.Println("Stop Service1")
}

type Service2 struct {
	govtx.ServiceType
}

func (ns *Service2) Start() {
	log.Println("Start Service2")
}

func (ns *Service2) Stop() {
	log.Println("Stop Service2")
}

type Service3 struct {
	govtx.ServiceType
}

func (ns *Service3) Start() {
	log.Println("Start Service3")
}

func (ns *Service3) Stop() {
	log.Println("Stop Service3")
}

type Service4 struct {
	govtx.ServiceType
}

func (ns *Service4) Start() {
	log.Println("Start Service4")
}

func (ns *Service4) Stop() {
	log.Println("Stop Service4")
}

type Service5 struct {
	govtx.ServiceType
}

func (ns *Service5) Start() {
	log.Println("Start Service5")
}

func (ns *Service5) Stop() {
	log.Println("Stop Service5")
}
