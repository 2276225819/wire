package wire

import (
	"runtime"
	"strconv"
	"unsafe"
)

type _newPtr struct {
	msg   string
	stack uintptr
}
type errPos interface {
	next() errPos
	stack() uintptr
}
type fp[T any] struct{ Func T }
type up struct{ p [1]uintptr }
type pp struct{ int }

func (u up) next() errPos    { return u }
func (u up) stack() uintptr  { return u.p[0] }
func (f fp[T]) next() errPos { return f }
func (f fp[T]) stack() (x uintptr) {
	switch v := any(f.Func).(type) {
	case uintptr:
		x = v + 1
	default:
		x = **(**uintptr)(unsafe.Pointer(&f.Func)) + 1
	}
	return
}
func (p pp) next() errPos { p.int += 1; return p }
func (p pp) stack() uintptr {
	var pcs [1]uintptr
	runtime.Callers(p.int+2, pcs[:])
	return pcs[0]
}
func stackPos(i int) errPos {
	var pcs [1]uintptr
	runtime.Callers(i+2, pcs[:])
	return up{pcs}
}
func offsetPos(i int) errPos {
	return pp{i + 1}
}
func funcPos[T any](ff T) errPos {
	return fp[T]{ff}
}

func errorsNew(message string, position errPos) (e *_newPtr) {
	return &_newPtr{message, position.stack()}
}

func (err *_newPtr) Error() string {
	f := Frame(err.stack)
	fn := runtime.FuncForPC(f.pc())
	file, line := fn.FileLine(f.pc())
	return err.msg + " " + file + ":" + strconv.Itoa(line)
}

//////////// go\pkg\mod\github.com\pkg\errors@v0.9.1\stackPos.go ///////////////////

// Frame represents a program counter inside a stackPos frame.
// For historical reasons if Frame is interpreted as a uintptr
// its value represents the program counter + 1.
type Frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f Frame) pc() uintptr { return uintptr(f) - 1 }
