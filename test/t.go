package main

import (
	"fmt"
	"reflect"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
}

func newStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}

func init() {
	registerType((*test)(nil))
}

type test struct {
	Name string
	Sex  int
}

func n(elem interface{}) reflect.Type {
	t := reflect.TypeOf(elem).Elem()
	return t
}
func main() {
	e := n((*test)(nil))
	x := reflect.New(e).Elem().Interface()

	fmt.Println(reflect.TypeOf(x))

	t, ok := x.(test)
	if !ok {
		return
	}
	t.Name = "i am test"
	fmt.Println(&t)
	fmt.Println(t, reflect.TypeOf(t))
}
