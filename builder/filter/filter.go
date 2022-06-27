package filter

import "go.mongodb.org/mongo-driver/bson"

const (
	EQ         = "$eq"
	NE         = "$ne"
	OR         = "$or"
	AND        = "$and"
	NOT        = "$not"
	GT         = "$gt"
	GTE        = "$gte"
	LT         = "$lt"
	LTE        = "$lte"
	IN         = "$in"
	NIN        = "$nin"
	EXISTS     = "$exists"
	REGEX      = "$regex"
	ELEM_MATCH = "$elemMatch"
)

type Filter interface {
	BSON() bson.M

	// return {"$or": [{...},...,{...}]}
	Or(filters ...Filter) Filter

	// return {"$and": [{...},...,{...}]}
	And(filters ...Filter) Filter

	EQField(
		field string,
		value interface{},
	) Filter

	EQ(
		field string, 
		value interface{},
	) Filter

	NEQ(
		field string, 
		value interface{},
	) Filter

	LT(
		field string, 
		value interface{},
	) Filter

	LTE(
		field string, 
		value interface{},
	) Filter

	GT(
		field string, 
		value interface{},
	) Filter

	GTE(
		field string, 
		value interface{},
	) Filter

	In(
		field string, 
		value ...interface{},
	) Filter

	NotIn(
		field string, 
		value ...interface{},
	) Filter

	Exist(field string) Filter

	NotExist(field string) Filter
	
	/* 
	Used to not make field in dot notaion
		Example:  
			filter.NewFindFilter().
				Field(
					"field",
					filter.NewFindFilter().
						EQField("some", 1).
						Field(
							"first",
							filter.NewFindFilter().
								EQField("one", 1).
								EQField("two", 2),
						),
				).
				BSON()
		Reptesent in bson: 
			{
				"field": {
					"first": {
						"one": 1,
						"two": 2,
					},
					"some": 1,
				},
			}
	*/
	Field(
		field string, 
		filter Filter,
	) Filter
}
