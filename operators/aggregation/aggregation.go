package aggregation

import (
	"github.com/0B1t322/MongoBuilder/operators/options"
	"github.com/0B1t322/MongoBuilder/operators/types"
	"github.com/0B1t322/MongoBuilder/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// if cond is true return defaultValue
func condDefaultOrValue(value interface{}, cond bool, defaultValye interface{}) interface{} {
	if cond {
		return defaultValye
	}
	return value
}

func defaultOrValue(value interface{}, defaultValue interface{}) interface{} {
	return condDefaultOrValue(
		value,
		value == nil,
		defaultValue,
	)
}

func condFuncDefaultOrValue(value interface{}, cond func() bool, defaultValue interface{}) interface{} {
	return condDefaultOrValue(value, cond(), defaultValue)
}

// Arithmetic Expression Operators

/*
Returns the absolute value of a number.

$abs has the following syntax:
	{ $abs: <number> }
The <number> expression can be any valid expression as long as it resolves to a number.
*/
func Abs(number interface{}) bson.M {
	return bson.M{
		"$abs": number,
	}
}

/*
Adds numbers together or adds numbers and a date. If one of the arguments is a date, $add treats the other arguments as milliseconds to add to the date.

The $add expression has the following syntax:
	{ $add: [ <expression1>, <expression2>, ... ] }
The arguments can be any valid expression as long as they resolve to either all numbers or to numbers and a date.
*/
func Add(expressions ...interface{}) bson.M {
	return bson.M{
		"$add": bson.A(expressions),
	}
}

/*
Returns the smallest integer greater than or equal to the specified number.

$ceil has the following syntax:
	{ $ceil: <number> }
The <number> expression can be any valid expression as long as it resolves to a number.
*/
func Ceil(number interface{}) bson.M {
	return bson.M{
		"$ceil": number,
	}
}

/*
Divides one number by another and returns the result. Pass the arguments to $divide in an array.

The $divide expression has the following syntax:
	{ $divide: [ <expression1>, <expression2> ] }
The first argument is the dividend, and the second argument is the divisor; i.e. the first argument is divided by the second argument.
The arguments can be any valid expression as long as they resolve to numbers.
*/
func Divide(dividend, divisor interface{}) bson.M {
	return bson.M{
		"$divide": bson.A{dividend, divisor},
	}
}

/*
Raises Euler's number (i.e. e ) to the specified exponent and returns the result.

$exp has the following syntax:
	{ $exp: <exponent> }
The <exponent> expression can be any valid expression as long as it resolves to a number.
*/
func Exp(exponent interface{}) bson.M {
	return bson.M{
		"$exp": exponent,
	}
}

/*
Returns the largest integer less than or equal to the specified number.

$floor has the following syntax:
	{ $floor: <number> }
The <number> expression can be any valid expression as long as it resolves to a number.
*/
func Floor(number interface{}) bson.M {
	return bson.M{
		"$floor": number,
	}
}

/*
Calculates the natural logarithm ln (i.e log e) of a number and returns the result as a double.

$ln has the following syntax:
	{ $ln: <number> }
The <number> expression can be any valid expression as long as it resolves to a non-negative number.
*/
func Ln(number interface{}) bson.M {
	return bson.M{
		"$ln": number,
	}
}

/*
Calculates the log of a number in the specified base and returns the result as a double.
$log has the following syntax:
	{ $log: [ <number>, <base> ] }
The <number> expression can be any valid expression as long as it resolves to a non-negative number.

The <base> expression can be any valid expression as long as it resolves to a positive number greater than 1.
*/
func Log(number, base interface{}) bson.M {
	return bson.M{
		"$log": bson.A{number, base},
	}
}

/*
Calculates the log base 10 of a number and returns the result as a double.

$log10 has the following syntax:
	{ $log10: <number> }
The <number> expression can be any valid expression as long as it resolves to a non-negative number.
*/
func Log10(number interface{}) bson.M {
	return bson.M{
		"$log10": number,
	}
}

/*
Divides one number by another and returns the remainder.

The $mod expression has the following syntax:
	{ $mod: [ <expression1>, <expression2> ] }
The first argument is the dividend, and the second argument is the divisor; i.e. first argument is divided by the second argument.
*/
func Mod(dividend, divisor interface{}) bson.M {
	return bson.M{
		"$mod": bson.A{dividend, divisor},
	}
}

/*
Raises a number to the specified exponent and returns the result. $pow has the following syntax:

{ $pow: [ <number>, <exponent> ] }

The <number> expression can be any valid expression as long as it resolves to a number.

The <exponent> expression can be any valid expression as long as it resolves to a number.

You cannot raise 0 to a negative exponent.
*/
func Pow(number, exponent interface{}) bson.M {
	return bson.M{
		"$pow": bson.A{number, exponent},
	}
}

/*
$round rounds a number to a whole integer or to a specified decimal place.

$round has the following syntax:
	{ $round : [ <number>, <place> ] }

	<number>:
		Number type
		Can be any valid expression that resolves to a number. Specifically, the expression must resolve to an integer, double, decimal, or long. $round returns an error if the expression resolves to a non-numeric data type.
	<place>:
		Integer type
		Optional Can be any valid expression that resolves to an integer between -20 and 100, exclusive. e.g. -20 < place < 100.
*/
func Round(number, place interface{}) bson.M {
	return bson.M{
		"$round": bson.A{number, defaultOrValue(place, 0)},
	}
}

/*
Calculates the square root of a positive number and returns the result as a double.

$sqrt has the following syntax:
	{ $sqrt: <number> }
The argument can be any valid expression as long as it resolves to a non-negative number.
*/
func Sqrt(number interface{}) bson.M {
	return bson.M{
		"$sqrt": number,
	}
}

/*
Subtracts two numbers to return the difference, or two dates to return the difference in milliseconds, or a date and a number in milliseconds to return the resulting date.

The $subtract expression has the following syntax:
	{ $subtract: [ <expression1>, <expression2> ] }
The second argument is subtracted from the first argument.

The arguments can be any valid expression as long as they resolve to numbers and/or dates.

To subtract a number from a date, the date must be the first argument.
*/
func Substract(first, second interface{}) bson.M {
	return bson.M{
		"$subtract": bson.A{first, second},
	}
}

/*
$trunc truncates a number to a whole integer or to a specified decimal place.
$sqrt has the following syntax:
	{ $trunc : [ <number>, <place> ] }

	<number>:
		Number type
		Can be any valid expression that resolves to a number. Specifically, the expression must resolve to an integer, double, decimal, or long.
	<place>:
		Integer type
		Optional Can be any valid expression that resolves to an integer between -20 and 100, exclusive. e.g. -20 < place < 100.
*/
func Trunc(number, place interface{}) bson.M {
	return bson.M{
		"$trunc": bson.A{number, defaultOrValue(place, 0)},
	}
}

/*
Work like
	func Trunc(number, place interface{}) bson.M
But place is zero in default
*/
func SingleTrunc(number interface{}) bson.M {
	return bson.M{
		"$trunc": number,
	}
}

// Array Expression Operators

/*
Returns the element at the specified array index.

$arrayElemAt has the following syntax:
	{ $arrayElemAt: [ <array>, <idx> ] }
The <array> expression can be any valid expression that resolves to an array.

	The <idx> expression can be any valid expression that resolves to an integer:
		If the <idx> expression resolves to zero or a positive integer, $arrayElemAt returns the element at the idx position, counting from the start of the array.
		If the <idx> expression resolves to a negative integer, $arrayElemAt returns the element at the idx position, counting from the end of the array.
If idx exceeds the array bounds, $arrayElemAt does not return a result.
*/
func ArrayElemAt(array, idx interface{}) bson.M {
	return bson.M{
		"$arrayElemAt": bson.A{array, idx},
	}
}

/*
Converts an array into a single document;

$arrayToObject has the following syntax:
	{ $arrayToObject: <expression> }
The <expression> can be any valid expression that resolves to an array of two-element arrays or array of documents that contains "k" and "v" fields.
*/
func ArrayToObject(expression interface{}) bson.M {
	return bson.M{
		"$arrayToObject": expression,
	}
}

/*
Concatenates arrays to return the concatenated array.

$concatArrays has the following syntax:
	{ $concatArrays: [ <array1>, <array2>, ... ] }
The <array> expressions can be any valid expression as long as they resolve to an array

If any argument resolves to a value of null or refers to a field that is missing, $concatArrays returns null.
*/
func ConcatArrays(arrayFirst, arraySecond interface{}, arrays ...interface{}) bson.M {
	return bson.M{
		"$concatArrays": func() (array bson.A) {
			array = append(array, arrayFirst, arraySecond)
			array = append(array, arrays...)
			return array
		}(),
	}
}

/*
Selects a subset of an array to return based on the specified condition. Returns an array with only those elements that match the condition. The returned elements are in the original order.

$filter has the following syntax:
	{ $filter: { input: <array>, as: <string>, cond: <expression> } }

input: An expression that resolves to an array.

as: A name for the variable that represents each individual element of the input array. If no name is specified, the variable name defaults to this.

cond: An expression that resolves to a boolean value used to determine if an element should be included in the output array. The expression references each element of the input array individually with the variable name specified in as.
*/
func Filter(
	input interface{},
	// if as equal "" set to "this"
	as string,
	cond interface{},
) bson.M {
	return bson.M{
		"$filter": bson.M{
			"input": input,
			"as":    condDefaultOrValue(as, as == "", "this"),
			"cond":  cond,
		},
	}
}

/*
$first has the following syntax:
	{ $first: <expression> }
The <expression> can be any valid expression as long as it resolves to an array, null or missing.
*/
func First(expression interface{}) bson.M {
	return bson.M{
		"$first": expression,
	}
}

/*
Returns a boolean indicating whether a specified value is in an array.

$in has the following operator expression syntax:
	{ $in: [ <expression>, <array expression> ] }
<expression> – Any valid expression expression.
<array expression> - Any valid expression that resolves to an array.
*/
func In(expression, arrayExpression interface{}) bson.M {
	return bson.M{
		"$in": bson.A{expression, arrayExpression},
	}
}

type IndexOfArrayOptionalParams interface {
	SetStart(start interface{}) IndexOfArrayOptionalParams
	SetEnd(end interface{}) IndexOfArrayOptionalParams

	getStart() interface{}
	getEnd() interface{}
}

type indexOfArrayOptionalParams struct {
	start interface{}
	end   interface{}
}

func (i indexOfArrayOptionalParams) merge(is ...IndexOfArrayOptionalParams) indexOfArrayOptionalParams {
	for _, ip := range is {
		i.start = ip.getStart()
		i.end = ip.getEnd()
	}
	return i
}

func (i indexOfArrayOptionalParams) getStart() interface{} {
	return i.start
}

func (i indexOfArrayOptionalParams) getEnd() interface{} {
	return i.end
}

func (i indexOfArrayOptionalParams) setStart(start interface{}) indexOfArrayOptionalParams {
	i.start = start
	return i
}

func (i indexOfArrayOptionalParams) setEnd(end interface{}) indexOfArrayOptionalParams {
	i.end = end
	return i
}

func (i indexOfArrayOptionalParams) SetStart(start interface{}) IndexOfArrayOptionalParams {
	return i.setStart(start)
}

func (i indexOfArrayOptionalParams) SetEnd(end interface{}) IndexOfArrayOptionalParams {
	return i.setEnd(end)
}

/*
Searches an array for an occurrence of a specified value and returns the array index (zero-based) of the first occurrence. If the value is not found, returns -1.
$indexOfArray has the following operator expression syntax:
	{ $indexOfArray: [ <array expression>, <search expression>, <start>, <end> ] }
*/
func IndexOfArray(
	arrayExpression,
	searchExpression interface{},
	optionalParams ...IndexOfArrayOptionalParams,
) bson.M {
	a := bson.A{arrayExpression, searchExpression}
	opt := indexOfArrayOptionalParams{}.
		merge(optionalParams...)

	if opt.start != nil {
		a = append(a, opt.start)
	}

	if opt.end != nil {
		a = append(a, opt.end)
	}

	return bson.M{
		"$indexOfArray": a,
	}
}

/*
Determines if the operand is an array. Returns a boolean.

$isArray has the following syntax:
	{ $isArray: [ <expression> ] }
*/
func IsArray(expression interface{}) bson.M {
	return bson.M{
		"$isArray": expression,
	}
}

/*
Returns the last element of an array.
The $last operator has the following syntax:
	{ $last: <expression> }
*/
func Last(expression interface{}) bson.M {
	return bson.M{
		"$last": expression,
	}
}

/*
Applies an expression to each item in an array and returns an array with the applied results.

The $map expression has the following syntax:
	{ $map: { input: <expression>, as: <string>, in: <expression> } }
*/
func Map(
	input interface{},
	as string,
	in interface{},
) bson.M {
	return bson.M{
		"$map": bson.M{
			"input": input,
			"as":    condDefaultOrValue(as, as == "", "this"),
			"in":    in,
		},
	}
}

/*
Converts a document to an array. The return array contains an element for each field/value pair in the original document. Each element in the return array is a document that contains two fields k and v:

	The k field contains the field name in the original document.

	The v field contains the value of the field in the original document.

$objectToArray has the following syntax:

	{ $objectToArray: <object> }

The <object> expression can be any valid expression as long as it resolves to a document object.
$objectToArray applies to the top-level fields of its argument.
If the argument is a document that itself contains embedded document fields, the $objectToArray does
not recursively apply to the embedded document fields.
*/
func ObjectToArray(object interface{}) bson.M {
	return bson.M{
		"$objectToArray": object,
	}
}

/*
Returns an array whose elements are a generated sequence of numbers. $range generates the sequence from the specified starting number by successively incrementing the starting number by the specified step value up to but not including the end point.

$range has the following operator expression syntax:
	{ $range: [ <start>, <end>, <non-zero step> ] }
*/
func RangeWithStep(start, end, step interface{}) bson.M {
	r := bson.A{start, end}
	if step != nil {
		r = append(r, step)
	}

	return bson.M{
		"$range": r,
	}
}

/*
Work like RangeWithStep, but without step.
*/
func Range(start, end interface{}) bson.M {
	return RangeWithStep(start, end, nil)
}

/*
New in version 3.4.

Applies an expression to each element in an array and combines them into a single value.

$reduce has the following syntax:
	{
		$reduce: {
			input: <array>,
			initialValue: <expression>,
			in: <expression>
		}
	}
*/
func Reduce(
	input interface{},
	initialValue interface{},
	in interface{},
) bson.M {
	return bson.M{
		"$reduce": bson.M{
			"input":        input,
			"initialValue": initialValue,
			"in":           in,
		},
	}
}

/*
Accepts an array expression as an argument and returns an array with the elements in reverse order.

$reverseArray has the following operator expression syntax:
	{ $reverseArray: <array expression> }

The argument can be any valid expression as long as it resolves to an array.
*/
func ReverseArray(
	arrayExpression interface{},
) bson.M {
	return bson.M{
		"$reverseArray": arrayExpression,
	}
}

/*
Counts and returns the total number of items in an array.

$size has the following syntax:
	{ $size: <expression> }

The argument for $size can be any expression as long as it resolves to an array.
*/
func Size(
	expression interface{},
) bson.M {
	return bson.M{
		"$size": expression,
	}
}

/*
Returns a subset of an array.

$slice has one of two syntax forms:

The following syntax returns elements from either the start or end of the array:

	{ $slice: [ <array>, <n> ] }

The following syntax returns elements from the specified position in the array:

	{ $slice: [ <array>, <position>, <n> ] }
*/
func Slice(
	array,
	n interface{},
	position ...interface{},
) bson.M {
	a := bson.A{array}
	if len(position) > 0 && position[0] != nil {
		a = append(a, position[0])
	}
	a = append(a, n)

	return bson.M{
		"$slice": a,
	}
}

type ZipArger interface {
	formatZipArg() bson.M
	AddInputs(...interface{}) ZipArger
	UseLongeestLength() ZipArger
	Defaults(interface{}) ZipArger
}

type zipArg struct {
	inputs           []interface{}
	useLongestLength bool
	defaults         interface{}
}

func (z zipArg) formatZipArg() bson.M {
	return bson.M{
		"inputs":           z.inputs,
		"useLongestLength": z.useLongestLength,
		"defaults":         z.defaults,
	}
}

func (z zipArg) AddInputs(inputs ...interface{}) ZipArger {
	z.inputs = append(z.inputs, inputs...)
	return z
}

func (z zipArg) UseLongeestLength() ZipArger {
	z.useLongestLength = true
	return z
}

func (z zipArg) Defaults(defaults interface{}) ZipArger {
	z.defaults = defaults
	return z
}

func ZipArg(inputs ...interface{}) ZipArger {
	return zipArg{inputs: inputs}
}

/*
Transposes an array of input arrays so that the first element of the output array would be an array containing, the first element of the first input array, the first element of the second input array, etc.

For example, $zip would transform [ [ 1, 2, 3 ], [ "a", "b", "c" ] ] into [ [ 1, "a" ], [ 2, "b" ], [ 3, "c" ] ].

$zip has the following syntax:
	{
		$zip: {
			inputs: [ <array expression1>,  ... ],
			useLongestLength: <boolean>,
			defaults:  <array expression>
		}
	}
*/
func Zip(
	arg ZipArger,
) bson.M {
	return bson.M{
		"$zip": arg.formatZipArg(),
	}
}

func And(
	args ...interface{},
) bson.M {
	return bson.M{
		"$and": bson.A(args),
	}
}

func Or(
	args ...interface{},
) bson.M {
	return bson.M{
		"$or": bson.A(args),
	}
}

func Not(
	arg interface{},
) bson.M {
	return bson.M{
		"$not": arg,
	}
}

func Cmp(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$cmp": bson.A{left, right},
	}
}

