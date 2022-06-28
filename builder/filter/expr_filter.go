package filter

import "go.mongodb.org/mongo-driver/bson"

type ExpressionFilter struct {
	raw bson.M
}

func (e *ExpressionFilter) BSON() bson.M {
	return e.raw
}

// return {"$or": [{...},...,{...}]}
func (e *ExpressionFilter) Or(filters ...Filter) Filter {
	var array bson.A
	{
		if or, find := e.raw[OR]; find {
			array = or.(bson.A)
		} else {
			array = bson.A{}
		}
	}

	for _, filter := range filters {
		array = append(array, filter.BSON())
	}

	e.raw[OR] = array
	return e
}

// return {"$and": [{...},...,{...}]}
func (e *ExpressionFilter) And(filters ...Filter) Filter {
	var array bson.A
	{
		if or, find := e.raw[AND]; find {
			array = or.(bson.A)
		} else {
			array = bson.A{}
		}
	}

	for _, filter := range filters {
		array = append(array, filter.BSON())
	}

	e.raw[AND] = array
	return e
}

// EQField returns {"field": value}
func (e *ExpressionFilter) EQField(field string, value interface{}) Filter {
	e.raw[field] = valueOrBSON(value)
	return e
}

func (e *ExpressionFilter) EQ(field string, value interface{}) Filter {
	e.raw[EQ] = bson.A{}
}

func (e *ExpressionFilter) NEQ(field string, value interface{}) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) LT(field string, value interface{}) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) LTE(field string, value interface{}) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) GT(field string, value interface{}) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) GTE(field string, value interface{}) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) In(field string, value ...interface{}) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) NotIn(field string, value ...interface{}) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) Exist(field string) Filter {
	panic("not implemented") // TODO: Implement
}

func (e *ExpressionFilter) NotExist(field string) Filter {
	panic("not implemented") // TODO: Implement
}

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
func (e *ExpressionFilter) Field(field string, filter Filter) Filter {
	panic("not implemented") // TODO: Implement
}

