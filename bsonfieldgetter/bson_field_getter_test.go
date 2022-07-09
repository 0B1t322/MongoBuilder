package bsonfieldgetter_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/bsonfieldgetter"
	"github.com/stretchr/testify/require"
)

func TestFunc_Func(t *testing.T) {
	b := bsonfieldgetter.NewBsonFieldGetter(
		[]struct {
			Name       string `bson:"name"`
			Age        int    `bson:"age"`
			Hide       int    `bson:"-"`
			NestedHide struct {
				Name string `bson:"name"`
			} `bson:"-"`

			Nested struct {
				Name string `bson:"name"`
				Time uint64 `bson:"time"`

				Obj struct {
					Name   string `bson:"name"`
					Inline struct {
						Some string `bson:"some"`
					} `bson:",inline"`
				} `bson:"obj"`
			} `bson:"nested"`

			Inline struct {
				Some string `bson:"some"`
			} `bson:",inline"`

			SliceField []struct {
				Name string `bson:"name"`
			} `bson:"sliceField"`

			SliceOfPtrField []*struct {
				Name string `bson:"name"`
			} `bson:"sliceOfPtrField"`
		}{},
	)

	require.Equal(
		t,
		map[string]string{
			"Name":            "name",
			"Age":             "age",
			"Nested":          "nested",
			"Some":            "some",
			"Nested.Name":     "nested.name",
			"Nested.Time":     "nested.time",
			"Nested.Obj":      "nested.obj",
			"Nested.Obj.Name": "nested.obj.name",
			"Nested.Obj.Some": "nested.obj.some",
			"SliceField":      "sliceField",
			"SliceField.Name": "sliceField.name",
			"SliceOfPtrField": "sliceOfPtrField",
			"SliceOfPtrField.Name": "sliceOfPtrField.name",
		},
		b.GetMap(),
	)
}
