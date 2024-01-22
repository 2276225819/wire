package wire

func Build(ss ...any) (_ struct{}) {
	ct := _new(ss)
	_, e := build[*Container](ct, false)
	if e != nil {
		panic(e)
		//panic(fmt.Sprintf("%+v", e))
	}
	return
}

func BuildE(ss ...any) (_ error) {
	ct := _new(ss)
	_, e := build[*Container](ct, false)
	if e != nil {
		return e
	}
	return nil
}

func BuildC(ss ...any) (func(), error) {
	ct := _new(ss)
	_, e := build[*Container](ct, false)
	if e != nil {
		panic(e)
	}
	return ct.cleans.cleanup, nil
}

func BuildCE(ss ...any) (func(), error) {
	ct := _new(ss)
	_, e := build[*Container](ct, true)
	if e != nil {
		return nil, e
	}
	return ct.cleans.cleanup, nil
}

func BuildA[ANY any](ss ...any) (null ANY) {
	ct := _new(ss)
	v, e := build[ANY](ct, false)
	if e != nil {
		panic(e)
		//panic(fmt.Sprintf("%+v", e))
	}
	return v
}

func BuildAE[ANY any](ss ...any) (null ANY, e error) {
	ct := _new(ss)
	v, e := build[ANY](ct, false)
	if e != nil {
		return null, e
	}
	return v, e
}

func BuildAC[ANY any](ss ...any) (null ANY, _ func()) {
	ct := _new(ss)
	v, e := build[ANY](ct, true)
	if e != nil {
		panic(e)
	}
	return v, ct.cleans.cleanup
}

func BuildACE[ANY any](ss ...any) (null ANY, _ func(), e error) {
	ct := _new(ss)
	v, e := build[ANY](ct, true)
	if e != nil {
		return null, nil, e
	}
	return v, ct.cleans.cleanup, e
}
