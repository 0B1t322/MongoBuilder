package sort

import "go.mongodb.org/mongo-driver/bson"

type SortOrder interface {
	getOrder() int
}

type ascOrder struct{}

func (ascOrder) getOrder() int {
	return 1
}

type descOrder struct{}

func (descOrder) getOrder() int {
	return -1
}

func ASC() SortOrder {
	return ascOrder{}
}

func DESC() SortOrder {
	return descOrder{}
}

// return
// 	{$sort: <sort-order>}
func SingleSort(order SortOrder) bson.M {
	return bson.M{"$sort": order.getOrder()}
}

type sortArg struct {
	field string
	order SortOrder
}

func (s sortArg) formatSortArg() bson.E {
	return bson.E{s.field, s.order.getOrder()}
}

type SortArger interface {
	formatSortArg() bson.E
}

func SortArg(field string, order SortOrder) sortArg {
	return sortArg{
		field: field,
		order: order,
	}
}

func Sort(args ...SortArger) bson.M {
	return bson.M{
		"$sort": bson.D(
			func() (slice []bson.E) {
				for _, arg := range args {
					slice = append(slice, arg.formatSortArg())
				}
				return slice
			}(),
		),
	}
}
