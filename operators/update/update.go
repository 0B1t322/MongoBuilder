package update

import (
	"github.com/0B1t322/MongoBuilder/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TypeSpecification interface {
	typeSpecification() interface{}
}

type booleanTypeSpecification bool

func (b booleanTypeSpecification) typeSpecification() interface{} {
	return bool(b)
}

func BooleanTypeSpecifaction(v bool) TypeSpecification {
	return booleanTypeSpecification(v)
}

type stringTypeSpecification string

func (t stringTypeSpecification) typeSpecification() interface{} {
	return bson.M{"$type": string(t)}
}

func TimestampTypeSpecification() TypeSpecification {
	return stringTypeSpecification("timestamp")
}

func DateTypeSpecification() TypeSpecification {
	return stringTypeSpecification("date")
}

type currentDateArg struct {
	Field             string
	TypeSpecification TypeSpecification
}

func (c currentDateArg) formatDateArg() bson.M {
	return bson.M{c.Field: c.TypeSpecification.typeSpecification()}
}

type CurrentDateArger interface {
	formatDateArg() bson.M
}

func CurrentDateArg(field string, typeSpecifacation TypeSpecification) CurrentDateArger {
	return currentDateArg{
		Field:             field,
		TypeSpecification: typeSpecifacation,
	}
}

/*
The $currentDate operator sets the value of a field to the current date, either as a Date or a timestamp. The default type is Date.

The $currentDate operator has the form:
	{ $currentDate: { <field1>: <typeSpecification1>, ... } }

<typeSpecification> can be either:

1. a boolean true to set the field value to the current date as a Date

2. a document { $type: "timestamp" } or { $type: "date" } which explicitly specifies the type. The operator is case-sensitive and accepts only the lowercase "timestamp" or the lowercase "date".

To specify a <field> in an embedded document or in an array, use dot notation.
*/
func CurrentDate(args ...CurrentDateArger) bson.M {
	return bson.M{
		"$currentDate": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatDateArg())
				}
				return slice
			}()...,
		),
	}
}

type incArg struct {
	Field  string
	Amount float64
}

func (i incArg) formatIncArg() bson.M {
	return bson.M{i.Field: i.Amount}
}

func IncArg(field string, amount float64) incArg {
	return incArg{
		Field:  field,
		Amount: amount,
	}
}

type IncArger interface {
	formatIncArg() bson.M
}

