package govtx

import (
	"errors"
	"github.com/satori/go.uuid"
	"sync"
)

type eventBus struct {
	handlerMap map[string][]handler
	mapMutex   sync.Mutex
}

type handler struct {
	F  func(message *Message)
	Id uuid.UUID
}

func NewEventBus() *eventBus {
	eb := eventBus{}
	eb.handlerMap = map[string][]handler{}
	return &eb
}

func (eb *eventBus) Consumer(key string, f func(message *Message)) uuid.UUID {
	id := uuid.NewV4()
	eb.mapMutex.Lock()
	defer eb.mapMutex.Unlock()
	_, ok := eb.handlerMap[key]
	if ok {
		eb.handlerMap[key] = append(eb.handlerMap[key], handler{f, id})
	} else {
		handlers := []handler{}
		handlers = append(handlers, handler{f, id})
		eb.handlerMap[key] = handlers
	}
	return id
}

func (eb *eventBus) Send(adr string, msg []byte, f func(result AsyncResult)) {
	_, ok := eb.handlerMap[adr]
	if !ok {
		f(AsyncResult{nil, errors.New("Adress not found")})
	} else {
		c := make(chan AsyncResult)
		defer close(c)
		go func() {
			m := Message{}
			m.SendBody = msg
			eb.handlerMap[adr][0].F(&m)
			c <- m.Result()
		}()
		f(<-c)
	}
}
