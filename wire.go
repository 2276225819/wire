package wire

import (
	"reflect"
	"strconv"
)

// NewSet 新增组（兼容）
func NewSet(ss ...any) ProviderSet {
	return ss
}

// Value 注册值 (兼容）
func Value[T0 any](a T0) ProvidedValue[T0] {
	return func(s *Container) (any, error) {
		return a, nil
	}
}

// Bind 绑定类型 (兼容）
func Bind[typ any, to any](_ *typ, _ *to) ProvidedValue[typ] {
	pcs := stackPos(2)
	return func(s *Container) (any, error) {
		vv, err := get[to](s, "bind", pcs)
		if err != nil {
			return nil, err
		}
		val, ok := any(vv).(typ)
		if !ok {
			return nil, errorsNew("bind error", pcs)
		}
		return val, nil
	}
}

// InterfaceValue 绑定类型 (兼容）
func InterfaceValue[typ any, val any](_ *typ, v val) ProvidedValue[typ] {
	return func(s *Container) (any, error) {
		return v, nil
	}
}

// Struct 结构体解析 (兼容）
func Struct[typ any](obj *typ, fieldNames ...string) ProvidedStruct[*typ] {
	pcs := stackPos(1)
	return []IOpt{
		ProvidedValue[typ](func(s *Container) (_ any, _e error) {
			v, err := get[*typ](s, "struct get", pcs)
			if err != nil {
				return nil, err
			}
			return *v, nil
		}),
		ProvidedValue[*typ](func(s *Container) (any, error) {
			rfv := reflect.ValueOf(obj).Elem()
			rft := rfv.Type()
			xx := map[string]reflect.Value{} //, 0)
			if len(fieldNames) == 1 && fieldNames[0] == "*" {
				for i, l := 0, rfv.NumField(); i < l; i++ {
					f := rft.Field(i)
					if !f.IsExported() {
						continue
					}
					xx[rft.Field(i).Name] = rfv.Field(i)
				}
			} else {
				for _, ss := range fieldNames {
					f, _ := rft.FieldByName(ss)
					if f.IsExported() {
						return nil, errorsNew("un exported field "+ss, pcs)
					}
					xx[f.Name] = rfv.FieldByName(ss)
				}
			}
			for name, rf := range xx {
				KEYX := reflect.NewAt(rf.Type(), nil).Interface()
				v, err := s.get(KEYX, "struct "+name, pcs)
				if err != nil {
					return nil, err
				}
				rf.Set(reflect.ValueOf(v))
			}
			return obj, nil
		}),
	}
}

// FieldOf 结构体字段解析（兼容）
func FieldOf[typ any](obj *typ, fieldNames ...string) InitOption {
	pcs := stackPos(1)
	return func(s *Container) {
		rfv := reflect.TypeOf(obj).Elem()
		xx := map[string]reflect.StructField{}
		if len(fieldNames) == 1 && fieldNames[0] == "*" {
			for i, l := 0, rfv.NumField(); i < l; i++ {
				f := rfv.Field(i)
				if !f.IsExported() {
					continue
				}
				xx[f.Name] = f
			}
		} else {
			for _, ss := range fieldNames {
				f, ok := rfv.FieldByName(ss)
				if !ok {
					continue
				}
				if !f.IsExported() {
					continue
				}
				xx[f.Name] = f
			}
		}
		for n_, rf_ := range xx {
			n := n_
			rf := rf_
			KEYF := reflect.NewAt(rf.Type, nil).Interface()
			s.set(KEYF, func(s *Container) (any, error) {
				var KEYX **typ
				obj, err := s.get(KEYX, "", pcs)
				if err != nil {
					return nil, err
				}
				rfv := reflect.ValueOf(obj).Elem().FieldByName(n)
				return rfv.Interface(), nil
			})
		}
	}
}

