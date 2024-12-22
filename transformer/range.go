package transformer

import "fmt"

type RangeTransformer struct {
	Min float32 `json:"min"`
	Max float32 `json:"max"`
}

func (s *RangeTransformer) Apply(x *float32) bool {
	if *x < s.Min {
		*x = s.Min
	} else if *x > s.Max {
		*x = s.Max
	}

	return true
}

func (s *RangeTransformer) String() string {
	return fmt.Sprintf("Range: [%f, %f]", s.Min, s.Max)
}