func EQ(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$eq": bson.A{left, right},
	}
}

func GT(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$gt": bson.A{left, right},
	}
}

func GTE(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$gte": bson.A{left, right},
	}
}

func LT(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$lt": bson.A{left, right},
	}
}

func LTE(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$lte": bson.A{left, right},
	}
}

func NE(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$ne": bson.A{left, right},
	}
}

func Cond(
	ifExpression interface{},
	thenExpression interface{},
	elseExpression interface{},
) bson.M {
	return bson.M{
		"$cond": bson.A{ifExpression, thenExpression, elseExpression},
	}
}

func IfNull(
	expression,
	replacement interface{},
) bson.M {
	return bson.M{
		"$ifNull": bson.A{expression, replacement},
	}
}

type SwithArger interface {
	formatSwitchArg() bson.M
	AddCase(caseExpression, then interface{}) SwithArger
	Default(exrpession interface{}) SwithArger
}

type switchArg struct {
	branches          []bson.M
	defaultExpression interface{}
}

func (s switchArg) formatSwitchArg() bson.M {
	return bson.M{
		"branches": s.branches,
		"default":  s.defaultExpression,
	}
}

func (s switchArg) AddCase(caseExpression, then interface{}) SwithArger {
	s.branches = append(s.branches, bson.M{"case": caseExpression, "then": then})
	return s
}

