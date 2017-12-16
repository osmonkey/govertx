package govtx

import (
	"encoding/binary"
	"github.com/satori/go.uuid"
	"sort"
)

type Service interface {
	Start() error
	Stop() error
}

type ServiceType struct {
	GoVertx *goVertx
}

type ServiceInfo struct {
	Priority int
	Service  Service
	Uuid     uuid.UUID
	Deployed bool
}

type goVertx struct {
	serviceMap map[int][]ServiceInfo
}

func NewGoVertx() *goVertx {
	gv := new(goVertx)
	gv.serviceMap = make(map[int][]ServiceInfo)
	return gv
}

func (gv *goVertx) Add(priority int, service Service) uuid.UUID {
	uuid := uuid.NewV4()
	_, ok := gv.serviceMap[priority]
	if ok {
		si := ServiceInfo{priority, service, uuid, false}
		gv.serviceMap[priority] = append(gv.serviceMap[priority], si)
	} else {
		si := ServiceInfo{priority, service, uuid, false}
		list := make([]ServiceInfo, 1)
		list[0] = si
		gv.serviceMap[priority] = list
	}
	return uuid
}

func (gv *goVertx) Deploy(f func(d AsyncResult)) {
	skeys := gv.sortedKeyList()
	i := 0
	for key := range skeys {
		k, err := gv.deploy(skeys[key])
		if err != nil {
			bs := make([]byte, 4)
			binary.LittleEndian.PutUint32(bs, uint32(i))
			f(AsyncResult{bs, err})
			return
		}
		i = i + k
	}
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(i))
	f(AsyncResult{bs, nil})
}

func (gv *goVertx) deploy(key int) (int, error) {
	services := gv.serviceMap[key]
	idx := []int{}
	for i := range services {
		if services[i].Deployed == false {
			idx = append(idx, i)
		}
	}
	chanels := make([]chan error, len(idx))
	for i := range chanels {
		chanels[i] = make(chan error)
	}
	for _, i := range idx {
		go func() {
			err := services[i].Service.Start()
			if err != nil {
				chanels[i] <- err
			} else {
				services[i].Deployed = true
				chanels[i] <- nil
			}
		}()
	}
	for i := range chanels {
		err := <-chanels[i]
		if err != nil {
			return 0, err
		}
	}

	return len(services), nil
}

func (gv *goVertx) Close() {
	chanels := make([]chan error, len(gv.serviceMap))
	for i := range chanels {
		chanels[i] = make(chan error)
	}
	for i, v := range gv.serviceMap {
		go func() {
			gv.close(&v)
			chanels[i] <- nil
		}()
	}
	for i := range chanels {
		<-chanels[i]
	}
}

func (gv *goVertx) close(si *[]ServiceInfo) {
	chanels := make([]chan error, len(*si))
	for i := range chanels {
		chanels[i] = make(chan error)
	}
	for i, v := range *si {
		go func() {
			chanels[i] <- v.Service.Stop()
		}()
	}
	for i := range chanels {
		<-chanels[i]
	}
}

func (gv *goVertx) sortedKeyList() []int {
	keyList := make([]int, 0)
	for k, _ := range gv.serviceMap {
		keyList = append(keyList, int(k))
	}
	sort.Ints(keyList)
	return keyList
}
