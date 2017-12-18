package govtx

import (
	"encoding/binary"
	"errors"
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
	chanel := make(chan error, len(idx))
	defer close(chanel)
	for _, x := range idx {
		go func(k int) {
			err := services[k].Service.Start()
			if err != nil {
				chanel <- err
			} else {
				services[k].Deployed = true
				chanel <- nil
			}
		}(x)
	}
	for _ = range idx {
		err := <-chanel
		if err != nil {
			return 0, err
		}
	}
	return len(services), nil
}

func (gv *goVertx) Undeploy(uuid uuid.UUID) error {
	return errors.New("Unimplemented")
}

func (gv *goVertx) Close() error {
	chanel := make(chan error, len(gv.serviceMap))
	defer close(chanel)
	for _, v := range gv.serviceMap {
		go func(v []ServiceInfo) {
			chanel <- gv.close(&v)
		}(v)
	}
	for i := 0; i < len(gv.serviceMap); i++ {
		err := <-chanel
		if err != nil {
			return err
		}
	}
	return nil
}

func (gv *goVertx) close(si *[]ServiceInfo) error {
	chanel := make(chan error, len(*si))
	defer close(chanel)
	for _, v := range *si {
		go func(v ServiceInfo) {
			chanel <- v.Service.Stop()
		}(v)
	}
	for _ = range chanel {
		err := <-chanel
		if err != nil {
			return err
		}
	}
	return nil
}

func (gv *goVertx) sortedKeyList() []int {
	keyList := make([]int, 0)
	for k, _ := range gv.serviceMap {
		keyList = append(keyList, k)
	}
	sort.Ints(keyList)
	return keyList
}