func (s switchArg) Default(exrpession interface{}) SwithArger {
	s.defaultExpression = exrpession
	return s
}

func SwitchArg() SwithArger {
	return switchArg{
		branches: make([]bson.M, 0),
	}
}

func Switch(
	arg SwithArger,
) bson.M {
	return bson.M{
		"$switch": arg.formatSwitchArg(),
	}
}

// TODO Date Expression Operators

func ToDate(
	expression interface{},
) bson.M {
	return bson.M{
		"$toDate": expression,
	}
}

func Literal(
	value interface{},
) bson.M {
	return bson.M{
		"$literal": value,
	}
}

func MergeObjects(
	docs ...interface{},
) bson.M {
	return bson.M{
		"$mergeObjects": condDefaultOrValue(bson.A(docs), len(docs) == 1, docs[0]),
	}
}

func SetField(
	field string,
	input interface{},
	value interface{},
) bson.M {
	return bson.M{
		"$setField": bson.M{
			"field": field,
			"input": input,
			"value": value,
		},
	}
}

func AllElementsTrue(
	expression interface{},
) bson.M {
	return bson.M{
		"$allElementsTrue": bson.A{expression},
	}
}

func AnyElemetsTrue(
	expression interface{},
) bson.M {
	return bson.M{
		"$anyElementTrue": bson.A{expression},
	}
}

