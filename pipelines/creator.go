package pipelines

import (
	"fmt"

	"github.com/fengyfei/definitions/pipelines/bits"
	"github.com/fengyfei/definitions/pipelines/operation"
)

func NewPipelineWithType(t uint32, all map[string]interface{}) (Pipeline, error) {
	switch t {
	case PipelineTypeScale:
		if _, ok := all["scale"]; !ok {
			return nil, fmt.Errorf("pipeline, scale missing")
		}
		return operation.NewScalePipeline(float32(all["scale"].(float64))), nil

	case PipelineTypeDelta:
		if _, ok := all["delta"]; !ok {
			return nil, fmt.Errorf("pipeline delta missing")
		}
		return operation.NewDeltaPipeline(float32(all["delta"].(float64))), nil

	case PipelineTypeBits:
		if _, ok := all["start"]; !ok {
			return nil, fmt.Errorf("pipeline start missing")
		}
		if _, ok := all["len"]; !ok {
			return nil, fmt.Errorf("pipeline len missing")
		}
		return bits.NewBitsPipeline(uint32(all["start"].(float64)), uint32(all["len"].(float64))), nil
	}

	return nil, fmt.Errorf("pipeline type missing")
}
