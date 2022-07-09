package bsonfieldgetter

import (
	"reflect"
	"strings"
)

type BsonFieldGetter struct {
	structFieldToBsonField map[string]string
}

func NewBsonFieldGetter(
	model interface{},
) *BsonFieldGetter {
	b := &BsonFieldGetter{
		structFieldToBsonField: make(map[string]string),
	}
	b.init(model)
	return b
}

func (b *BsonFieldGetter) init(model interface{}) {
	t := reflect.TypeOf(model)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
		if t.Kind() == reflect.Ptr  {
			t = t.Elem()
		}
	}

	if t.Kind() != reflect.Struct {
		panic("model must be a struct")
	}

	b.initType(t)
}

func (b *BsonFieldGetter) initTypeWithParentField(parentField string, t reflect.Type) {
	numsField := t.NumField()
	for i := 0; i < numsField; i++ {
		field := t.Field(i)
		b.initField(parentField, field)
	}
}

// init first level of struct
func (b *BsonFieldGetter) initType(t reflect.Type) {
	b.initTypeWithParentField("", t)
}

func (b *BsonFieldGetter) initField(parentField string, field reflect.StructField) {
	t := field.Type

	if t.Kind() == reflect.Ptr  {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
		if t.Kind() == reflect.Ptr  {
			t = t.Elem()
		}
	}

	tag := field.Tag.Get("bson")

	if tag == "-" || tag == "" {
		return
	}

	if strings.Contains(tag, ",inline") {
		if t.Kind() != reflect.Struct {
			return
		}
		b.initTypeWithParentField(parentField, t)
		return
	}

	if parentField == "" {
		b.structFieldToBsonField[field.Name] = tag
	} else {
		b.structFieldToBsonField[parentField+"."+field.Name] = b.structFieldToBsonField[parentField] + "." + tag
	}

	if t.Kind() == reflect.Struct {
		if parentField == "" {
			b.initTypeWithParentField(field.Name, t)
		} else {
			b.initTypeWithParentField(parentField+"."+field.Name, t)
		}
	}
}

func (b BsonFieldGetter) GetMap() map[string]string {
	return b.structFieldToBsonField
}

func (b *BsonFieldGetter) Get(field string) string {
	return b.structFieldToBsonField[field]
}