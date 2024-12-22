package bits

import (
	"fmt"
	"reflect"
)

type BitsPipeline struct {
	Start uint32 `json:"start"`
	Len   uint32 `json:"len"`
}

func (p *BitsPipeline) Apply(x interface{}) error {
	if reflect.TypeOf(x).Kind() != reflect.Ptr {
		return fmt.Errorf("[bits-pipeline] x is not a pointer")
	}

	if reflect.TypeOf(x).Elem().Kind() != reflect.Int64 {
		return fmt.Errorf("[bits-pipeline] x is not a pointer to int64")
	}

	v := reflect.ValueOf(x).Elem()
	v.SetInt((v.Int() >> p.Start) & ((1 << p.Len) - 1))

	return nil
}

func NewBitsPipeline(start, len uint32) *BitsPipeline {
	return &BitsPipeline{
		Start: start,
		Len:   len,
	}
}
