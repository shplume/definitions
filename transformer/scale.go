package transformer

import "fmt"

type ScaleTransformer struct {
	Scale float32 `json:"scale"`
}

func (s *ScaleTransformer) Apply(x *float32) bool {
	*x *= s.Scale
	return true
}

func (s *ScaleTransformer) String() string {
	return fmt.Sprintf("Scale: [%f]", s.Scale)
}
