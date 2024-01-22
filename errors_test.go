package wire

import (
	"strings"
	"testing"
)

func __FuncPos()                 {}
func __Stack1() error            { return errorsNew("xx", stackPos(0)) }
func __Stack2() (error, error)   { return __Stack1(), errorsNew("xx", stackPos(1).next()) }
func __Stack3() (error, error)   { return __Stack2() }
func __Offset1(pos errPos) error { return errorsNew("xx", pos) }
func __Offset2(pos errPos) error { return __Offset1(pos.next()) }
func __Offset3() (error, error)  { return __Offset2(offsetPos(0).next()), errorsNew("xx", offsetPos(0)) }

func TestFuncPos(t *testing.T) {
	x := funcPos(__FuncPos)
	e1 := errorsNew("xx", x)
	if !strings.Contains(e1.Error(), "errors_test.go:8") {
		t.Fatal()
	}
	x = x.next().next()
	e2 := errorsNew("xx", x)
	if !strings.Contains(e2.Error(), "errors_test.go:8") {
		t.Fatal()
	}
}

func TestStack(t *testing.T) {
	e1, e2 := __Stack3()
	if !strings.Contains(e1.Error(), "errors_test.go:9") {
		t.Fatal()
	}
	if !strings.Contains(e2.Error(), "errors_test.go:11") {
		t.Fatal()
	}
}

func TestOffset(t *testing.T) {
	e1, e2 := __Offset3()
	if !strings.Contains(e1.Error(), "errors_test.go:14") {
		t.Fatal()
	}
	if !strings.Contains(e2.Error(), "errors_test.go:14") {
		t.Fatal()
	}
}
