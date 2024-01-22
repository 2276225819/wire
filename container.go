package wire

type execOption func(s *Container) (a any, e error)
type running struct{}
type clean []func()

// Container 容器
type Container struct {
	cache   map[any]any
	imports map[any]map[any]execOption
	cleans  clean
}

func New(ss ProviderSet) *Container {
	last := &Container{
		cache:   map[any]any{},
		cleans:  []func(){},
		imports: map[any]map[any]execOption{},
	}
	ss.set(last)
	return last
}

func (c clean) cleanup() {
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
		e = ct.use(KEY)
		if e != nil {
			return nil, e
		}
		return vv, nil
	case running:
		return nil, errorsNew(txt+" loop dependency ", ep.next())
	case nil:
		return nil, errorsNew(txt+" get nil ", ep.next())
	default:
		return v, nil
	}
}

func (ct *Container) setExe(KEY any, v execOption) {
	ct.cache[KEY] = v
}

func (ct *Container) setAny(KEY any, v any) {
	ct.cache[KEY] = v
}

func (ct *Container) use(KEY any) error {
	ls := ct.imports[KEY]
	if len(ls) != 0 {
		for _, importFn := range ls {
			_, ee := importFn(ct) //丢弃值
			if ee != nil {
				return ee
			}
		}
	}
	return nil
}
