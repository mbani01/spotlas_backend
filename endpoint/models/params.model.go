package models

type Params struct {
	Latitude  float64 `validate:"required"min=-90,max=90`
	Longitude float64 `validate:"required"min=-180,max=180`
	Radius    float64 `validate:"required,gte=0"`
	Type      string  `validate:"required,oneof=circle square"`
}
