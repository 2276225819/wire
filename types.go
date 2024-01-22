package wire

import (
	"sync"
)

// ProvidedValue 解析取值
type ProvidedValue[T0 any] execOption

func (p ProvidedValue[T0]) get(s *Container) (a any, e error) {
	return p(s)
}

func (p ProvidedValue[T0]) set(s *Container) {
	var KEY *T0
	s.set(KEY, execOption(p))
}

//func (p ProvidedValue[T0]) use(s *Container, useKey any) {
//	var KEY *T0
//	if _, ok := s.imports[useKey]; !ok {
//		s.imports[useKey] = map[any]execOption{}
//	}
//	s.imports[useKey][KEY] = func(ct *Container) (a any, e error) {
//		_, e = p(ct)
//		if e != nil {
//			return nil, e
//		}
//		e = use[T0](ct)
//		if e != nil {
//			return nil, e
//		}
//
//		return
//	}
//}

// InitOption 构造处理
type InitOption func(s *Container)

func (i InitOption) set(s *Container) {
	i(s)
}

//
//func (i execOption) use(s *Container, useKey any) {
//	var KEY = &i
//	if _, ok := s.imports[useKey]; !ok {
//		s.imports[useKey] = map[any]execOption{}
//	}
//	s.imports[useKey][KEY] = i
//}

// ProviderSet 构造分组
type ProviderSet []any

func (l ProviderSet) set(s *Container) {
	for _, fn := range l {
		if x, ok := fn.(ISet); ok {
			x.set(s)
		} else {
			Func(fn).set(s)
		}
	}
	return
}

// ProvidedStruct 解析组
type ProvidedStruct[T any] []IOpt

func (l ProvidedStruct[T]) get(s *Container) (a any, e error) {
	for _, i2 := range l {
		return i2.get(s)
	}
	return nil, nil
}

func (l ProvidedStruct[T]) set(s *Container) {
	for _, i2 := range l {
		i2.set(s)
	}
}

//func (l ProvidedStruct[T]) use(s *Container, key any) {
//	for _, i2 := range l {
//		i2.use(s, key)
//		return
//	}
//}

// ProviderFunc 反射解析函数
type ProviderFunc struct {
	key any
	fn  execOption
}

func (f ProviderFunc) set(s *Container) {
	s.set(f.key, f.fn)
}

func (f ProviderFunc) get(s *Container) (a any, e error) {
	return f.fn(s)
}

// InjectOption 解析后处理
type InjectOption execOption

func (ao InjectOption) set(s *Container) {
	var KEY **Container
	s.use(KEY, execOption(ao))
}

type IOpt interface {
	set(s *Container)
	get(s *Container) (a any, e error)
	//use(s *Container, key any)
}

type IGet interface {
	get(s *Container) (a any, e error)
}

type ISet interface {
	set(s *Container)
}

var __structLock sync.Mutex      //todo ThreadLocal
var __structFields []structField //todo ThreadLocal

func __getFields[STRUCT any](v *STRUCT, fn func(*STRUCT)) (*STRUCT, []structField) {
	__structLock.Lock()
	__structFields = nil
	fn(v)
	var fs = __structFields
	__structFields = nil
	__structLock.Unlock()
	return v, fs
}

type fieldAttr[T1 any] struct{ val *T1 }

type structField interface {
	SetVal(any)
	GetVal() any
	GetKey() any
}

func (x *fieldAttr[T1]) SetVal(_v any) {
	*x.val = _v.(T1)
}

func (x *fieldAttr[T1]) GetVal() any {
	return *x.val
}

func (x *fieldAttr[T1]) GetKey() any {
	var v *T1
	return v
}
