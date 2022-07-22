package bsonfieldgetter_test

import (
	"testing"

	"github.com/0B1t322/MongoBuilder/bsonfieldgetter"
	"github.com/stretchr/testify/require"
)

func TestFunc_Tags(t *testing.T) {
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
			} `bson:"sliceOfPtrField,omitempty"`
		}{},
	)

	require.Equal(
		t,
		"name",
		b.Get("Name"),
	)

	require.Equal(
		t,
		"age",
		b.Get("Age"),
	)

	require.Equal(
		t,
		"nested",
		b.Get("Nested"),
	)

	require.Equal(
		t,
		"some",
		b.Get("Some"),
	)

	require.Equal(
		t,
		"nested.name",
		b.Get("Nested.Name"),
	)

	require.Equal(
		t,
		"nested.time",
		b.Get("Nested.Time"),
	)

	require.Equal(
		t,
		"nested.obj",
		b.Get("Nested.Obj"),
	)

	require.Equal(
		t,
		"nested.obj.name",
		b.Get("Nested.Obj.Name"),
	)

	require.Equal(
		t,
		"nested.obj.some",
		b.Get("Nested.Obj.Some"),
	)

	require.Equal(
		t,
		"sliceField",
		b.Get("SliceField"),
	)

	require.Equal(
		t,
		"sliceField.name",
		b.Get("SliceField.Name"),
	)

	require.Equal(
		t,
		"sliceOfPtrField",
		b.Get("SliceOfPtrField"),
	)

	require.Equal(
		t,
		"sliceOfPtrField.name",
		b.Get("SliceOfPtrField.Name"),
	)
}

type TypeF struct {
	Name   string   `bson:"name"`
	TypesS []*TypeS `bson:"typesS"`
}

type TypeS struct {
	SName  string   `bson:"sname"`
	TypesF []*TypeF `bson:"typesF"`
}

func TestFunc_RecursiveTypes(t *testing.T) {
	b := bsonfieldgetter.NewBsonFieldGetter(TypeF{})
	require.Equal(
		t,
		"typesS",
		b.Get("TypesS"),
	)
	require.Equal(
		t,
		"typesS.sname",
		b.Get("TypesS.SName"),
	)
	require.Equal(
		t,
		"typesS.typesF",
		b.Get("TypesS.TypesF"),
	)
	require.Equal(
		t,
		"typesS.typesF.name",
		b.Get("TypesS.TypesF.Name"),
	)
}
