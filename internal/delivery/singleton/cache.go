package singleton

import (
	"GorillaWebSocket/internal/delivery"
	"sync"
)

var instance *Cache = nil
var once sync.Once

type Cache struct {
	sync.RWMutex
	msg delivery.Response
}

func GetInstance() *Cache {
	once.Do(func() {
		instance = &Cache{}
	})
	return instance
}

func (c *Cache) Set(message delivery.Response) {
	c.Lock()
	c.msg = message
	c.Unlock()
}

func (c *Cache) Get() delivery.Response {
	c.RLock()
	defer c.RUnlock()
	return c.msg
}
