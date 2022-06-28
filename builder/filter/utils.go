package filter

func valueOrBSON(v interface{}) interface{} {
	if b, ok := v.(ToBSON); ok {
		return b.BSON()
	}
	return v
}