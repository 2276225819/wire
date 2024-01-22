package wire

import (
	"fmt"
	"testing"
)

//	type s1[T1 any] struct {
//		v1 T1
//	}
//
//	type s2[T1, T2 any] struct {
//		v1 T1
//		v2 T1
//	}
//
// type fx[T1 any] IOpt
//
//	func Struct1[TS any, T1 any]() ProvidedValue[TS] {
//		return func(s *Container) (a any, e error) {
//			x := new(s1[T1])
//			x.v1, e = get[T1](s, "xx", nil) // f1.get(s)
//			if e != nil {
//				return nil, e
//			}
//			vv := *(*TS)(unsafe.Pointer(&x))
//			return vv, nil
//		}
//	}
//
// var _ = Struct1[*xx, yy]()
//
// type yy struct{ int }
//
//	type xx struct {
//		d yy
//		q int
//	}
type A struct{ *A }
type C struct{ int }
type D struct{}

func TestName(t *testing.T) {
	var x, y = int32(0), int(0)
	Build(
		Out(&x), Out(&y),
		Value(&C{1111111}),
		As[int32](Func(func(c *C) (int, error) {
			return 1423, fmt.Errorf("%v", c)
		})),
	)

	t.Log(x)
}