func SetDifference(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$setDifference": bson.A{left, right},
	}
}

func SetEquals(
	expressions ...interface{},
) bson.M {
	return bson.M{
		"$setEquals": bson.A(expressions),
	}
}

func SetIntersection(
	arrays ...interface{},
) bson.M {
	return bson.M{
		"$setIntersection": bson.A(arrays),
	}
}

func SetIsSubset(
	left,
	right interface{},
) bson.M {
	return bson.M{
		"$setIsSubset": bson.A{left, right},
	}
}

func SetUnion(
	expressions ...interface{},
) bson.M {
	return bson.M{
		"$setUnion": bson.A(expressions),
	}
}

func Concat(
	expressions ...interface{},
) bson.M {
	return bson.M{
		"$concat": bson.A(expressions),
	}
}

func trimBody(
	input interface{},
	chars interface{},
) bson.M {
	b := bson.M{
		"input": input,
	}
	if chars != nil {
		b["chars"] = chars
	}
	return b
}

func LTrimChars(
	input interface{},
	chars interface{},
) bson.M {
	return bson.M{
		"$ltrim": trimBody(input, chars),
	}
}

func LTrim(
	input interface{},
) bson.M {
	return LTrimChars(input, nil)
}

func regexBody(
	input,
	regex interface{},
	opts ...options.RegexOptions,
) bson.M {
	b := bson.M{
		"input": input,
		"regex": regex,
	}

	if options := options.MergeRegexOptions(opts...).BuildRegexOption(); options != "" {
		b["options"] = options
	}
	return b
}

