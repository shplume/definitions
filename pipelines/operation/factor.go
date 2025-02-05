package operation

import (
	"github.com/fengyfei/definitions/pipelines/defs"
)

type FactorPipeline struct {
	update bool

	factor   float32
	Type     uint32                    `json:"type"`
	Variable defs.KernelVariableHeader `json:"variable"`

	scale *ScalePipeline
	delta *DeltaPipeline
}

func (p *FactorPipeline) Apply(x interface{}) error {
	if p.Type == defs.PipelineTypeScale {
		if p.scale == nil || p.update {
			p.scale = NewScalePipeline(p.factor)

			p.update = false
		}

		return p.scale.Apply(x)
	}

	if p.Type == defs.PipelineTypeDelta {
		if p.delta == nil || p.update {
			p.delta = NewDeltaPipeline(p.factor)

			p.update = false
		}

		return p.delta.Apply(x)
	}
	return nil
}

func (p *FactorPipeline) Prepare(factor float32) {
	if p.factor != factor {
		p.update = true
		p.factor = factor
	}
}

func (p *FactorPipeline) GetVariable() defs.KernelVariableHeader {
	return p.Variable
}

func NewFactorPipeline(pipelineType uint32, stationID int32, domain int32, name string) *FactorPipeline {
	return &FactorPipeline{Type: pipelineType, Variable: defs.KernelVariableHeader{
		Station:      stationID,
		Domain:       domain,
		VariableName: name,
	}}
}
