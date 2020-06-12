package entity

// ParameterInput ...
type ParameterInput struct {
	Path      string `json:"path"`
	Value     string `json:"value"`
	IsEncrypt bool   `json:"is_encrypt"`
}

// RawInput ...
type RawInput struct {
	Stage      string           `json:"stage"`
	Name       string           `json:"name"`
	Parameters []ParameterInput `json:"parameters"`
}

// RawOutput ...
type RawOutput struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
