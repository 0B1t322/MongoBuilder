package aggregation_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/aggregation"
	"github.com/0B1t322/MongoBuilder/object"
	op "github.com/0B1t322/MongoBuilder/operators/aggregation"
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
					"field1": "value1",
					"field2": "value2",
					"obj": bson.M{
						"field3": "value3",
					},
					"field4": bson.M{
						"$first": "$some",
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

	
}
