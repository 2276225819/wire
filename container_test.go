package wire

import (
	"strings"
	"testing"
)

func TestGetNil(t *testing.T) {
	ct := _new(NewSet())
	_, e := get[string](ct, "", offsetPos(0))
	if !strings.Contains(e.Error(), "nil") {
		t.Fatal()
	}
}

func TestGetLoopDep(t *testing.T) {
	var s = ""
	ct := _new(NewSet(func(string) string { return "" }, Out(&s)))
	_, e := get[string](ct, "", offsetPos(0))
	if !strings.Contains(e.Error(), "loop dependency") {
		t.Fatal()
	}
}

func TestGetTypeErr(t *testing.T) {
	ct := _new(NewSet())
	ProvidedValue[int64](func(s *Container) (a any, e error) {
		return float32(22), nil
	}).set(ct)
	_, ee := get[int64](ct, "test", offsetPos(0))
	if !strings.Contains(ee.Error(), "type err") {
		t.Fatal()
	}
}
