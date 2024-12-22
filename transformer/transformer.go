package transformer

type Transformer struct {
	*RangeTransformer `json:"range"`
	*ScaleTransformer `json:"scale"`
}

func (s *Transformer) Apply(x *float32) bool {
	if s.RangeTransformer != nil {
		return s.RangeTransformer.Apply(x)
	}

	if s.ScaleTransformer != nil {
		return s.ScaleTransformer.Apply(x)
	}

	return false
}

func (s *Transformer) String() string {
	if s.RangeTransformer != nil {
		return s.RangeTransformer.String()
	}

	if s.ScaleTransformer != nil {
		return s.ScaleTransformer.String()
	}

	return ""
}