// Func 构造函数解析 （兼容）
func Func[T any](x T) ProviderFunc {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Func {
		p := errorsNew("func error", stackPos(1))
		return ProviderFunc{
			key: reflect.NewAt(t.Out(0), nil).Interface(),
			fn: func(s *Container) (a any, e error) {
				return nil, p
			},
		}
	}
	fpp := v.Pointer()
	return ProviderFunc{
		key: reflect.NewAt(t.Out(0), nil).Interface(),
		fn: func(s *Container) (a any, e error) {
			ln := t.NumIn()
			xx := make([]reflect.Value, 0, ln)
			for i := 0; i < ln; i++ {
				KEYin := reflect.NewAt(t.In(i), nil).Interface()
				t1, err := s.get(KEYin, "param"+strconv.Itoa(i), funcPos(fpp))
				if err != nil {
					return nil, err
				}
				xx = append(xx, reflect.ValueOf(t1))
			}
			result := v.Call(xx)
			switch len(result) {
			case 1:
				return result[0].Interface(), nil
			case 2:
				e := result[1].Interface()
				if e != nil {
					if e_, ok := e.(error); ok {
						return nil, e_
					}
					return nil, errorsNew("return1 error", funcPos(fpp))
				}
				return result[0].Interface(), nil
			case 3:
				e := result[2].Interface()
				if e != nil {
					if e_, ok := e.(error); ok {
						return nil, e_
					}
					return nil, errorsNew("return2 error", funcPos(fpp))
				}
				s.cls(result[1].Interface().(func()))
				return result[0].Interface(), nil
			default:
				return nil, errorsNew("param error", funcPos(fpp))
			}
		},
	}
}

///////////////////////////////////////////////////////////////////

// AddSet 合并组
func AddSet(n *ProviderSet, ss ...any) (_ struct{}) {
	*n = append(*n, ss...)
	return
}

// As 类型转换
func As[Interface any](f IGet) ProvidedValue[Interface] {
	pcs := stackPos(1)
	return func(s *Container) (any, error) {
		vv, err := f.get(s)
		if err != nil {
			return nil, err
		}
		val, ok := vv.(Interface)
		if !ok {
			return nil, errorsNew("as error", pcs)
		}
		return val, nil
	}
}

// Out 解析返回值
func Out[T any](x *T) InjectOption {
	pcs := stackPos(1)
	return func(s *Container) (a any, e error) {
		ex, err := get[T](s, "out", pcs)
		if err != nil {
			return nil, err
		}
		*x = ex
		return nil, nil
	}
}

// StructI 结构体解析  Struct 泛型版本
func StructI[STRUCT any](fn func(*STRUCT)) ProvidedStruct[*STRUCT] {
	return []IOpt{
		ProvidedValue[STRUCT](func(s *Container) (_ any, _e error) {
			v, err := get[*STRUCT](s, "struct get", funcPos(fn))
			if err != nil {
				return nil, err
			}
			return *v, nil
		}),
		ProvidedValue[*STRUCT](func(s *Container) (x any, e error) {
			x, fs := __getFields(new(STRUCT), fn)
			for i, f_ := range fs {
				v, err := s.get(f_.GetKey(), "struct."+strconv.Itoa(i), funcPos(fn))
				if err != nil {
					return nil, err
				}
				f_.SetVal(v)
			}
			return x, nil
		}),
	}
}

// StructO 结构体字段提取 FieldOf 泛型版本
func StructO[STRUCT any](fn func(*STRUCT)) InitOption {
	return func(s *Container) {
		_, fs := __getFields[STRUCT](new(STRUCT), fn)
		for i_, f_ := range fs {
			i := i_
			s.set(f_.GetKey(), func(s *Container) (a any, e error) {
				nx, err := get[*STRUCT](s, "fieldOf", funcPos(fn))
				if err != nil {
					return nil, err
				}
				_, nfs := __getFields(nx, fn)
				for _, f__ := range nfs {
					s.set(f__.GetKey(), func(s *Container) (a any, e error) {
						return f__.GetVal(), nil
					})
				}
				return nfs[i].GetVal(), nil
			})
		}
	}
}

// S 字段定义 StructI / StructO
func S[T1 any](v *T1) {
	__structFields = append(__structFields, &fieldAttr[T1]{val: v})
}

//// Invoke 解析配置并返回值
//func Invoke(fn ...IOpt) ISet {
//	return InitOption(func(s *Container) {
//		for _, opt := range fn {
//			opt.init(s, key[afterBuild]{})
//		}
//	})
//}
