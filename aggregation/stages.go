package aggregation

import (
	"github.com/0B1t322/MongoBuilder/object"
	"github.com/0B1t322/MongoBuilder/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type AddFieldsArger interface {
	formatAddFieldsArg() bson.M
	AddField(field string, value interface{}) AddFieldsArger
	AddFieldArger(field string, value AddFieldsArger) AddFieldsArger
}

type addFieldsArger struct {
	builder object.ObjectBuilder
}

func (a addFieldsArger) formatAddFieldsArg() bson.M {
	return a.builder.Build()
}

func (a addFieldsArger) AddField(field string, value interface{}) AddFieldsArger {
	a.builder.AddField(field, value)
	return a
}

func (a addFieldsArger) AddFieldArger(field string, value AddFieldsArger) AddFieldsArger {
	a.builder.AddField(field, value.formatAddFieldsArg())
	return a
}

func AddFieldArg() AddFieldsArger {
	return addFieldsArger{builder: object.Object()}
}

func AddFields(args ...AddFieldsArger) bson.M {
	return utils.MergeBsonM(
		func() (slice []bson.M) {
			for _, arg := range args {
				slice = append(slice, arg.formatAddFieldsArg())
			}
			return slice
		}()...,
	)
}

type BucketArger interface {
	formatBucketArg() bson.M
	GroupBy(field interface{}) BucketArger
	AddBondary(boundary interface{}) BucketArger
	AddBondaries(boundaries ...interface{}) BucketArger
	Default(value interface{}) BucketArger

	// You can use objectBuilder to build the output value
	Output(value interface{}) BucketArger
}

type bucketArger struct {
	groupBy      interface{}
	boundaries   []interface{}
	defaultValue interface{}
	output       interface{}
}

func (b *bucketArger) formatBucketArg() bson.M {
	optionals := bson.M{}
	{
		if b.defaultValue != nil {
			optionals["default"] = b.defaultValue
		}

		if b.output != nil {
			optionals["output"] = b.output
		}
	}
	return bson.M{
		"$bucket": utils.MergeBsonM(
			bson.M{
				"groupBy":    b.groupBy,
				"boundaries": bson.A(b.boundaries),
			},
			optionals,
		),
	}
}

func (b *bucketArger) GroupBy(field interface{}) BucketArger {
	b.groupBy = field
	return b
}

func (b *bucketArger) AddBondary(boundary interface{}) BucketArger {
	b.boundaries = append(b.boundaries, boundary)
	return b
}

func (b *bucketArger) AddBondaries(boundaries ...interface{}) BucketArger {
	b.boundaries = append(b.boundaries, boundaries...)
	return b
}

func (b *bucketArger) Default(value interface{}) BucketArger {
	b.defaultValue = value
	return b
}

func (b *bucketArger) Output(value interface{}) BucketArger {
	b.output = value
	return b
}

func BucketArg() BucketArger {
	return &bucketArger{}
}

func Bucket(
	arg BucketArger,
) bson.M {
	return arg.formatBucketArg()
}
