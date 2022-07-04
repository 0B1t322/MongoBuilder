package object

import "go.mongodb.org/mongo-driver/bson"

type ObjectBuilder interface {
	// if value is ObjectBuilder will build the object
	AddField(field string, value interface{}) ObjectBuilder
	Build() bson.M
}

type objectBulder struct {
	object bson.M
}

func (o *objectBulder) AddField(field string, value interface{}) ObjectBuilder {
	if v, ok := value.(ObjectBuilder); ok {
		o.object[field] = v.Build()
	} else {
		o.object[field] = value
	}
	return o
}

func (o *objectBulder) Build() bson.M {
	return o.object
}

func Object() ObjectBuilder {
	return &objectBulder{object: bson.M{}}
}
