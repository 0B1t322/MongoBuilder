package aggregation

import (
	"github.com/0B1t322/MongoBuilder/object"
	"github.com/0B1t322/MongoBuilder/operators/sort"
	"github.com/0B1t322/MongoBuilder/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type AddFieldsArger interface {
	formatAddFieldsArg() bson.M
	AddField(field string, value interface{}) AddFieldsArger
	AddFieldArger(field string, value AddFieldsArger) AddFieldsArger
}

type addFieldsArger struct {
	builder object.ObjectBuilder
}

func (a addFieldsArger) formatAddFieldsArg() bson.M {
	return a.builder.Build()
}

func (a addFieldsArger) AddField(field string, value interface{}) AddFieldsArger {
	a.builder.AddField(field, value)
	return a
}

func (a addFieldsArger) AddFieldArger(field string, value AddFieldsArger) AddFieldsArger {
	a.builder.AddField(field, value.formatAddFieldsArg())
	return a
}

func AddFieldArg() AddFieldsArger {
	return addFieldsArger{builder: object.Object()}
}

func AddFields(args ...AddFieldsArger) bson.M {
	return bson.M{
		"$addFields": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, arg := range args {
					slice = append(slice, arg.formatAddFieldsArg())
				}
				return slice
			}()...,
		),
	}
}

type BucketArger interface {
	formatBucketArg() bson.M
	GroupBy(field interface{}) BucketArger
	AddBondary(boundary interface{}) BucketArger
	AddBondaries(boundaries ...interface{}) BucketArger
	Default(value interface{}) BucketArger

	// You can use objectBuilder to build the output value
	Output(value interface{}) BucketArger
}

type bucketArger struct {
	groupBy      interface{}
	boundaries   []interface{}
	defaultValue interface{}
	output       interface{}
}

func (b *bucketArger) formatBucketArg() bson.M {
	optionals := bson.M{}
	{
		if b.defaultValue != nil {
			optionals["default"] = b.defaultValue
		}

		if b.output != nil {
			optionals["output"] = b.output
		}
	}
	return bson.M{
		"$bucket": utils.MergeBsonM(
			bson.M{
				"groupBy":    b.groupBy,
				"boundaries": bson.A(b.boundaries),
			},
			optionals,
		),
	}
}

func (b *bucketArger) GroupBy(field interface{}) BucketArger {
	b.groupBy = field
	return b
}

func (b *bucketArger) AddBondary(boundary interface{}) BucketArger {
	b.boundaries = append(b.boundaries, boundary)
	return b
}

func (b *bucketArger) AddBondaries(boundaries ...interface{}) BucketArger {
	b.boundaries = append(b.boundaries, boundaries...)
	return b
}

func (b *bucketArger) Default(value interface{}) BucketArger {
	b.defaultValue = value
	return b
}

func (b *bucketArger) Output(value interface{}) BucketArger {
	b.output = value
	return b
}

func BucketArg() BucketArger {
	return &bucketArger{}
}

func Bucket(
	arg BucketArger,
) bson.M {
	return arg.formatBucketArg()
}

func Count(field string) bson.M {
	return bson.M{
		"$count": field,
	}
}

type GroupArger interface {
	formatGroupArg() bson.M
	// _id field in query
	GroupBy(expression interface{}) GroupArger
	AddField(field string, value interface{}) GroupArger
}

type groupArger struct {
	_id    interface{}
	fields object.ObjectBuilder
}

func (g *groupArger) formatGroupArg() bson.M {
	return bson.M{
		"$group": utils.MergeBsonM(
			bson.M{
				"_id": g._id,
			},
			g.fields.Build(),
		),
	}
}

func (g *groupArger) GroupBy(expression interface{}) GroupArger {
	g._id = expression
	return g
}

func (g *groupArger) AddField(field string, value interface{}) GroupArger {
	g.fields.AddField(field, value)
	return g
}

func GroupArg() GroupArger {
	return &groupArger{fields: object.Object()}
}

func Group(
	arg GroupArger,
) bson.M {
	return arg.formatGroupArg()
}

func Limit(limit int) bson.M {
	return bson.M{
		"$limit": limit,
	}
}

type LetArger interface {
	formatLetArg() bson.M
	Add(field string, value interface{}) LetArger
}

