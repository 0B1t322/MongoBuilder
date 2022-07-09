package bsonfieldgetter

import (
	"sync"
)

var casher Casher

func init() {
	casher = NewBsonFieldsCacher()
}

func GetCasher() Casher {
	return casher
}

type Casher interface {
	Get(model interface{}) *BsonFieldGetter
}

type BsonFieldsCacher struct {
	cache map[interface{}]*BsonFieldGetter
	sync.RWMutex
}

func NewBsonFieldsCacher() *BsonFieldsCacher {
	return &BsonFieldsCacher{
		cache: make(map[interface{}]*BsonFieldGetter),
		RWMutex: sync.RWMutex{},
	}
}

func (c *BsonFieldsCacher) Get(model interface{}) *BsonFieldGetter {
	c.RLock()
	if bg, ok := c.cache[model]; ok {
		c.RUnlock()
		return bg
	}
	c.RUnlock()
	c.Lock()
	defer c.Unlock()
	if bg, ok := c.cache[model]; ok {
		return bg
	}
	bg := NewBsonFieldGetter(model)
	c.cache[model] = bg
	return bg
}
