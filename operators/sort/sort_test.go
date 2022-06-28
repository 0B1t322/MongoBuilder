package sort_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/operators/sort"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)


func TestFunc_Sort(t *testing.T) {
	t.Run(
		"SingleSort",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$sort": -1,
				},
				sort.SingleSort(sort.DESC()),
			)
		},
	)

	t.Run(
		"Sort",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$sort": bson.D{
						{"field.f", 1},
						{"field.as", -1},
					},
				},
				sort.Sort(
					sort.SortArg("field.f", sort.ASC()),
					sort.SortArg("field.as", sort.DESC()),
				),
			)
		},
	)
}