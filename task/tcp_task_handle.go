package task

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type TCPTaskHandler interface {
	ServeTCP(c chan []byte, msg interface{})
}

type TCPTaskFunc func(c chan []byte, msg interface{})

func (t TCPTaskFunc) ServeTCP(c chan []byte, msg interface{}) {
	t(c, msg)
}

type tcpTaskHandle struct {
	handler TCPTaskHandler
	timeOut time.Duration
}

var (
	tcpHandlePoolMu sync.Mutex
	tcpHandlePool   = make(map[int32]*tcpTaskHandle)
)

func RegisterTCPTaskHandle(id int32, handler TCPTaskHandler, timeOut time.Duration) {
	tcpHandlePoolMu.Lock()
	defer tcpHandlePoolMu.Unlock()
	newHandle := &tcpTaskHandle{
		handler: handler,
		timeOut: timeOut,
	}
	tcpHandlePool[id] = newHandle
}

func GetTCPTaskHandle(id int32) (*tcpTaskHandle, error) {
	tcpHandlePoolMu.Lock()
	defer tcpHandlePoolMu.Unlock()
	if handle, ok := tcpHandlePool[id]; ok {
		return handle, nil
	} else {
		return nil, errors.New("can't find  handle")
	}
}

func DumpTCPTaskHandle() {
	tcpHandlePoolMu.Lock()
	defer tcpHandlePoolMu.Unlock()
	for k, v := range tcpHandlePool {
		fmt.Println(k, v)
	}
}
