package main

import (
	"fmt"
	"reflect"
)

type Pack struct {
	a, b int
	res  map[string]interface{}
}

var P = Pack{
	a: 1,
	b: 2,
}

func (p Pack) Add(arg map[string]interface{}) Pack {
	sum := 0
	str1 := ""
	p.res = make(map[string]interface{})
	for k, v := range arg {
		if value, ok := v.(int); ok {
			sum += value
			p.res[k] = value
			str1 = str1 + fmt.Sprint(value, "+")
		}
	}
	str := str1 + fmt.Sprint(p.a, "+", p.b, "=", p.a+p.b+sum)

	p.res["a"] = p.a
	p.res["b"] = p.b
	p.res["s"] = str
	return p
}
func main() {
	pc := reflect.ValueOf(&P)
	m := pc.MethodByName("Add")
	param := make([]reflect.Value, 1)
	mp := map[string]interface{}{
		"c": 3,
		"d": 4,
		"e": 5,
	}
	var mp_interface interface{} = mp
	param[0] = reflect.ValueOf(mp_interface)
	v := m.Call(param)
	fmt.Println(v)
	pack, _ := v[0].Interface().(Pack)
	fmt.Println(pack)
	for k, v := range pack.res {
		fmt.Print(k, v)
	}
}