/*
The $inc operator increments a field by a specified value and has the following form:
	{ $inc: { <field1>: <amount1>, <field2>: <amount2>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Inc(args ...IncArger) bson.M {
	return bson.M{
		"$inc": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatIncArg())
				}
				return slice
			}()...,
		),
	}
}

type minArg struct {
	Field string
	Value interface{}
}

func (m minArg) formatMinArg() bson.M {
	return bson.M{m.Field: m.Value}
}

type MinArger interface {
	formatMinArg() bson.M
}

func MinArg(field string, value interface{}) minArg {
	return minArg{
		Field: field,
		Value: value,
	}
}

/*
The $min updates the value of the field to a specified value if the specified value is less than the current value of the field.
The $min operator can compare values of different types, using the BSON comparison order.
	{ $min: { <field1>: <value1>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Min(args ...MinArger) bson.M {
	return bson.M{
		"$min": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatMinArg())
				}
				return slice
			}()...,
		),
	}
}

type maxArg struct {
	Field string
	Value interface{}
}

func (m maxArg) formatMaxArg() bson.M {
	return bson.M{m.Field: m.Value}
}

type MaxArger interface {
	formatMaxArg() bson.M
}

func MaxArg(field string, value interface{}) maxArg {
	return maxArg{
		Field: field,
		Value: value,
	}
}

/*
The $max operator updates the value of the field to a specified value if the specified value is greater than the current value of the field.
The $max operator can compare values of different types, using the BSON comparison order.

The $max operator expression has the form:
	{ $max: { <field1>: <value1>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Max(args ...MaxArger) bson.M {
	return bson.M{
		"$max": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatMaxArg())
				}
				return slice
			}()...,
		),
	}
}

type MulNumber interface {
	getMulNumber() interface{}
}

type mulNumber struct {
	number interface{}
}

func (m mulNumber) getMulNumber() interface{} {
	return m.number
}

func MulInt32(v int32) mulNumber {
	return mulNumber{
		number: v,
	}
}

func MulInt(v int) mulNumber {
	return mulNumber{
		number: v,
	}
}

func MulInt64(v int64) mulNumber {
	return mulNumber{
		number: v,
	}
}

func MulFloat64(v float64) mulNumber {
	return mulNumber{
		number: v,
	}
}

func MulDecimal(h, l uint64) mulNumber {
	return mulNumber{
		number: primitive.NewDecimal128(h, l),
	}
}

func MulDecimalString(v string) mulNumber {
	d, _ := primitive.ParseDecimal128(v)
	return mulNumber{
		number: d,
	}
}

type mulArg struct {
	Field  string
	Number MulNumber
}

func (m mulArg) formatMulArg() bson.M {
	return bson.M{m.Field: m.Number.getMulNumber()}
}

type MulArger interface {
	formatMulArg() bson.M
}

func MulArg(field string, number MulNumber) mulArg {
	return mulArg{
		Field:  field,
		Number: number,
	}
}

/*
Multiply the value of a field by a number. To specify a $mul expression, use the following prototype:
	{ $mul: { <field1>: <number1>, ... } }
The field to update must contain a numeric value.

To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Mul(args ...MulArger) bson.M {
	return bson.M{
		"$mul": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatMulArg())
				}
				return slice
			}()...,
		),
	}
}

type renameArg struct {
	Field   string
	NewName string
}

func (r renameArg) formatRenameArg() bson.M {
	return bson.M{r.Field: r.NewName}
}

type RenameArger interface {
	formatRenameArg() bson.M
}

func RenameArg(field, newName string) renameArg {
	return renameArg{
		Field:   field,
		NewName: newName,
	}
}

/*
The $rename operator updates the name of a field and has the following form:
	{$rename: { <field1>: <newName1>, <field2>: <newName2>, ... } }
The new field name must differ from the existing field name. To specify a <field> in an embedded document, use dot notation.
*/
func Rename(args ...RenameArger) bson.M {
	return bson.M{
		"$rename": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatRenameArg())
				}
				return slice
			}()...,
		),
	}
}

type setArg struct {
	Field string
	Value interface{}
}

func (s setArg) formatSetArg() bson.M {
	return bson.M{s.Field: s.Value}
}

type SetArger interface {
	formatSetArg() bson.M
}

func SetArg(field string, value interface{}) setArg {
	return setArg{
		Field: field,
		Value: value,
	}
}

// Marshal value to bson and set
func SetArgWithMarshalling(field string, value interface{}) SetArger {
	out := bson.M{}
	{
		d, err := bson.Marshal(value)
		if err != nil {
			return nil
		}

		if err := bson.Unmarshal(d, &out); err != nil {
			return nil
		}
	}

	return setArg{
		Field: field,
		Value: out,
	}
}

/*
The $set operator replaces the value of a field with the specified value.

The $set operator expression has the following form:
	{ $set: { <field1>: <value1>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Set(args ...SetArger) bson.M {
	return bson.M{
		"$set": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatSetArg())
				}
				return slice
			}()...,
		),
	}
}

/*
If an update operation with upsert: true results in an insert of a document,
then $setOnInsert assigns the specified values to the fields in the document.

If the update operation does not result in an insert, $setOnInsert does nothing.
*/
func SetOnInsert(args ...SetArger) bson.M {
	return bson.M{
		"$setOnInsert": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatSetArg())
				}
				return slice
			}()...,
		),
	}
}

type unsetArg struct {
	Field string
}

func (u unsetArg) formatUnsetArg() bson.M {
	return bson.M{u.Field: ""}
}

type UnsetArger interface {
	formatUnsetArg() bson.M
}

func UnsetArg(field string) unsetArg {
	return unsetArg{
		Field: field,
	}
}

