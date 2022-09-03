package di

import (
	"github.com/facebookgo/inject"
)

var graph inject.Graph

func Inject(objects ...interface{}) {
	for _, obj := range objects {
		err := graph.Provide(&inject.Object{Value: obj})
		if err != nil {
			panic(err)
		}
	}

}
func Populate() {
	if err := graph.Populate(); err != nil {
		panic(err)
	}
}
