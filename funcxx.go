package wire

func Func01[T0 any](fn func() T0) ProvidedValue[T0] {
	return newSet0x[T0](fn)
}
func Func02[T0 any](fn func() (T0, error)) ProvidedValue[T0] {
	return newSet0x[T0](fn)
}
func Func03[T0 any](fn func() (T0, func(), error)) ProvidedValue[T0] {
	return newSet0x[T0](fn)
}
func Func11[T0, T1 any](fn func(T1) T0) ProvidedValue[T0] {
	return newSet1x[T1, T0](fn)
}
func Func12[T0, T1 any](fn func(T1) (T0, error)) ProvidedValue[T0] {
	return newSet1x[T1, T0](fn)
}
func Func13[T0, T1 any](fn func(T1) (T0, func(), error)) ProvidedValue[T0] {
	return newSet1x[T1, T0](fn)
}
func Func21[T0, T1, T2 any](fn func(T1, T2) T0) ProvidedValue[T0] {
	return newSet2x[T1, T2, T0](fn)
}
func Func22[T0, T1, T2 any](fn func(T1, T2) (T0, error)) ProvidedValue[T0] {
	return newSet2x[T1, T2, T0](fn)
}
func Func23[T0, T1, T2 any](fn func(T1, T2) (T0, func(), error)) ProvidedValue[T0] {
	return newSet2x[T1, T2, T0](fn)
}
func Func31[T0, T1, T2, T3 any](fn func(T1, T2, T3) T0) ProvidedValue[T0] {
	return newSet3x[T1, T2, T3, T0](fn)
}
func Func32[T0, T1, T2, T3 any](fn func(T1, T2, T3) (T0, error)) ProvidedValue[T0] {
	return newSet3x[T1, T2, T3, T0](fn)
}
func Func33[T0, T1, T2, T3 any](fn func(T1, T2, T3) (T0, func(), error)) ProvidedValue[T0] {
	return newSet3x[T1, T2, T3, T0](fn)
}
func Func41[T0, T1, T2, T3, T4 any](fn func(T1, T2, T3, T4) T0) ProvidedValue[T0] {
	return newSet4x[T1, T2, T3, T4, T0](fn)
}
func Func42[T0, T1, T2, T3, T4 any](fn func(T1, T2, T3, T4) (T0, error)) ProvidedValue[T0] {
	return newSet4x[T1, T2, T3, T4, T0](fn)
}
func Func43[T0, T1, T2, T3, T4 any](fn func(T1, T2, T3, T4) (T0, func(), error)) ProvidedValue[T0] {
	return newSet4x[T1, T2, T3, T4, T0](fn)
}
func Func51[T0, T1, T2, T3, T4, T5 any](fn func(T1, T2, T3, T4, T5) T0) ProvidedValue[T0] {
	return newSet5x[T1, T2, T3, T4, T5, T0](fn)
}
func Func52[T0, T1, T2, T3, T4, T5 any](fn func(T1, T2, T3, T4, T5) (T0, error)) ProvidedValue[T0] {
	return newSet5x[T1, T2, T3, T4, T5, T0](fn)
}
func Func53[T0, T1, T2, T3, T4, T5 any](fn func(T1, T2, T3, T4, T5) (T0, func(), error)) ProvidedValue[T0] {
	return newSet5x[T1, T2, T3, T4, T5, T0](fn)
}

func FuncA0[T0 any](fn func() T0) ProvidedValue[T0] {
	return newSet0x[T0](fn)
}
func FuncA1[T1, T0 any](fn func(T1) T0) ProvidedValue[T0] {
	return newSet1x[T1, T0](fn)
}
func FuncA2[T1, T2, T0 any](fn func(T1, T2) T0) ProvidedValue[T0] {
	return newSet2x[T1, T2, T0](fn)
}
func FuncA3[T1, T2, T3, T0 any](fn func(T1, T2, T3) T0) ProvidedValue[T0] {
	return newSet3x[T1, T2, T3, T0](fn)
}
func FuncA4[T1, T2, T3, T4, T0 any](fn func(T1, T2, T3, T4) T0) ProvidedValue[T0] {
	return newSet4x[T1, T2, T3, T4, T0](fn)
}
func FuncA5[T1, T2, T3, T4, T5, T0 any](fn func(T1, T2, T3, T4, T5) T0) ProvidedValue[T0] {
	return newSet5x[T1, T2, T3, T4, T5, T0](fn)
}