func RegexFind(
	input,
	regex interface{},
	opts ...options.RegexOptions,
) bson.M {
	return bson.M{
		"$regexFind": regexBody(input, regex, opts...),
	}
}

func RegexFindAll(
	input,
	regex interface{},
	opts ...options.RegexOptions,
) bson.M {
	return bson.M{
		"$regexFindAll": regexBody(input, regex, opts...),
	}
}

func RegexMatch(
	input,
	regex interface{},
	opts ...options.RegexOptions,
) bson.M {
	return bson.M{
		"$regexMatch": regexBody(input, regex, opts...),
	}
}

func replace(
	op string,
	input,
	find,
	replacement interface{},
) bson.M {
	return bson.M{
		op: bson.M{
			"input":       input,
			"find":        find,
			"replacement": replacement,
		},
	}
}

func ReplaceOne(
	input,
	find,
	replacement interface{},
) bson.M {
	return replace("$replaceOne", input, find, replacement)
}

func ReplaceAll(
	input,
	find,
	replacement interface{},
) bson.M {
	return replace("$replaceAll", input, find, replacement)
}

func RTtrim(
	input interface{},
) bson.M {
	return RTrimChars(input, nil)
}

func RTrimChars(
	input interface{},
	chars interface{},
) bson.M {
	return bson.M{
		"$rtrim": trimBody(input, chars),
	}
}

