package operation

import (
	"fmt"
	"reflect"
)

type DeltaPipeline struct {
	Delta float32 `json:"delta"`
}

func (p *DeltaPipeline) Apply(x interface{}) error {
	if reflect.TypeOf(x).Kind() != reflect.Ptr {
		return fmt.Errorf("[delta-pipeline] x is not a pointer")
	}

	kind := reflect.TypeOf(x).Elem().Kind()

	if kind != reflect.Float32 && kind != reflect.Int64 {
		return fmt.Errorf("[delta-pipeline] x is not a pointer to float32 or int64")
	}

	if kind == reflect.Float32 {
		v := reflect.ValueOf(x).Elem()
		v.SetFloat(v.Float() + float64(p.Delta))
	} else {
		v := reflect.ValueOf(x).Elem()
		v.SetInt(v.Int() + int64(p.Delta))
	}

	return nil
}

func NewDeltaPipeline(delta float32) *DeltaPipeline {
	return &DeltaPipeline{Delta: delta}
}
