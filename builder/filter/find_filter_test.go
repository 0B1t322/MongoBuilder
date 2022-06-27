package filter_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/builder/filter"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFunc_FindFilter(t *testing.T) {
	t.Parallel()

	t.Run(
		"EQField",
		func(t *testing.T) {
			expected := bson.M{
				"field_1": 1,
				"field_2": 2,
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().
					EQField("field_1", 1).
					EQField("field_2", 2).
					BSON(),

			)
		},
	)

	t.Run(
		"EQ",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$eq": 1,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().EQ("field", 1).BSON(),
			)
		},
	)

	t.Run(
		"NE",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$ne": 1,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().NEQ("field", 1).BSON(),
			)
		},
	)

	t.Run(
		"LT",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$lt": 1,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().LT("field", 1).BSON(),
			)
		},
	)

	t.Run(
		"LTE",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$lte": 1,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().LTE("field", 1).BSON(),
			)
		},
	)

	t.Run(
		"GT",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$gt": 1,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().GT("field", 1).BSON(),
			)
		},
	)

	t.Run(
		"GTE",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$gte": 1,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().GTE("field", 1).BSON(),
			)
		},
	)

	t.Run(
		"In",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$in": bson.A{"1", "2", "some"},
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().In("field", "1", "2", "some").BSON(),
			)
		},
	)

	t.Run(
		"NotIn",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$nin": bson.A{"1", "2", "some"},
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().NotIn("field", "1", "2", "some").BSON(),
			)
		},
	)

	t.Run(
		"Exist",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$exists": true,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().Exist("field").BSON(),
			)
		},
	)

	t.Run(
		"NotExist",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"$exists": false,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().NotExist("field").BSON(),
			)
		},
	)

	t.Run(
		"Field",
		func(t *testing.T) {
			expected := bson.M{
				"field": bson.M{
					"first": bson.M{
						"one": 1,
						"two": 2,
					},
					"some": 1,
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().
				Field(
					"field",
					filter.NewFindFilter().
						EQField("some", 1).
						Field(
							"first",
							filter.NewFindFilter().
								EQField("one", 1).
								EQField("two", 2),
						),
				).
				BSON(),
			)
		},
	)

	t.Run(
		"And",
		func(t *testing.T) {
			expected := bson.M{
				"$and": bson.A{
					bson.M{
						"field": bson.M{
							"$lte": 12,
						},
					},
					bson.M{
						"field": bson.M{
							"$gte": 16,
						},
					},
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().
					And(
						filter.NewFindFilter().LTE("field", 12),
						filter.NewFindFilter().GTE("field", 16),
					).
					BSON(),
			)
		},
	)

	t.Run(
		"Or",
		func(t *testing.T) {
			expected := bson.M{
				"$or": bson.A{
					bson.M{
						"field": bson.M{
							"$lte": 12,
						},
					},
					bson.M{
						"field": bson.M{
							"$gte": 16,
						},
					},
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().
					Or(
						filter.NewFindFilter().LTE("field", 12),
						filter.NewFindFilter().GTE("field", 16),
					).
					BSON(),
			)
		},
	)

	t.Run(
		"And_Or",
		func(t *testing.T) {
			expected := bson.M{
				"$or": bson.A{
					bson.M{
						"$and": bson.A{
							bson.M{
								"field_1": 1,
							},
						},
						"field_1": 12,
					},
				},
				"$and": bson.A{
					bson.M{
						"$or": bson.A{
							bson.M{
								"field_1": 1,
							},
						},
						"field_1": 12,
					},
				},
			}

			require.Equal(
				t,
				expected,
				filter.NewFindFilter().
					Or(
						filter.NewFindFilter().
							And(
								filter.NewFindFilter().EQField("field_1", 1),
							).
							EQField("field_1", 12),
					).
					And(
						filter.NewFindFilter().
						Or(
							filter.NewFindFilter().EQField("field_1", 1),
						).
						EQField("field_1", 12),
					).
					BSON(),
			)
		},
	)
}