func Split(
	input,
	delimeter interface{},
) bson.M {
	return bson.M{
		"$split": bson.M{
			"input":     input,
			"delimeter": delimeter,
		},
	}
}

func StrLenBytes(
	input interface{},
) bson.M {
	return bson.M{
		"$strLenBytes": input,
	}
}

func StrLenCP(
	input interface{},
) bson.M {
	return bson.M{
		"$strLenCP": input,
	}
}

func StrCaseCMP(
	stringExpression,
	caseExpression interface{},
) bson.M {
	return bson.M{
		"$strcasecmp": bson.A{stringExpression, caseExpression},
	}
}

func SubStrBytes(
	stringExpression,
	byteIndex,
	byteCount interface{},
) bson.M {
	return bson.M{
		"$substrBytes": bson.A{stringExpression, byteIndex, byteCount},
	}
}

func ToLower(
	expression interface{},
) bson.M {
	return bson.M{
		"$toLower": expression,
	}
}

func ToString(
	expression interface{},
) bson.M {
	return bson.M{
		"$toString": expression,
	}
}

func Trim(
	input interface{},
) bson.M {
	return TrimChars(input, nil)
}

func TrimChars(
	input interface{},
	chars interface{},
) bson.M {
	return bson.M{
		"$trim": trimBody(input, chars),
	}
}

