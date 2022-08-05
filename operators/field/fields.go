package field

func Field(field string) string {
	return "$" + field
}

func Var(field string) string {
	return "$$" + field
}