type letArger struct {
	fields object.ObjectBuilder
}

func (l *letArger) formatLetArg() bson.M {
	return l.fields.Build()
}

func (l *letArger) Add(field string, value interface{}) LetArger {
	l.fields.AddField(field, value)
	return l
}

func LetArg() LetArger {
	return &letArger{fields: object.Object()}
}

type LookupArger interface {
	formatLookupArg() bson.M
	From(from string) LookupArger
	LocalField(field string) LookupArger
	ForeignField(field string) LookupArger
	As(as string) LookupArger
	Let(arg LetArger) LookupArger
	SetPipelines(pipelines ...bson.M) LookupArger
}

type lookupArger struct {
	from         string
	localField   string
	foreignField string
	as           string
	let          bson.M
	pipelines    []bson.M
}

func (l *lookupArger) formatLookupArg() bson.M {
	optional := bson.M{}
	{
		if l.localField != "" {
			optional["localField"] = l.localField
		}

		if l.foreignField != "" {
			optional["foreignField"] = l.foreignField
		}

		if l.pipelines != nil {
			optional["pipeline"] = l.pipelines
		}

		if len(l.let) > 0 {
			optional["let"] = l.let
		}
	}
	return bson.M{
		"$lookup": utils.MergeBsonM(
			bson.M{
				"from": l.from,
				"as":   l.as,
			},
			optional,
		),
	}
}

func (l *lookupArger) From(from string) LookupArger {
	l.from = from
	return l
}

func (l *lookupArger) LocalField(field string) LookupArger {
	l.localField = field
	return l
}

func (l *lookupArger) ForeignField(field string) LookupArger {
	l.foreignField = field
	return l
}

func (l *lookupArger) As(as string) LookupArger {
	l.as = as
	return l
}

func (l *lookupArger) Let(arg LetArger) LookupArger {
	l.let = arg.formatLetArg()
	return l
}

func (l *lookupArger) SetPipelines(pipelines ...bson.M) LookupArger {
	l.pipelines = pipelines
	return l
}

func LookupArg() LookupArger {
	return &lookupArger{}
}

func Lookup(
	arg LookupArger,
) bson.M {
	return arg.formatLookupArg()
}

func Match(query interface{}) bson.M {
	return bson.M{
		"$match": query,
	}
}

type ProjectionArger interface {
	formatProjectionArg() bson.M
	AddField(field string, value interface{}) ProjectionArger
	AddFieldArger(field string, arg ProjectionArger) ProjectionArger
	ExcludeField(field string) ProjectionArger
	IncludeField(field string) ProjectionArger
}

type projectionArger struct {
	fields object.ObjectBuilder
}

func (p *projectionArger) formatProjectionArg() bson.M {
	return p.fields.Build()
}

func (p *projectionArger) AddField(field string, value interface{}) ProjectionArger {
	p.fields.AddField(field, value)
	return p
}

func (p *projectionArger) AddFieldArger(field string, arg ProjectionArger) ProjectionArger {
	p.fields.AddField(field, arg.formatProjectionArg())
	return p
}

func (p *projectionArger) ExcludeField(field string) ProjectionArger {
	p.fields.AddField(field, 0)
	return p
}

func (p *projectionArger) IncludeField(field string) ProjectionArger {
	p.fields.AddField(field, 1)
	return p
}

func ProjectionArg() ProjectionArger {
	return &projectionArger{fields: object.Object()}
}

func Projection(
	arg ProjectionArger,
) bson.M {
	return bson.M{
		"$project": arg.formatProjectionArg(),
	}
}

type MergeOnArger interface {
	formatMergeOnArg() bson.M
	On(on ...string) MergeOnArger
}

func MergeOnArg() MergeOnArger {
	return &mergeOnArger{
		on: []string{},
	}
}

type mergeOnArger struct {
	on []string
}

func (m *mergeOnArger) formatMergeOnArg() bson.M {
	return bson.M{
		"on": func() interface{} {
			if len(m.on) == 1 {
				return m.on[0]
			} else {
				array := bson.A{}
				for _, o := range m.on {
					array = append(array, o)
				}
				return array
			}
		}(),
	}
}

func (m *mergeOnArger) On(on ...string) MergeOnArger {
	m.on = append(m.on, on...)
	return m
}

