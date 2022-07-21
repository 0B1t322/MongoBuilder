package bsonfieldgetter

import (
	"reflect"
	"sync"
)

var casher Casher
var once sync.Once

func GetCasher() Casher {
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
	var strType string
	{
		t := reflect.TypeOf(model)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			t = t.Elem()
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
		}

		strType = t.String()
	}
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