func ToUpper(
	expression interface{},
) bson.M {
	return bson.M{
		"$toUpper": expression,
	}
}

type ConvertOptionalsArger interface {
	convertOptionals()
	OnNull(interface{}) ConvertOptionalsArger
	OnError(interface{}) ConvertOptionalsArger
}

type convertOptionals struct {
	onError interface{}
	onNull  interface{}
}

func (convertOptionals) convertOptionals() {}

func (c convertOptionals) Merge(opts ...ConvertOptionalsArger) convertOptionals {
	for _, opt := range opts {
		c.onNull = opt.OnNull(c.onNull)
		c.onError = opt.OnError(c.onError)
	}
	return c
}

func (c convertOptionals) OnNull(v interface{}) ConvertOptionalsArger {
	c.onNull = v
	return c
}

func (c convertOptionals) OnError(v interface{}) ConvertOptionalsArger {
	c.onError = v
	return c
}

func ConvertOptionalsArgs() ConvertOptionalsArger {
	return convertOptionals{}
}

func Convert(
	input interface{},
	to types.Type,
	opts ...ConvertOptionalsArger,
) bson.M {
	optionals := bson.M{}
	{
		opt := convertOptionals{}.Merge(opts...)
		if opt.onNull != nil {
			optionals["onNull"] = opt.onNull
		}
		if opt.onError != nil {
			optionals["onError"] = opt.onError
		}
	}

	return bson.M{
		"$convert": utils.MergeBsonM(
			bson.M{
				"input": input,
				"to":    to.StringIdentifier(),
			},
			optionals,
		),
	}
}

func AddToSet(expression interface{}) bson.M {
	return bson.M{
		"$addToSet": expression,
	}
}

func Avg(
	expressions ...interface{},
) bson.M {
	return bson.M{
		"$avg": condDefaultOrValue(bson.A(expressions), len(expressions) == 1, expressions[0]),
	}
}

func Count() bson.M {
	return bson.M{
		"$count": bson.M{},
	}
}

func Push(expression interface{}) bson.M {
	return bson.M{
		"$push": expression,
	}
}

func Sum(expressions ...interface{}) bson.M {
	return bson.M{
		"$sum": condDefaultOrValue(bson.A(expressions), len(expressions) == 1, expressions[0]),
	}
}