/*
The $unset operator deletes a particular field. Consider the following syntax:
	{ $unset: { <field1>: "", ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Unset(args ...UnsetArger) bson.M {
	return bson.M{
		"$unset": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, a := range args {
					slice = append(slice, a.formatUnsetArg())
				}
				return slice
			}()...,
		),
	}
}

// Modifers

type Modfier interface {
	formatModifier() bson.M
}

type PushModifier interface {
	Modfier
	pushModfier()
}

type AddToSetModifier interface {
	Modfier
	addToSetModifer()
}

type eachModifier struct {
	Value interface{}
}

func (e eachModifier) formatModifier() bson.M {
	return bson.M{"$each": e.Value}
}

func (eachModifier) pushModfier() {}

func (eachModifier) addToSetModifer() {}

// return each modifer with array values
// array can be any slice or bson.A
func EachModifer(array interface{}) eachModifier {
	return eachModifier{
		Value: array,
	}
}

// return each modifer with bson.A
func EachModiferFromValues(values ...interface{}) eachModifier {
	return eachModifier{
		Value: bson.A(values),
	}
}

func EmptyEachModifer() eachModifier {
	return eachModifier{
		Value: bson.A{},
	}
}

type positionModifer struct {
	Value int
}

func (positionModifer) pushModfier() {}

func (e positionModifer) formatModifier() bson.M {
	return bson.M{"$position": e.Value}
}

func (e positionModifer) addPushEachModifer(p pushEachArg) pushEachArg {
	p.Position = &e
	return p
}

func PositionModifer(number int) positionModifer {
	return positionModifer{
		Value: number,
	}
}

type sliceModfier struct {
	Value int
}

func (sliceModfier) pushModfier() {}

func (s sliceModfier) formatModifier() bson.M {
	return bson.M{"$slice": s.Value}
}

func (s sliceModfier) addPushEachModifer(p pushEachArg) pushEachArg {
	p.Slice = &s
	return p
}

func SliceModifer(num int) sliceModfier {
	return sliceModfier{
		Value: num,
	}
}

type sortModifer struct {
	Value bson.M
}

func (sortModifer) pushModfier() {}

func (s sortModifer) formatModifier() bson.M {
	return s.Value
}

func (s sortModifer) addPushEachModifer(p pushEachArg) pushEachArg {
	p.Sort = &s
	return p
}

// To build sort use sort package instead
func SortModifer(sort bson.M) sortModifer {
	return sortModifer{
		Value: sort,
	}
}

type AddToSetArger interface {
	formatAddToSetArg() bson.M
}

type addToSetArg struct {
	Field string
	Value interface{}
}

func (a addToSetArg) formatAddToSetArg() bson.M {
	return bson.M{a.Field: a.Value}
}

func AddToSetArg(field string, value interface{}) addToSetArg {
	return addToSetArg{
		Field: field,
		Value: value,
	}
}

type addEachToSetArg struct {
	Field string
	Each  eachModifier
}

func (a addEachToSetArg) formatAddToSetArg() bson.M {
	return bson.M{
		a.Field: a.Each.formatModifier(),
	}
}

func AddEachToSetArg(field string, each eachModifier) addEachToSetArg {
	return addEachToSetArg{
		Field: field,
		Each:  each,
	}
}

/*
The $addToSet operator adds a value to an array unless the value is already present, in which case $addToSet does nothing to that array.

The $addToSet operator has the form:
	{ $addToSet: { <field1>: <value1>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func AddToSet(args ...AddToSetArger) bson.M {
	return bson.M{
		"$addToSet": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, arg := range args {
					slice = append(slice, arg.formatAddToSetArg())
				}
				return slice
			}()...,
		),
	}
}

type PushArger interface {
	formatPushArg() bson.M
}

type pushArg struct {
	Field string
	Value interface{}
}

func (p pushArg) formatPushArg() bson.M {
	return bson.M{p.Field: p.Value}
}

func PushArg(field string, value interface{}) pushArg {
	return pushArg{
		Field: field,
		Value: value,
	}
}

type pushEachArg struct {
	Field    string
	Each     eachModifier
	Sort     *sortModifer
	Slice    *sliceModfier
	Position *positionModifer
}

func (p pushEachArg) formatPushArg() bson.M {
	ms := []bson.M{p.Each.formatModifier()}
	{
		if p.Sort != nil {
			ms = append(ms, p.Sort.formatModifier())
		}

		if p.Slice != nil {
			ms = append(ms, p.Slice.formatModifier())
		}

		if p.Position != nil {
			ms = append(ms, p.Position.formatModifier())
		}
	}
	return bson.M{
		p.Field: utils.MergeBsonM(
			ms...,
		),
	}
}

type PushEachArgModifers interface {
	addPushEachModifer(pushEachArg) pushEachArg
}

func PushEachArg(field string, each eachModifier, opts ...PushEachArgModifers) pushEachArg {
	arg := pushEachArg{
		Field: field,
		Each:  each,
	}

	for _, opt := range opts {
		arg = opt.addPushEachModifer(arg)
	}

	return arg
}

/*
The $push operator appends a specified value to an array.

The $push operator has the form:
	{ $push: { <field1>: <value1>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Push(args ...PushArger) bson.M {
	return bson.M{
		"$push": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, arg := range args {
					slice = append(slice, arg.formatPushArg())
				}
				return slice
			}()...,
		),
	}
}

type PopValuer interface {
	getPopValue() int
}

type firstPopValue struct{}

func (firstPopValue) getPopValue() int {
	return -1
}

type lastPopValue struct{}

func (lastPopValue) getPopValue() int {
	return 1
}

func First() PopValuer {
	return firstPopValue{}
}

func Last() PopValuer {
	return lastPopValue{}
}

type PopArger interface {
	formatPopArg() bson.M
}

type popArg struct {
	Field string
	Value PopValuer
}

func (p popArg) formatPopArg() bson.M {
	return bson.M{p.Field: p.Value.getPopValue()}
}

func PopArg(field string, value PopValuer) popArg {
	return popArg{
		Field: field,
		Value: value,
	}
}

/*
The $pop operator removes the first or last element of an array. Pass $pop a value of -1 to remove the first element of an array and 1 to remove the last element in an array.

The $pop operator has the form:
	{ $pop: { <field>: <-1 | 1>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Pop(args ...PopArger) bson.M {
	return bson.M{
		"$pop": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, arg := range args {
					slice = append(slice, arg.formatPopArg())
				}
				return slice
			}()...,
		),
	}
}

type PullArger interface {
	formatPullArg() bson.M
}

type pullArg struct {
	Field string
	Value interface{}
}

func (p pullArg) formatPullArg() bson.M {
	return bson.M{p.Field: p.Value}
}

func PullArg(field string, value interface{}) pullArg {
	return pullArg{
		Field: field,
		Value: value,
	}
}

/*
The $pull operator removes from an existing array all instances of a value or values that match a specified condition.

The $pull operator has the form:
	{ $pull: { <field1>: <value|condition>, <field2>: <value|condition>, ... } }
To specify a <field> in an embedded document or in an array, use dot notation.
*/
func Pull(args ...PullArger) bson.M {
	return bson.M{
		"$pull": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, arg := range args {
					slice = append(slice, arg.formatPullArg())
				}
				return slice
			}()...,
		),
	}
}

type PullAllArger interface {
	formatPullAllArg() bson.M
}

type pullAllArg struct {
	Field string
	Value interface{}
}

func (p pullAllArg) formatPullAllArg() bson.M {
	return bson.M{p.Field: p.Value}
}

func PullAllArg(field string, values ...interface{}) pullAllArg {
	return pullAllArg{
		Field: field,
		Value: bson.A(values),
	}
}

func PullAll(args ...PullAllArger) bson.M {
	return bson.M{
		"$pullAll": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, arg := range args {
					slice = append(slice, arg.formatPullAllArg())
				}
				return slice
			}()...,
		),
	}
}