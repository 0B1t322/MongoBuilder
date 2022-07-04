package types

type Type interface {
	StringIdentifier() string
	NumericIdentifier() int8
}

type typeBase int8

const (
	Dobule typeBase = iota + 1
	String
	Object
	Array
	BinaryData
	Undefined
	ObjectId
	Boolean
	Date
	Null
	Regex
	DBPointer
	JavaScript
	Symbol
	JavaScriptWithScope
	Int32
	Timestamp
	Int64
	Decimal128

	MinKey typeBase = -1
	MaxKey typeBase = 127
)

func (t typeBase) StringIdentifier() string {
	if(t == MinKey) {
		return "minKey"
	}

	if(t == MaxKey) {
		return "maxKey"
	}

	return []string{
		"double",
		"string",
		"object",
		"array",
		"binData",
		"undefined",
		"objectId",
		"bool",
		"date",
		"null",
		"regex",
		"dbPointer",
		"javascript",
		"symbol",
		"javascriptWithScope",
		"int",
		"timestamp",
		"long",
		"decimal",
	}[t]
}

func (t typeBase) NumericIdentifier() int8 {
	return int8(t)
}