func FuncAE0[T0 any](fn func() (T0, error)) ProvidedValue[T0] {
	return newSet0x[T0](fn)
}
func FuncAE1[T1, T0 any](fn func(T1) (T0, error)) ProvidedValue[T0] {
	return newSet1x[T1, T0](fn)
}
func FuncAE2[T1, T2, T0 any](fn func(T1, T2) (T0, error)) ProvidedValue[T0] {
	return newSet2x[T1, T2, T0](fn)
}
func FuncAE3[T1, T2, T3, T0 any](fn func(T1, T2, T3) (T0, error)) ProvidedValue[T0] {
	return newSet3x[T1, T2, T3, T0](fn)
}
func FuncAE4[T1, T2, T3, T4, T0 any](fn func(T1, T2, T3, T4) (T0, error)) ProvidedValue[T0] {
	return newSet4x[T1, T2, T3, T4, T0](fn)
}
func FuncAE5[T1, T2, T3, T4, T5, T0 any](fn func(T1, T2, T3, T4, T5) (T0, error)) ProvidedValue[T0] {
	return newSet5x[T1, T2, T3, T4, T5, T0](fn)
}

func FuncACE0[T0 any](fn func() (T0, func(), error)) ProvidedValue[T0] {
	return newSet0x[T0](fn)
}
func FuncACE1[T1, T0 any](fn func(T1) (T0, func(), error)) ProvidedValue[T0] {
	return newSet1x[T1, T0](fn)
}
func FuncACE2[T1, T2, T0 any](fn func(T1, T2) (T0, func(), error)) ProvidedValue[T0] {
	return newSet2x[T1, T2, T0](fn)
}
func FuncACE3[T1, T2, T3, T0 any](fn func(T1, T2, T3) (T0, func(), error)) ProvidedValue[T0] {
	return newSet3x[T1, T2, T3, T0](fn)
}
func FuncACE4[T1, T2, T3, T4, T0 any](fn func(T1, T2, T3, T4) (T0, func(), error)) ProvidedValue[T0] {
	return newSet4x[T1, T2, T3, T4, T0](fn)
}
func FuncACE5[T1, T2, T3, T4, T5, T0 any](fn func(T1, T2, T3, T4, T5) (T0, func(), error)) ProvidedValue[T0] {
	return newSet5x[T1, T2, T3, T4, T5, T0](fn)
}

func newSet0x[T0 any, Tx any](fn Tx) ProvidedValue[T0] {
	return func(s *Container) (a any, e error) {
		switch f := any(fn).(type) {
		case func() T0:
			a = f()
		case func() (T0, error):
			a, e = f()
		case func() (T0, func(), error):
			var cb func() = nil
			a, cb, e = f()
			if cb != nil {
				s.cleans = append(s.cleans, cb)
			}
		}
		return
	}
}

func newSet1x[T1, T0 any, Tx any](fn Tx) ProvidedValue[T0] {
	return func(s *Container) (a any, e error) {
		t1, e := get[T1](s, "param0", funcPos(fn))
		if e != nil {
			return
		}
		switch f := any(fn).(type) {
		case func(T1) T0:
			a = f(t1)
		case func(T1) (T0, error):
			a, e = f(t1)
		case func(T1) (T0, func(), error):
			var cb func() = nil
			a, cb, e = f(t1)
			if cb != nil {
				s.cleans = append(s.cleans, cb)
			}
		}
		return
	}
}

