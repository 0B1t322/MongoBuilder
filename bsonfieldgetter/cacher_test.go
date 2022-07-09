package bsonfieldgetter_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/0B1t322/MongoBuilder/bsonfieldgetter"
	"github.com/stretchr/testify/require"
)

type structF struct {
	Name string `bson:"name1"`
}

type structG struct {
	Name string `bson:"name2"`
}

type structH struct {
	Name string `bson:"name3"`
}

func TestFunc_Cacher(t *testing.T) {
	c := bsonfieldgetter.GetCasher()

	wg := sync.WaitGroup{}
	wg.Add(6)

	var (
		fBuffer = make(chan *bsonfieldgetter.BsonFieldGetter, 2)
		gBuffer = make(chan *bsonfieldgetter.BsonFieldGetter, 2)
		hBuffer = make(chan *bsonfieldgetter.BsonFieldGetter, 2)
	)

	go func() {
		defer wg.Done()
		bg := c.Get(structF{})
		fBuffer <- bg
		require.Equal(
			t,
			"name1",
			bg.Get("Name"),
		)
	}()
	go func() {
		defer wg.Done()
		bg := c.Get(structF{})
		fBuffer <- bg
		require.Equal(
			t,
			"name1",
			bg.Get("Name"),
		)
	}()

	go func() {
		defer wg.Done()
		bg := c.Get(structG{})
		gBuffer <- bg
		require.Equal(
			t,
			"name2",
			bg.Get("Name"),
		)
	}()
	go func() {
		defer wg.Done()
		bg := c.Get(structG{})
		gBuffer <- bg
		require.Equal(
			t,
			"name2",
			bg.Get("Name"),
		)
	}()

	go func() {
		defer wg.Done()
		bg := c.Get(structH{})
		hBuffer <- bg
		require.Equal(
			t,
			"name3",
			bg.Get("Name"),
		)
	}()
	go func() {
		defer wg.Done()
		bg := c.Get(structH{})
		hBuffer <- bg
		require.Equal(
			t,
			"name3",
			bg.Get("Name"),
		)
	}()
	wg.Wait()

	require.Equal(
		t,
		reflect.ValueOf(
			func() *bsonfieldgetter.BsonFieldGetter {
				bg := <-fBuffer
				return bg
			}(),
		).Pointer(),
		reflect.ValueOf(
			func() *bsonfieldgetter.BsonFieldGetter {
				bg := <-fBuffer
				return bg
			}(),
		).Pointer(),
	)

	require.Equal(
		t,
		reflect.ValueOf(
			func() *bsonfieldgetter.BsonFieldGetter {
				bg := <-hBuffer
				return bg
			}(),
		).Pointer(),
		reflect.ValueOf(
			func() *bsonfieldgetter.BsonFieldGetter {
				bg := <-hBuffer
				return bg
			}(),
		).Pointer(),
	)

	require.Equal(
		t,
		reflect.ValueOf(
			func() *bsonfieldgetter.BsonFieldGetter {
				bg := <-gBuffer
				return bg
			}(),
		).Pointer(),
		reflect.ValueOf(
			func() *bsonfieldgetter.BsonFieldGetter {
				bg := <-gBuffer
				return bg
			}(),
		).Pointer(),
	)
}
