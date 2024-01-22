package wire

type execOption func(s *Container) (a any, e error)
type running struct{}
type cleaner []func()

// Container 容器
type Container struct {
	cache   map[any]any
	imports map[any]map[any]execOption
	cleans  cleaner
}

func _new(ss ProviderSet) *Container {
	last := &Container{
		cache:   map[any]any{},
		cleans:  cleaner{},
		imports: map[any]map[any]execOption{},
	}
	ss.set(last)
	return last
}

func (c cleaner) cleanup() {
	for i := len(c) - 1; i >= 0; i-- {
		c[i]()
	}
}

func (ct *Container) get(KEY any, txt string, ep errPos) (any, error) {
	vx := ct.cache[KEY]
	switch v := vx.(type) {
	case execOption:
		ct.cache[KEY] = running{}
		vv, e := v(ct)
		if e != nil {
			return nil, e
		}
		ct.cache[KEY] = vv
		e = ct.init(KEY)
		if e != nil {
			return nil, e
		}
		return vv, nil
	case running:
		return nil, errorsNew(txt+" loop dependency ", ep.next())
	case nil:
		return nil, errorsNew(txt+" nil ", ep.next())
	default:
		return v, nil
	}
}

func (ct *Container) set(KEY any, v execOption) {
	ct.cache[KEY] = v
}

func (ct *Container) init(KEY any) error {
	ls := ct.imports[KEY]
	if len(ls) != 0 {
		for _, importFn := range ls {
			_, ee := importFn(ct)
			if ee != nil {
				return ee
			}
		}
	}
	return nil
}

func (ct *Container) use(KEY any, ao execOption) {
	if _, ok := ct.imports[KEY]; !ok {
		ct.imports[KEY] = map[any]execOption{}
	}
	ct.imports[KEY][&ao] = ao
}

func (ct *Container) cls(cb func()) {
	ct.cleans = append(ct.cleans, cb)
}

func get[T0 any](ct *Container, txt string, ep errPos) (null T0, _ error) {
	var KEY *T0
	vv, err := ct.get(KEY, txt, ep.next())
	if err != nil {
		return null, err
	}
	vvv, ok := vv.(T0)
	if !ok {
		return null, errorsNew(txt+" type err", ep.next())
	}
	return vvv, nil
}

func use[T0 any](ct *Container) error {
	var KEY *T0
	return ct.init(KEY)
}

func build[T any](ct *Container, must2cleanup bool) (null T, e error) {
	var KEY **Container
	ct.cache[KEY] = execOption(func(*Container) (any, error) { return ct, nil })
	v, e := get[T](ct, "build", offsetPos(2))
	if e != nil {
		ct.cleans.cleanup()
		return null, e
	}
	e = use[*Container](ct)
	if e != nil {
		ct.cleans.cleanup()
		return null, e
	}
	if !must2cleanup && len(ct.cleans) != 0 {
		return null, errorsNew("no cleanup", offsetPos(2))
	}
	return v, e
}