func newSet2x[T1, T2, T0 any, Tx any](fn Tx) ProvidedValue[T0] {
	return func(s *Container) (a any, e error) {
		t1, e := get[T1](s, "param0", funcPos(fn))
		if e != nil {
			return
		}
		t2, e := get[T2](s, "param1", funcPos(fn))
		if e != nil {
			return
		}
		switch f := any(fn).(type) {
		case func(T1, T2) T0:
			a = f(t1, t2)
		case func(T1, T2) (T0, error):
			a, e = f(t1, t2)
		case func(T1, T2) (T0, func(), error):
			var cb func() = nil
			a, cb, e = f(t1, t2)
			if cb != nil {
				s.cleans = append(s.cleans, cb)
			}
		}
		return
	}
}

func newSet3x[T1, T2, T3, T0 any, Tx any](fn Tx) ProvidedValue[T0] {
	return func(s *Container) (a any, e error) {
		t1, e := get[T1](s, "param0", funcPos(fn))
		if e != nil {
			return
		}
		t2, e := get[T2](s, "param1", funcPos(fn))
		if e != nil {
			return
		}
		t3, e := get[T3](s, "param2", funcPos(fn))
		if e != nil {
			return
		}
		switch f := any(fn).(type) {
		case func(T1, T2, T3) T0:
			a = f(t1, t2, t3)
		case func(T1, T2, T3) (T0, error):
			a, e = f(t1, t2, t3)
		case func(T1, T2, T3) (T0, func(), error):
			var cb func() = nil
			a, cb, e = f(t1, t2, t3)
			if cb != nil {
				s.cleans = append(s.cleans, cb)
			}
		}
		return
	}
}

func newSet4x[T1, T2, T3, T4, T0 any, Tx any](fn Tx) ProvidedValue[T0] {
	return func(s *Container) (a any, e error) {
		t1, e := get[T1](s, "param0", funcPos(fn))
		if e != nil {
			return
		}
		t2, e := get[T2](s, "param1", funcPos(fn))
		if e != nil {
			return
		}
		t3, e := get[T3](s, "param2", funcPos(fn))
		if e != nil {
			return
		}
		t4, e := get[T4](s, "param3", funcPos(fn))
		if e != nil {
			return
		}
		switch f := any(fn).(type) {
		case func(T1, T2, T3, T4) T0:
			a = f(t1, t2, t3, t4)
		case func(T1, T2, T3, T4) (T0, error):
			a, e = f(t1, t2, t3, t4)
		case func(T1, T2, T3, T4) (T0, func(), error):
			var cb func() = nil
			a, cb, e = f(t1, t2, t3, t4)
			if cb != nil {
				s.cleans = append(s.cleans, cb)
			}
		}
		return
	}
}

func newSet5x[T1, T2, T3, T4, T5, T0 any, Tx any](fn Tx) ProvidedValue[T0] {
	return func(s *Container) (a any, e error) {
		t1, e := get[T1](s, "param0", funcPos(fn))
		if e != nil {
			return
		}
		t2, e := get[T2](s, "param1", funcPos(fn))
		if e != nil {
			return
		}
		t3, e := get[T3](s, "param2", funcPos(fn))
		if e != nil {
			return
		}
		t4, e := get[T4](s, "param3", funcPos(fn))
		if e != nil {
			return
		}
		t5, e := get[T5](s, "param4", funcPos(fn))
		if e != nil {
			return
		}
		switch f := any(fn).(type) {
		case func(T1, T2, T3, T4, T5) T0:
			a = f(t1, t2, t3, t4, t5)
		case func(T1, T2, T3, T4, T5) (T0, error):
			a, e = f(t1, t2, t3, t4, t5)
		case func(T1, T2, T3, T4, T5) (T0, func(), error):
			var cb func() = nil
			a, cb, e = f(t1, t2, t3, t4, t5)
			if cb != nil {
				s.cleans = append(s.cleans, cb)
			}
		}
		return
	}
}
