package bsonfieldgetter

import (
	"reflect"
	"sync"
)

var casher *BsonFieldsCacher
var once sync.Once

func GetCasher() Casher {
	once.Do(func() {
		casher = NewBsonFieldsCacher()
	})

	return casher
}

func getCasher() *BsonFieldsCacher {
	once.Do(func() {
		casher = NewBsonFieldsCacher()
	})

	return casher
}

type Casher interface {
	Get(model interface{}) *BsonFieldGetter
}

type BsonFieldsCacher struct {
	cache map[string]*BsonFieldGetter
	sync.RWMutex
}

func NewBsonFieldsCacher() *BsonFieldsCacher {
	return &BsonFieldsCacher{
		cache:   make(map[string]*BsonFieldGetter),
		RWMutex: sync.RWMutex{},
	}
}

func (c *BsonFieldsCacher) Get(model interface{}) *BsonFieldGetter {
	strType := getTypeName(model)
	c.RLock()
	if bg, ok := c.cache[strType]; ok {
		c.RUnlock()
		return bg
	}
	c.RUnlock()
	c.Lock()
	defer c.Unlock()
	if bg, ok := c.cache[strType]; ok {
		return bg
	}
	bg := NewBsonFieldGetter(model)
	c.cache[strType] = bg
	return bg
}

func (c *BsonFieldsCacher) set(model interface{}, bg *BsonFieldGetter) {
	strType := getTypeName(model)

	c.cache[strType] = bg
}

func (c *BsonFieldsCacher) get(model interface{}) *BsonFieldGetter {
	strType := getTypeName(model)
	if bg, ok := c.cache[strType]; ok {
		return bg
	}
	bg := NewBsonFieldGetter(model)
	c.cache[strType] = bg
	return bg
}

func (c *BsonFieldsCacher) createIfNotExist(model interface{}) *BsonFieldGetter {
	strType := getTypeName(model)
	bg := c.getByTypeName(strType)
	if bg != nil {
		return bg
	}
	bg = NewBsonFieldGetter(model)
	c.cache[strType] = bg
	return bg
}

func (c *BsonFieldsCacher) getByTypeName(typeName string) *BsonFieldGetter {
	if bg, ok := c.cache[typeName]; ok {
		return bg
	}
	return nil
}

func getTypeName(of interface{}) string {
	t := reflect.TypeOf(of)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		t = t.Elem()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}

	return t.String()
}