type WhenMatchedArger interface {
	formatWhenMatchedArg() interface{}
	Replace() WhenMatchedArger
	KeepExisting() WhenMatchedArger
	Merge() WhenMatchedArger
	Fail() WhenMatchedArger
	Pipeline(stages ...bson.M) WhenMatchedArger
}

type whenMatchedArger struct {
	op interface{}
}

func WhenMatchedArg() WhenMatchedArger {
	arg := &whenMatchedArger{}
	return arg.Merge()
}

func (w *whenMatchedArger) formatWhenMatchedArg() interface{} {
	return w.op
}

func (w *whenMatchedArger) Replace() WhenMatchedArger {
	w.op = "replace"
	return w
}

func (w *whenMatchedArger) KeepExisting() WhenMatchedArger {
	w.op = "keepExisting"
	return w
}

func (w *whenMatchedArger) Merge() WhenMatchedArger {
	w.op = "merge"
	return w
}

func (w *whenMatchedArger) Fail() WhenMatchedArger {
	w.op = "fail"
	return w
}

func (w *whenMatchedArger) Pipeline(stages ...bson.M) WhenMatchedArger {
	w.op = stages
	return w
}

type WhenNotMatchedArger interface {
	formatWhenNotMatchedArg() interface{}
	Insert() WhenNotMatchedArger
	Discard() WhenNotMatchedArger
	Fail() WhenNotMatchedArger
}

type whenNotMatchedArger struct {
	op interface{}
}

func WhenNotMatchedArg() WhenNotMatchedArger {
	arg := &whenNotMatchedArger{}
	return arg.Insert()
}

func (w *whenNotMatchedArger) formatWhenNotMatchedArg() interface{} {
	return w.op
}

func (w *whenNotMatchedArger) Insert() WhenNotMatchedArger {
	w.op = "insert"
	return w
}

func (w *whenNotMatchedArger) Discard() WhenNotMatchedArger {
	w.op = "discard"
	return w
}

func (w *whenNotMatchedArger) Fail() WhenNotMatchedArger {
	w.op = "fail"
	return w
}

type MergeArger interface {
	formatMergeArg() bson.M
	IntoCollection(collection string) MergeArger
	IntoDatabase(database, collection string) MergeArger
	On(arg MergeOnArger) MergeArger
	Let(arg LetArger) MergeArger
	WhenMatched(arg WhenMatchedArger) MergeArger
	WhenNotMatched(arg WhenNotMatchedArger) MergeArger
}

type mergeArger struct {
	into           interface{}
	on             MergeOnArger
	let            LetArger
	whenMatched    WhenMatchedArger
	whenNotMatched WhenNotMatchedArger
}

func (m *mergeArger) formatMergeArg() bson.M {
	b := bson.M{
		"into": m.into,
	}

	if m.on != nil {
		b["on"] = m.on.formatMergeOnArg()["on"]
	}

	if m.let != nil {
		b["let"] = m.let.formatLetArg()
	}

	if m.whenMatched != nil {
		b["whenMatched"] = m.whenMatched.formatWhenMatchedArg()
	}

	if m.whenNotMatched != nil {
		b["whenNotMatched"] = m.whenNotMatched.formatWhenNotMatchedArg()
	}

	return b
}

func (m *mergeArger) IntoCollection(collection string) MergeArger {
	m.into = collection
	return m
}

func (m *mergeArger) IntoDatabase(database, collection string) MergeArger {
	m.into = bson.M{
		"db":   database,
		"coll": collection,
	}
	return m
}

func (m *mergeArger) On(arg MergeOnArger) MergeArger {
	m.on = arg
	return m
}

func (m *mergeArger) Let(arg LetArger) MergeArger {
	m.let = arg
	return m
}

func (m *mergeArger) WhenMatched(arg WhenMatchedArger) MergeArger {
	m.whenMatched = arg
	return m
}

func (m *mergeArger) WhenNotMatched(arg WhenNotMatchedArger) MergeArger {
	m.whenNotMatched = arg
	return m
}

func MergeArg() MergeArger {
	return &mergeArger{}
}

func Merge(arg MergeArger) bson.M {
	return bson.M{
		"$merge": arg.formatMergeArg(),
	}
}

func OutCollection(collection string) bson.M {
	return bson.M{
		"$out": collection,
	}
}

