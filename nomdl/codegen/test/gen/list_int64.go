// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// ListOfInt64

type ListOfInt64 struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfInt64() ListOfInt64 {
	return ListOfInt64{types.NewList(), &ref.Ref{}}
}

type ListOfInt64Def []int64

func (def ListOfInt64Def) New() ListOfInt64 {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = types.Int64(d)
	}
	return ListOfInt64{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfInt64) Def() ListOfInt64Def {
	d := make([]int64, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = int64(l.l.Get(i).(types.Int64))
	}
	return d
}

func ListOfInt64FromVal(val types.Value) ListOfInt64 {
	// TODO: Do we still need FromVal?
	if val, ok := val.(ListOfInt64); ok {
		return val
	}
	// TODO: Validate here
	return ListOfInt64{val.(types.List), &ref.Ref{}}
}

func (l ListOfInt64) InternalImplementation() types.List {
	return l.l
}

func (l ListOfInt64) Equals(other types.Value) bool {
	if other, ok := other.(ListOfInt64); ok {
		return l.Ref() == other.Ref()
	}
	return false
}

func (l ListOfInt64) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfInt64) Chunks() (futures []types.Future) {
	futures = append(futures, l.TypeRef().Chunks()...)
	futures = append(futures, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfInt64.
var __typeRefForListOfInt64 types.TypeRef

func (m ListOfInt64) TypeRef() types.TypeRef {
	return __typeRefForListOfInt64
}

func init() {
	__typeRefForListOfInt64 = types.MakeCompoundTypeRef(types.ListKind, types.MakePrimitiveTypeRef(types.Int64Kind))
	types.RegisterFromValFunction(__typeRefForListOfInt64, func(v types.Value) types.Value {
		return ListOfInt64FromVal(v)
	})
}

func (l ListOfInt64) Len() uint64 {
	return l.l.Len()
}

func (l ListOfInt64) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfInt64) Get(i uint64) int64 {
	return int64(l.l.Get(i).(types.Int64))
}

func (l ListOfInt64) Slice(idx uint64, end uint64) ListOfInt64 {
	return ListOfInt64{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfInt64) Set(i uint64, val int64) ListOfInt64 {
	return ListOfInt64{l.l.Set(i, types.Int64(val)), &ref.Ref{}}
}

func (l ListOfInt64) Append(v ...int64) ListOfInt64 {
	return ListOfInt64{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfInt64) Insert(idx uint64, v ...int64) ListOfInt64 {
	return ListOfInt64{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfInt64) Remove(idx uint64, end uint64) ListOfInt64 {
	return ListOfInt64{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfInt64) RemoveAt(idx uint64) ListOfInt64 {
	return ListOfInt64{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfInt64) fromElemSlice(p []int64) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.Int64(v)
	}
	return r
}

type ListOfInt64IterCallback func(v int64, i uint64) (stop bool)

func (l ListOfInt64) Iter(cb ListOfInt64IterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(int64(v.(types.Int64)), i)
	})
}

type ListOfInt64IterAllCallback func(v int64, i uint64)

func (l ListOfInt64) IterAll(cb ListOfInt64IterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(int64(v.(types.Int64)), i)
	})
}

type ListOfInt64FilterCallback func(v int64, i uint64) (keep bool)

func (l ListOfInt64) Filter(cb ListOfInt64FilterCallback) ListOfInt64 {
	nl := NewListOfInt64()
	l.IterAll(func(v int64, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}
