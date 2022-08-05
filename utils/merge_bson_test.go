package utils_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/operators/query"
	"github.com/0B1t322/MongoBuilder/utils"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFunc_MergeBson(t *testing.T) {
	require.Equal(
		t,
		bson.M{
			"field": bson.M{
				"$eq": 1,
				"$ne": 12,
			},
		},
		utils.MergeBsonM(
			bson.M{
				"field": bson.M {
					"$eq": 1,
				},
			},
			bson.M{
				"field": bson.M{
					"$ne": 12,
				},
			},
		),
	)

	require.Equal(
		t,
		bson.M{
			"$elemMatch": bson.M{
				"$gte": 80,
				"$lt": 85,
			},
		},
		bson.M{
			"$elemMatch": utils.MergeBsonM(
				query.SingleGTE(80),
				query.SingleLT(85),
			),
		},
	)

	require.Equal(
		t,
		bson.M{
			"result": bson.M{
				"$elemMatch": bson.M{
					"product": "xyz",
					"score": bson.M{
						"$gte": 8,
					},
				},
			},
		},
		bson.M{
			"result": bson.M{
				"$elemMatch": utils.MergeBsonM(
					query.EQField("product", "xyz"),
					query.GTE("score", 8),
				),
			},
		},
	)

	require.Equal(
		t,
		bson.M{},
		utils.MergeBsonM(),
	)
}