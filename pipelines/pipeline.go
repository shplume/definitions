package pipelines

import (
	"encoding/json"
	"fmt"

	"github.com/fengyfei/definitions/pipelines/defs"
	"github.com/fengyfei/definitions/pipelines/operation"
)

var _ DynamicPipeline = (*operation.FactorPipeline)(nil)

type Pipeline interface {
	Apply(x interface{}) error
}

type DynamicPipeline interface {
	Pipeline
	Prepare(factor float32)
	GetVariable() defs.KernelVariableHeader
}

type PipelineDescriptor struct {
	Type     uint32 `json:"type"`
	Pipeline Pipeline
}

func (p *PipelineDescriptor) Apply(x interface{}) {
	p.Pipeline.Apply(x)
}

func (p *PipelineDescriptor) UnmarshalJSON(b []byte) error {
	var all map[string]interface{}

	if err := json.Unmarshal(b, &all); err != nil {
		return err
	}

	if _, ok := all["type"]; !ok {
		return fmt.Errorf("pipeline type missing")
	}

	p.Type = uint32(all["type"].(float64))

	pipeline, err := NewPipelineWithType(p.Type, all)
	if err != nil {
		return err
	}

	p.Pipeline = pipeline

	return nil
}

type PipelineDescriptors struct {
	Descriptors []PipelineDescriptor `json:"descriptors"`
}

func (p *PipelineDescriptors) Apply(x interface{}) {
	if len(p.Descriptors) == 0 {
		return
	}

	for _, d := range p.Descriptors {
		d.Apply(x)
	}
}
