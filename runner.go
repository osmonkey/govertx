package main

import (
	"./govtx"
	"errors"
	"log"
	"os"
)

func result(dr govtx.AsyncResult) {
	log.Println("Deploy done")
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
	os.Exit(0)
}

type Service1 struct {
	govtx.ServiceType
}

func (ns *Service1) Start() error {
	log.Println("Start Service1")
	return nil
}

func (ns *Service1) Stop() error {
	log.Println("Stop Service1")
	return nil
}

type Service2 struct {
	govtx.ServiceType
}

func (ns *Service2) Start() error {
	log.Println("Start Service2")
	return nil
}

func (ns *Service2) Stop() error {
	log.Println("Stop Service2")
	return nil
}

type Service3 struct {
	govtx.ServiceType
}

func (ns *Service3) Start() error {
	log.Println("Start Service3")
	return errors.New("3")
}

func (ns *Service3) Stop() error {
	log.Println("Stop Service3")
	return nil
}

type Service4 struct {
	govtx.ServiceType
}

func (ns *Service4) Start() error {
	log.Println("Start Service4")
	return nil
}

func (ns *Service4) Stop() error {
	log.Println("Stop Service4")
	return nil
}

type Service5 struct {
	govtx.ServiceType
}

func (ns *Service5) Start() error {
	log.Println("Start Service5")
	return nil
}

func (ns *Service5) Stop() error {
	log.Println("Stop Service5")
	return nil
}