func OutDatabase(database, collection string) bson.M {
	return bson.M{
		"$out": bson.M{
			"db":   database,
			"coll": collection,
		},
	}
}

func Redact(expression interface{}) bson.M {
	return bson.M{
		"$redact": expression,
	}
}

func ReplaceRoot(newRoot interface{}) bson.M {
	return bson.M{
		"$replaceRoot": bson.M{
			"newRoot": newRoot,
		},
	}
}

func ReplaceWith(replaceDocument interface{}) bson.M {
	return bson.M{
		"$replaceWith": replaceDocument,
	}
}

func Sample(size interface{}) bson.M {
	return bson.M{
		"$sample": bson.M{
			"size": size,
		},
	}
}

func Set(args ...AddFieldsArger) bson.M {
	return bson.M{
		"$set": utils.MergeBsonM(
			func() (slice []bson.M) {
				for _, arg := range args {
					slice = append(slice, arg.formatAddFieldsArg())
				}
				return slice
			}()...,
		),
	}
}

func Skip(skip interface{}) bson.M {
	return bson.M{
		"$skip": skip,
	}
}

func Sort(args ...sort.SortArger) bson.M {
	return sort.Sort(args...)
}

func SortByCount(expression interface{}) bson.M {
	return bson.M{
		"$sortByCount": expression,
	}
}

type UnionWithArger interface {
	formatUnionWithArg() interface{}
	SetPipeline(pipelines ...bson.M) UnionWithArger
}

type unionWithArger struct {
	collection string
	pipeline   []bson.M
}

func UnionWithArg(collection string) UnionWithArger {
	return &unionWithArger{collection: collection}
}

func (u *unionWithArger) formatUnionWithArg() interface{} {
	if len(u.pipeline) == 0 {
		return u.collection
	} else {
		return bson.M{
			"coll":     u.collection,
			"pipeline": u.pipeline,
		}
	}
}

func (u *unionWithArger) SetPipeline(pipelines ...bson.M) UnionWithArger {
	u.pipeline = pipelines
	return u
}

func UnionWith(arg UnionWithArger) bson.M {
	return bson.M{
		"$unionWith": arg.formatUnionWithArg(),
	}
}

type UnsetArger interface {
	formatUnsetArg() interface{}
	Unset(fields ...string) UnsetArger
}

type unsetArger struct {
	fields bson.A
}

func UnsetArg(fields ...string) UnsetArger {
	arg := &unsetArger{}
	return arg.Unset(fields...)
}

func (u *unsetArger) formatUnsetArg() interface{} {
	if len(u.fields) == 1 {
		return u.fields[0]
	} else {
		return u.fields
	}
}

func (u *unsetArger) Unset(fields ...string) UnsetArger {
	u.fields = append(
		u.fields,
		func() (array bson.A) {
			for _, field := range fields {
				array = append(array, field)
			}
			return array
		}()...,
	)
	return u
}

func Unset(arg UnsetArger) bson.M {
	return bson.M{
		"$unset": arg.formatUnsetArg(),
	}
}

type UnwindArger interface {
	formatUnwindArg() interface{}
	PreserveNullAndEmptyArrays(bool) UnwindArger
	IncludeArrayIndex(string) UnwindArger
}

type unwindArger struct {
	path         string
	preserveNull bool
	includeIndex string
}

func UnwindArg(path string) UnwindArger {
	return &unwindArger{path: path}
}

func (u *unwindArger) formatUnwindArg() interface{} {
	b := bson.M{}
	{
		if u.preserveNull != false {
			b["preserveNullAndEmptyArrays"] = u.preserveNull
		}
		if u.includeIndex != "" {
			b["includeArrayIndex"] = u.includeIndex
		}
	}
	if len(b) == 0 {
		return u.path
	} else {
		b["path"] = u.path
		return b
	}
}

func (u *unwindArger) PreserveNullAndEmptyArrays(preserveNull bool) UnwindArger {
	u.preserveNull = preserveNull
	return u
}

func (u *unwindArger) IncludeArrayIndex(includeIndex string) UnwindArger {
	u.includeIndex = includeIndex
	return u
}

func Unwind(arg UnwindArger) bson.M {
	return bson.M{
		"$unwind": arg.formatUnwindArg(),
	}
}