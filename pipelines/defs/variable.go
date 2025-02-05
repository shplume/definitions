package defs

type KernelVariableHeader struct {
	Station      int32  `json:"station"`
	Domain       int32  `json:"domain"`
	VariableName string `json:"variable_name"`
}
