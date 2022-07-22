package bsonfieldgetter

import (
	"reflect"
	"strings"
)

type BsonFieldGetter struct {
	structFieldToBsonField map[string]string
	// in key field represent name in value nil value
	subTypes map[string]interface{}
}

func NewBsonFieldGetter(
	model interface{},
) *BsonFieldGetter {
	b := &BsonFieldGetter{
		structFieldToBsonField: make(map[string]string),
		subTypes:               make(map[string]interface{}),
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
		if t.Kind() == reflect.Ptr {
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

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}

	tag := field.Tag.Get("bson")

	if tag == "-" || tag == "" {
		return
	}

	if strings.Contains(tag, ",inline") && t.Kind() == reflect.Struct {
		b.initTypeWithParentField(parentField, t)
		return
	}

	tag = strings.Split(tag, ",")[0]

	if parentField == "" {
		b.structFieldToBsonField[field.Name] = tag
	} else {
		b.structFieldToBsonField[parentField+"."+field.Name] = b.structFieldToBsonField[parentField] + "." + tag
	}

	if t.Kind() == reflect.Struct {
		in := reflect.New(t).Interface()
		if parentField == "" {
			b.subTypes[field.Name] = in
		} else {
			b.subTypes[parentField+"."+field.Name] = in
		}
		return
	}
}

// TODO: Refactor recursive types
// func (b BsonFieldGetter) GetMap() map[string]string {
// 	newMap := make(map[string]string)
// 	{
// 		for k, v := range b.structFieldToBsonField {
// 			newMap[k] = v
// 		}

// 		for subTypeField, subType := range b.subTypes {
// 			for k, v := range getCasher().Get(subType).GetMap() {
// 				newMap[subTypeField+"."+k] = b.structFieldToBsonField[subTypeField] + "." + v
// 			}
// 		}
// 	}
// 	return newMap
// }

func (b *BsonFieldGetter) Get(field string) string {
	splited := strings.SplitN(field, ".", 2)
	if len(splited) > 1 {
		if subType, ok := b.subTypes[splited[0]]; ok {
			return b.structFieldToBsonField[splited[0]] + "." + getCasher().Get(subType).Get(splited[1])
		}
	}
	return b.structFieldToBsonField[field]
}
