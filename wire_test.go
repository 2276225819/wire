package wire

import (
	"testing"
)

type A struct{ A, D int }
type C struct{ V int }
type D struct{ V string }
type ID interface{ ff() }
type IC interface{ ff() }

func (i *D) ff() {}
func (i *C) ff() {}

func TestName(t *testing.T) {
	var x, y, z = int64(0), int(0), float32(0)
	Build(
		//x
		InterfaceValue(new(ID), &D{"222"}),
		NewSet(Struct(new(D), "*"), Value("111")),
		As[int64](Func(func(c IC, d D, id ID) (int64, error) { return 1423, nil })),
		NewSet(Out(&x)),

		//y
		NewSet(Bind(new(IC), new(*C)), Value(&C{1111111})),
		NewSet(FieldOf(new(C), "*"), Out(&y)),

		//z
		StructI(func(s *A) { S(&s.A); s.D = 100 }),
		StructO(func(s *A) { f := float32(s.D); S(&f) }),
		Out(&z),
	)
	if x != 1423 {
		t.Fatal()
	}
	if y != 1111111 {
		t.Fatal()
	}
	if z != 100 {
		t.Fatal()
	}

}
