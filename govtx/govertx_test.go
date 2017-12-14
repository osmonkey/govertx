package govtx

import (
	"github.com/satori/go.uuid"
	"log"
	"testing"
)

func TestNewGoVertx(t *testing.T) {
	gvtx := NewGoVertx()
	if len(gvtx.serviceMap) != 0 {
		t.Fail()
	}
}

func TestGoVertx_Add(t *testing.T) {
	ns := NewService{}
	gvtx := NewGoVertx()
	id := gvtx.Add(5, &ns)
	if len(gvtx.serviceMap) != 1 {
		t.Fail()
	} else {
		log.Printf("map len: %v\n", len(gvtx.serviceMap))
	}
	if id == uuid.Nil {
		t.Fail()
	} else {
		log.Printf("UUID: %s\n", id.String())
	}
}

type NewService struct {
	ServiceType
}

func (ns *NewService) Start() {
	log.Println("Start NewService")
}

func (ns *NewService) Stop() {
	log.Println("Stop NewService")
}
