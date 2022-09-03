package main

import (
	"fmt"
	"github.com/facebookgo/inject"
)

type MyService struct {
	Name string
}

type OrderService struct {
	Name string
	Ms   *MyService `inject:""`
}

type TestService struct {
	Name string
}

func main() {

	var graph inject.Graph

	var od OrderService = OrderService{}
	err := graph.Provide(
		&inject.Object{Value: &od},
		&inject.Object{Value: &TestService{}},

		&inject.Object{Value: &MyService{Name: "x"}},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	er := graph.Populate()
	if er != nil {
		fmt.Println(er)
		return
	}
	for _, obj := range graph.Objects() {
		fmt.Printf("object known: %v\n", obj)
	}

	fmt.Println(od)
	fmt.Println(od.Name)
	//fmt.Println(od.Ms.Name)

	//var odx OrderService
	//fmt.Println(odx.Ms)
}
