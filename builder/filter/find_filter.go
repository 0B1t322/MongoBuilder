package filter

import "go.mongodb.org/mongo-driver/bson"

type FindFilter struct {
	raw bson.M
}

func (f *FindFilter) BSON() bson.M {
	return f.raw
}

func NewFindFilter() *FindFilter {
	return &FindFilter{raw: bson.M{}}
}

// return {"$or": [{...},...,{...}]}
func (f *FindFilter) Or(filters ...Filter) Filter {
	var array bson.A
	{
		if or, find := f.raw[OR]; find {
			array = or.(bson.A)
		} else {
			array = bson.A{}
		}
	}

	for _, filter := range filters {
		array = append(array, filter.BSON())
	}

	f.raw[OR] = array
	return f
}

// return {"$and": [{...},...,{...}]}
func (f *FindFilter) And(filters ...Filter) Filter {
	var array bson.A
	{
		if or, find := f.raw[AND]; find {
			array = or.(bson.A)
		} else {
			array = bson.A{}
		}
	}

	for _, filter := range filters {
		array = append(array, filter.BSON())
	}
	
	f.raw[AND] = array
	return f
}

// EQField returns {"field": value}
func (f *FindFilter) EQField(
	field string,
	value interface{},
) Filter {
	f.raw[field] = value
	return f
}


// EQ returns {"field": {"$eq": value}}
func (f *FindFilter) EQ(field string, value interface{}) Filter {
	f.raw[field] = bson.M{EQ: value}
	return f
}

// NEQ returns {"field": {"$ne": value}}
func (f *FindFilter) NEQ(field string, value interface{}) Filter {
	f.raw[field] = bson.M{NE: value}
	return f
}

// LT returns {"field": {"$lt": value}}
func (f *FindFilter) LT(field string, value interface{}) Filter {
	f.raw[field] = bson.M{LT: value}
	return f
}

// LT returns {"field": {"$lte": value}}
func (f *FindFilter) LTE(field string, value interface{}) Filter {
	f.raw[field] = bson.M{LTE: value}
	return f
}

// GT returns {"field": {"$gt": value}}
func (f *FindFilter) GT(field string, value interface{}) Filter {
	f.raw[field] = bson.M{GT: value}
	return f
}

// GTE returns {"field": {"$gte": value}}
func (f *FindFilter) GTE(field string, value interface{}) Filter {
	f.raw[field] = bson.M{GTE: value}
	return f
}

/* 
In returns {"field": {"$in": [...]}}
Don't pass slice as single value
Do it like
	Example:
	strs := []string{"1", "2"}
	f := filter.NewFindFilter().In(
		"field",
		func () (slice []interface{}) {
			for _, elem := range strs {
				slice = append(slice, elem)
			}
			return slice
		}()...
	)
*/
func (f *FindFilter) In(field string, values ...interface{}) Filter {
	f.raw[field] = bson.M{IN: bson.A(values)}
	return f
}


func (f *FindFilter) NotIn(field string, values ...interface{}) Filter {
	f.raw[field] = bson.M{NIN: bson.A(values)}
	return f
}

func (f *FindFilter) Exist(field string) Filter {
	f.raw[field] = bson.M{EXISTS: true}
	return f
}

func (f *FindFilter) NotExist(field string) Filter {
	f.raw[field] = bson.M{EXISTS: false}
	return f
}

func (f *FindFilter) Field(
	field string, 
	filter Filter,
) Filter {
	f.raw[field] = filter.BSON()
	return f
}
