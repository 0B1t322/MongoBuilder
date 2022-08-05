package update_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/operators/query"
	"github.com/0B1t322/MongoBuilder/operators/sort"
	"github.com/0B1t322/MongoBuilder/operators/update"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFunc_Update(t *testing.T) {
	t.Run(
		"CurrentDate",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$currentDate": bson.M{
						"lastModified": true,
						"cancallation.date": bson.M{
							"$type": "timestamp",
						},
						"date": bson.M{
							"$type": "date",
						},
					},
				},
				update.CurrentDate(
					update.CurrentDateArg(
						"lastModified",
						update.BooleanTypeSpecifaction(true),
					),
					update.CurrentDateArg(
						"cancallation.date",
						update.TimestampTypeSpecification(),
					),
					update.CurrentDateArg(
						"date",
						update.DateTypeSpecification(),
					),
				),
			)
		},
	)

	t.Run(
		"Inc",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$inc": bson.M{
						"quantity":      -2.0,
						"metrics.order": 1.0,
					},
				},
				update.Inc(
					update.IncArg("quantity", -2.0),
					update.IncArg("metrics.order", 1.0),
				),
			)
		},
	)

	t.Run(
		"Min",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$min": bson.M{
						"lowerScore": 250,
					},
				},
				update.Min(update.MinArg("lowerScore", 250)),
			)
		},
	)

	t.Run(
		"Max",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$max": bson.M{
						"hightScore": 250,
					},
				},
				update.Max(update.MaxArg("hightScore", 250)),
			)
		},
	)

	t.Run(
		"Mul",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$mul": bson.M{
						"price": 1,
					},
				},
				update.Mul(
					update.MulArg(
						"price",
						update.MulInt(1),
					),
				),
			)
		},
	)

	t.Run(
		"Rename",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$rename": bson.M{
						"nickname": "alias",
						"cell":     "mobile",
					},
				},
				update.Rename(
					update.RenameArg(
						"nickname",
						"alias",
					),
					update.RenameArg(
						"cell",
						"mobile",
					),
				),
			)
		},
	)

	t.Run(
		"Set",
		func(t *testing.T) {
			obj := struct {
				Field  string `bson:"field"`
				Field2 string `bson:"field_2"`
			}{
				Field:  "name",
				Field2: "a",
			}

			require.Equal(
				t,
				bson.M{
					"$set": bson.M{
						"nested.array": []string{"1", "2"},
						"valuefield":   12,
						"objField": bson.M{
							"field":   "name",
							"field_2": "a",
						},
					},
				},
				update.Set(
					update.SetArg(
						"nested.array",
						[]string{"1", "2"},
					),
					update.SetArg(
						"valuefield",
						12,
					),
					update.SetArgWithMarshalling(
						"objField",
						obj,
					),
				),
			)
		},
	)

	t.Run(
		"Empty Set",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$set": bson.M{},
				},
				update.Set(),
			)
		},
	)

	t.Run(
		"Unset",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$unset": bson.M{
						"field": "",
						"some":  "",
					},
				},
				update.Unset(
					update.UnsetArg("field"),
					update.UnsetArg("some"),
				),
			)
		},
	)

	t.Run(
		"AddToSet",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$addToSet": bson.M{
						"tags": "camera",
						"some": "a",
					},
				},
				update.AddToSet(
					update.AddToSetArg("tags", "camera"),
					update.AddToSetArg("some", "a"),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$addToSet": bson.M{
						"tags": bson.A{"camera", "phone"},
					},
				},
				update.AddToSet(
					update.AddToSetArg("tags", bson.A{"camera", "phone"}),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$addToSet": bson.M{
						"tags": bson.M{
							"$each": bson.A{"camera", "phone"},
						},
					},
				},
				update.AddToSet(
					update.AddEachToSetArg("tags", update.EachModifer(bson.A{"camera", "phone"})),
				),
			)
		},
	)

	t.Run(
		"Push",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$push": bson.M{
						"scores": 89,
					},
				},
				update.Push(
					update.PushArg(
						"scores",
						89,
					),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$push": bson.M{
						"scores": bson.M{
							"$each": bson.A{90, 92, 85},
						},
					},
				},
				update.Push(
					update.PushEachArg(
						"scores",
						update.EachModiferFromValues(90, 92, 85),
					),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$push": bson.M{
						"scores": bson.M{
							"$each": bson.A{
								bson.M{
									"wk":    5,
									"score": 8,
								},
								bson.M{
									"wk":    6,
									"score": 7,
								},
							},
							"$sort":     bson.D{{"score", -1}},
							"$slice":    2,
							"$position": 1,
						},
					},
				},
				update.Push(
					update.PushEachArg(
						"scores",
						update.EachModiferFromValues(
							bson.M{
								"wk":    5,
								"score": 8,
							},
							bson.M{
								"wk":    6,
								"score": 7,
							},
						),
						update.SortModifer(
							sort.Sort(
								sort.SortArg("score", sort.DESC()),
							),
						),
						update.SliceModifer(2),
						update.PositionModifer(1),
					),
				),
			)
		},
	)

	t.Run(
		"Pop",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$pop": bson.M{
						"scores": -1,
						"others": 1,
					},
				},
				update.Pop(
					update.PopArg(
						"scores",
						update.First(),
					),
					update.PopArg(
						"others",
						update.Last(),
					),
				),
			)
		},
	)

	t.Run(
		"Pull",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$pull": bson.M{
						"votes": bson.M{
							"$gte": 6,
						},
					},
				},
				update.Pull(
					update.PullArg(
						"votes",
						query.SingleGTE(6),
					),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$pull": bson.M{
						"results": bson.M{
							"score": 8,
							"item": "B",
						},
					},
				},
				update.Pull(
					update.PullArg(
						"results",
						bson.M{
							"score": 8,
							"item": "B",
						},
					),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$pull": bson.M{
						"results": bson.M{
							"answers": bson.M{
								"$elemMatch": bson.M{
									"q": 2,
									"a": bson.M{
										"$gte": 8,
									},
								},
							},
						},
					},
				},
				update.Pull(
					update.PullArg(
						"results",
						query.ElemMatch(
							"answers",
							query.EQField("q", 2),
							query.GTE("a", 8),
						),
					),
				),
			)
		},
	)

	t.Run(
		"PullAll",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$pullAll": bson.M{
						"scores": bson.A{0, 5},
					},
				},
				update.PullAll(
					update.PullAllArg(
						"scores",
						0, 5,
					),
				),
			)
		},
	)
}
