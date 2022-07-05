package aggregation_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/aggregation"
	"github.com/0B1t322/MongoBuilder/object"
	op "github.com/0B1t322/MongoBuilder/operators/aggregation"
	"github.com/0B1t322/MongoBuilder/operators/query"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFunc_Aggregation(t *testing.T) {
	t.Run(
		"AddFields",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$addFields": bson.M{
						"field1": "value1",
						"field2": "value2",
						"obj": bson.M{
							"field3": "value3",
						},
						"field4": bson.M{
							"$first": "$some",
						},
					},
				},
				aggregation.AddFields(
					aggregation.AddFieldArg().
						AddField("field1", "value1").
						AddField("field2", "value2").
						AddFieldArger(
							"obj",
							aggregation.AddFieldArg().
								AddField("field3", "value3"),
						).
						AddField("field4", op.First("$some")),
				),
			)
		},
	)

	t.Run(
		"Bucket",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$bucket": bson.M{
						"groupBy": "$year_born",
						"boundaries": bson.A{
							1840,
							1850,
							1860,
							1870,
							1880,
						},
						"default": "Other",
						"output": bson.M{
							"count": bson.M{
								"$sum": 1,
							},
							"artist": bson.M{
								"$push": bson.M{
									"name": bson.M{
										"$concat": bson.A{"$first_name", " ", "$last_name"},
									},
									"year_born": "$year_born",
								},
							},
						},
					},
				},
				aggregation.Bucket(
					aggregation.BucketArg().
						GroupBy("$year_born").
						AddBondaries(1840, 1850, 1860, 1870, 1880).
						Default("Other").
						Output(
							object.Object().
								AddField("count", op.Sum(1)).
								AddField(
									"artist", 
									op.Push(
										object.Object().
											AddField("name", op.Concat("$first_name", " ", "$last_name")).
											AddField("year_born", "$year_born").
											Build(),
									),
							).
							Build(),
						),
				),
			)
		},
	)

	t.Run(
		"Group",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$group": bson.M{
						"_id": "$item",
						"totalSaleAmount": bson.M{
							"$sum": bson.M{
								"$multiply": bson.A{
									"$price",
									"$quantity",
							},
						},
					},
				},
			},
			aggregation.Group(
				aggregation.GroupArg().
					GroupBy("$item").
					AddField("totalSaleAmount", op.Sum(op.Multiply("$price", "$quantity"))),
				),
			)
		},
	)

	t.Run(
		"Lookup",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$lookup": bson.M{
						"from": "inventory",
						"localField": "item",
						"foreignField": "sku",
						"as": "inventory_docs",
					},
				},
				aggregation.Lookup(
					aggregation.LookupArg().
						From("inventory").
						LocalField("item").
						ForeignField("sku").
						As("inventory_docs"),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$lookup": bson.M{
						"from": "warehouses",
						"let": bson.M{
							"order_item": "$item",
							"order_qty": "$ordered",
						},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$and": bson.A{
											bson.M{
												"$eq": bson.A{
													"$stock_item",
													"$$order_item",
												},
											},
											bson.M{
												"$gte": bson.A{
													"$instock",
													"$$order_qty",
												},
											},
										},
									},
								},
							},
						},
						"as": "stockdata",
					},
				},
				aggregation.Lookup(
					aggregation.LookupArg().
						From("warehouses").
						Let(
							aggregation.LetArg().
								Add("order_item", "$item").
								Add("order_qty", "$ordered"),
						).
						SetPipelines(
							aggregation.Match(
								query.Expr(
									op.And(
										op.EQ("$stock_item", "$$order_item"),
										op.GTE("$instock", "$$order_qty"),
									),
								),
							),
						).
						As("stockdata"),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$lookup": bson.M{
						"from": "restaurants",
						"localField": "restaurant_name",
						"foreignField": "name",
						"let": bson.M{
							"orders_drink": "$drink",
						},
						"pipeline": []bson.M{
							{
								"$match": bson.M{
									"$expr": bson.M{
										"$in": bson.A{
											"$$orders_drink",
											"$beverages",
										},
									},
								},
							},
						},
						"as": "matches",
					},
				},
				aggregation.Lookup(
					aggregation.LookupArg().
						From("restaurants").
						LocalField("restaurant_name").
						ForeignField("name").
						Let(
							aggregation.LetArg().
								Add("orders_drink", "$drink"),
						).
						SetPipelines(
							aggregation.Match(
								query.Expr(
									op.In("$$orders_drink", "$beverages"),
								),
							),
						).
						As("matches"),
				),
			)
		},
	)

	t.Run(
		"Projection",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$project": bson.M{
						"_id": 0,
						"title": 1,
						"author": 1,
					},
				},
				aggregation.Projection(
					aggregation.ProjectionArg().
						ExcludeField("_id").
						IncludeField("title").
						IncludeField("author"),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$project": bson.M{
						"author": bson.M{
							"first": 0,
						},
						"lastModified": 0,
					},
				},
				aggregation.Projection(
					aggregation.ProjectionArg().
						AddFieldArger("author", aggregation.ProjectionArg().ExcludeField("first")).
						ExcludeField("lastModified"),
				),
			)
		},
	)

	t.Run(
		"Merge",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$merge": bson.M{
						"into": "newrestaurants",
						"on": bson.A{"date", "postcode"},
						"whenMatched": "replace",
						"whenNotMatched": "insert",
					},
				},
				aggregation.Merge(
					aggregation.MergeArg().
						IntoCollection("newrestaurants").
						On(aggregation.MergeOnArg().On("date", "postcode")).
						WhenMatched(aggregation.WhenMatchedArg().Replace()).
						WhenNotMatched(aggregation.WhenNotMatchedArg().Insert()),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$merge": bson.M{
						"into": bson.M{
							"db": "test-db",
							"coll": "some",
						},
						"on": "_id",
						"whenMatched": []bson.M{
							{
								"$addFields": bson.M{
									"thumbspup": bson.M{
										"$add": bson.A{"$thumbsup", "$$new.thumbsup"},
									},
									"thumbsdown": bson.M{
										"$add": bson.A{"$thumbsdown", "$$new.thumbsdown"},
									},
								},
							},
						},
						"whenNotMatched": "insert",
					},
				},
				aggregation.Merge(
					aggregation.MergeArg().
						IntoDatabase("test-db", "some").
						On(aggregation.MergeOnArg().On("_id")).
						WhenMatched(
							aggregation.WhenMatchedArg().
								Pipeline(
									aggregation.AddFields(
										aggregation.AddFieldArg().
											AddField("thumbspup", op.Add("$thumbsup", "$$new.thumbsup")).
											AddField("thumbsdown", op.Add("$thumbsdown", "$$new.thumbsdown")),
									),
								),
						).
						WhenNotMatched(aggregation.WhenNotMatchedArg().Insert()),
				),
			)
		},
	)

	t.Run(
		"Unwind",
		func(t *testing.T) {
			require.Equal(
				t,
				bson.M{
					"$unwind": "$tags",
				},
				aggregation.Unwind(
					aggregation.UnwindArg("$tags"),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$unwind": bson.M{
						"path": "$tags",
						"preserveNullAndEmptyArrays": true,
					},
				},
				aggregation.Unwind(
					aggregation.UnwindArg("$tags").
						PreserveNullAndEmptyArrays(true),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$unwind": bson.M{
						"path": "$tags",
						"includeArrayIndex": "tagsIndex",
					},
				},
				aggregation.Unwind(
					aggregation.UnwindArg("$tags").
						IncludeArrayIndex("tagsIndex"),
				),
			)

			require.Equal(
				t,
				bson.M{
					"$unwind": bson.M{
						"path": "$tags",
						"preserveNullAndEmptyArrays": true,
						"includeArrayIndex": "tagsIndex",
					},
				},
				aggregation.Unwind(
					aggregation.UnwindArg("$tags").
						PreserveNullAndEmptyArrays(true).
						IncludeArrayIndex("tagsIndex"),
				),
			)
		},
	)
}
