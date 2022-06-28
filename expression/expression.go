package expression

type ExpressionType int

const (
	AGGRAGATION_VARIABLES ExpressionType = iota
	LITERAL
	EXPRESSION_OBJECTS
	OPERATOR_EXPRESSION

)

type Expression interface {
	ExpressionType() ExpressionType
}