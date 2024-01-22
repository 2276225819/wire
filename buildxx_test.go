package wire

import (
	//"github.com/samber/do"
	//"go.uber.org/dig"
	//"go.uber.org/fx" 
	"testing"
)

func assert(v bool) {
	if !v {
		panic(nil)
	}
}

type X struct {
	y Y
	z Z
}
type Y struct {
	z Z
}
type Z struct {
	i int
}

func NewXX(y Y, z Z) X {
	return X{y, z}
}
func NewYY(z Z) Y {
	return Y{z}
}
func NewZZ() Z {
	return Z{231}
}

//
//func TestName1(t *testing.T) {
//	for i := 0; i < 1000000; i++ {
//		fx.New(
//			fx.NopLogger,
//			fx.Provide(NewXX),
//			fx.Provide(NewYY),
//			fx.Provide(NewZZ),
//			fx.Invoke(func(xx X) {
//				xx.z.i = 11
//			}),
//		)
//	}
//
//}
//
//func TestName3(t *testing.T) {
//	for i := 0; i < 1000000; i++ {
//		x := dig.New()
//		x.Provide(NewXX)
//		x.Provide(NewYY)
//		x.Provide(NewZZ)
//		_ = x.Invoke(func(xx X) {
//			xx.z.i = 11
//		})
//	}
//}
//
//func TestName14(t *testing.T) {
//	for i := 0; i < 1000000; i++ {
//		injector := do.New()
//
//		do.Provide(injector, func(di *do.Injector) (X, error) {
//			y, _ := do.Invoke[Y](di)
//			z, _ := do.Invoke[Z](di)
//			return NewXX(y, z), nil
//		})
//
//		do.Provide(injector, func(di *do.Injector) (Y, error) {
//			z, _ := do.Invoke[Z](di)
//			return NewYY(z), nil
//		})
//
//		do.Provide(injector, func(di *do.Injector) (Z, error) {
//			return NewZZ(), nil
//		})
//
//		xx, _ := do.Invoke[X](injector)
//		xx.z.i = 11
//	}
//
//}

func TestName2(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		xx := BuildA[X](NewXX, NewYY, NewZZ, NewZZ)
		xx.z.i = 11
	}
}

func TestName22(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		xx := BuildA[X](
			Func21(NewXX),
			Func11(NewYY),
			Func01(NewZZ),
			Func01(NewZZ),
		)
		xx.z.i = 11
	}
}

func TestName44(t *testing.T) {
	type Xx struct {
		int
	}

	cx := func(c *func() int) (x Xx) {
		*c = func() int {
			return x.int
		}
		return
	}

	var cb func() int
	xx := cx(&cb)
	xx.int = 2
	cb()
}
