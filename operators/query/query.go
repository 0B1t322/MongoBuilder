package query

import (
	"github.com/0B1t322/MongoBuilder/operators/options"
	"github.com/0B1t322/MongoBuilder/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func singleOper(op string, value interface{}) bson.M {
	return bson.M{op: value}
}

// Comprassion

// return { $eq: value }
func SingleEQ(value interface{}) bson.M {
	return singleOper("$eq", value)
}

// return { field: { $eq: value } }
func EQ(field string, value interface{}) bson.M {
	return bson.M{field: SingleEQ(value)}
}

// return { field: value }
func EQField(field string, value interface{}) bson.M {
	return bson.M{field: value}
}

// return { $gt: value }
func SingleGT(value interface{}) bson.M {
	return singleOper("$gt", value)
}

// return { field: { $gt: value } 
func GT(field string, value interface{}) bson.M {
	return bson.M{field: SingleGT(value)}
}

// return { $gte: value }
func SingleGTE(value interface{}) bson.M {
	return singleOper("$gte", value)
}

// return { field: { $gte: value } 
func GTE(field string, value interface{}) bson.M {
	return bson.M{field: SingleGTE(value)}
}

// return { field: { $in: [values[0], values[1], ... values[n] ] } }
func In(field string, values ...interface{}) bson.M {
	return bson.M{field: bson.M{"$in": bson.A(values)}}
}

// return { $lt: value }
func SingleLT(value interface{}) bson.M {
	return singleOper("$lt", value)
}

// return { field: { $lt: value } 
func LT(field string, value interface{}) bson.M {
	return bson.M{field: SingleLT(value)}
}

// return { $lte: value }
func SingleLTE(value interface{}) bson.M {
	return singleOper("$lte", value)
}

// return { field: { $lte: value } 
func LTE(field string, value interface{}) bson.M {
	return bson.M{field: SingleLTE(value)}
}

// return { $ne: value }
func SingleNE(value interface{}) bson.M {
	return singleOper("$ne", value)
}

// return { field: { $ne: value } 
func NE(field string, value interface{}) bson.M {
	return bson.M{field: SingleNE(value)}
}

// return { field: { $nin: [values[0], values[1], ... values[n] ] } }
func Nin(field string, values ...interface{}) bson.M {
	return bson.M{field: bson.M{"$nin": bson.A(values)}}
}

// Logical

// return { $and: [ { expression[0] }, { expression[1] } , ... , { expression[n] } ] }
func And(expressions ...bson.M) bson.M {
	array := bson.A{}
	for _, expr := range expressions {
		array = append(array, expr)
	}
	return bson.M{"$and": array}
}

// return { $or: [ { expression[0] }, { expression[1] } , ... , { expression[n] } ] }
func Or(expressions ...bson.M) bson.M {
	array := bson.A{}
	for _, expr := range expressions {
		array = append(array, expr)
	}
	return bson.M{"$or": array}
}

// return { $nor: [ { expression[0] }, { expression[1] } , ... , { expression[n] } ] }
func Nor(expressions ...bson.M) bson.M {
	array := bson.A{}
	for _, expr := range expressions {
		array = append(array, expr)
	}
	return bson.M{"$nor": array}
}

// return { field: { $not: { operatorExpression } } }
func Not(field string, operatorExpression bson.M) bson.M {
	return bson.M{field: bson.M{"$not": operatorExpression}}
}

// Element

// return { field: { $exists: value } }
func Exists(field string, value bool) bson.M {
	return bson.M{field: bson.M{"$exists": value}}
}

// If only one type return { field: { $type: <BSON type> } }
// 
// If more than one type return { field: { $type: [ <BSON type1> , <BSON type2>, ... ] } }
func Type(field string, types ...bsontype.Type) bson.M {	
	if len(types) == 1 {
		return bson.M{field: bson.M{"$type": types[0].String()}}
	}

	typesArray := bson.A{}
	for _, t := range types {
		typesArray = append(typesArray, t.String())
	}
	return bson.M{field: bson.M{"$type": typesArray}}
}

// Evaluation

// return { $expr: { <expression> } }
// expression can be any valud aggragate expression
func Expr(aggragateExpression bson.M) bson.M {
	return bson.M{"$expr": aggragateExpression}
}

// TODO
// func JsonSchema()

// return { field: { $mod: [ divisor, remainder ] } }
func Mod(field string, divisor, remainder float64) bson.M {
	return bson.M{field: bson.M{"$mod": bson.A{divisor, remainder}}}
}



// return if given options { <field>: { $regex: 'pattern', $options: '<options>' } }
// if not return { <field>: { $regex: 'pattern' } }
func Regex(field string, pattern string, opts ...options.RegexOptions) bson.M {
	regexBson := bson.M{}
	{
		regexBson["$regex"] = pattern
		if builded := options.MergeRegexOptions(opts...).BuildRegexOption(); builded != "" {
			regexBson["$options"] = builded
		}
	}
	return bson.M{field: regexBson}
}

/*
Return: 
	{
		$text:
			{
				$search: <string>,
				$language: <string>,
				$caseSensitive: <bool>,
				$diacriticSensitive: <bool>
			}
	}
*/
func Text(search string, language string, caseSensitive, diacriticSensitive bool) bson.M {
	textParams := bson.M{
		"$search": search,
		"$caseSensitive": caseSensitive,
		"$diacriticSensitive": diacriticSensitive,
	}

	if language != "" {
		textParams["$language"] = language
	}

	return bson.M{
		"$text": textParams,
	}
}

// TODO
// func Where()

// Geospatial

// TODO all Geospatial


// Array

/*
The $all operator selects the documents where the value of a field is an array that contains all the specified elements. To specify an $all expression, use the following prototype: 

{ <field>: { $all: [ <value1> , <value2> ... ] } }
*/
func All(field string, values ...interface{}) bson.M {
	return bson.M{field: bson.M{"$all": bson.A(values)}}
}

/*
The $elemMatch operator matches documents that contain an array field with at least one element that matches all the specified query criteria.

{ <field>: { $elemMatch: { <query1>, <query2>, ... } } }
*/
func ElemMatch(field string, querys ...bson.M) bson.M {
	return bson.M{
		field: bson.M{
			"$elemMatch": utils.MergeBsonM(querys...),
		},
	}
}
/*
The $size operator matches any array with the number of elements specified by the argument.

{ field: { $size: size } }
*/
func Size(field string, size int) bson.M {
	return bson.M{field: singleOper("$size", size)}
}

type EQFieldArger interface {
	formatEQFieldArg() bson.M
}

type eqFieldArg struct {
	Field string
	Value interface{}
}

func (e eqFieldArg) formatEQFieldArg() bson.M {
	return EQField(e.Field, e.Value)
}

func EQFieldArg(field string, value interface{}) eqFieldArg {
	return eqFieldArg{
		Field: field,
		Value: value,
	}
}

func EQFieldArgWithAnothers(field string, anothers ...eqFieldArg) eqFieldArg {
	return eqFieldArg{
		Field: field,
		Value: utils.MergeBsonM(
			func () (slice []bson.M) {
				for _, another := range anothers {
					slice = append(slice, another.formatEQFieldArg())
				}
				return slice
			}()...,
		),
	}
}

func EQFields(args ...EQFieldArger) bson.M {
	return utils.MergeBsonM(
		func () (slice []bson.M) {
			for _, arg := range args {
				slice = append(slice, arg.formatEQFieldArg())
			}
			return slice
		}()...,
	)
}