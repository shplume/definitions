package operation

import (
	"fmt"
	"reflect"
)

type ScalePipeline struct {
	Scale float32 `json:"scale"`
}

func (p *ScalePipeline) Apply(x interface{}) error {
	if reflect.TypeOf(x).Kind() != reflect.Ptr {
		return fmt.Errorf("[scale-pipeline] x is not a pointer")
	}

	kind := reflect.TypeOf(x).Elem().Kind()

	if kind != reflect.Float32 && kind != reflect.Int64 {
		return fmt.Errorf("[scale-pipeline] x is not a pointer to float32 or int64")
	}

	if kind == reflect.Float32 {
		v := reflect.ValueOf(x).Elem()
		v.SetFloat(v.Float() * float64(p.Scale))
	} else {
		v := reflect.ValueOf(x).Elem()
		v.SetInt(v.Int() * int64(p.Scale))
	}

	return nil
}

func NewScalePipeline(factor float32) *ScalePipeline {
	return &ScalePipeline{Scale: factor}
}
