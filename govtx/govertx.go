package govtx

import (
	"github.com/satori/go.uuid"
	"sort"
)

type Service interface {
	Start()
	Stop()
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

func (gv *goVertx) Deploy(f func(d DeployResult)) {
	skeys := gv.sortedKeyList()
	for key := range skeys {
		gv.deploy(skeys[key])
	}

}

func (gv *goVertx) deploy(key int) int {
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
			services[i].Service.Start()
			services[i].Deployed = true
			chanels[i] <- nil
		}()
	}
	for i := range chanels {
		<-chanels[i]
	}

	return len(services)
}

//todo make async
func (gv *goVertx) Close() {
	for _, v := range gv.serviceMap {
		for i := range v {
			v[i].Service.Stop()
		}
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
