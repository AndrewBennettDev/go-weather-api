package main

import (
	"encode/json"
	"fmt"
)

type InputData struct {
	Location	struct	`json:"location"`
	Current		struct	`json:"current"`
	WindMph		float32	`json:""`
	WindKph		float32	`json:""`
	WingDegree	float32	`json:""`
	WindDirection	float32	`json:""`
	PressureMb	float32	`json:""`
	PressureIn	float32	`json:""`
	PrecipMm	float32	`json:""`
	PrecipIn	float32	`json:""`
	Humidity	float32	`json:""`
	Cloud		float32	`json:""`
	FeelsLikeC	float32	`json:""`
	FeelsLikeF	float32	`json:""`
	VisKm		float32	`json:""`
	VisMiles	float32	`json:""`
	Uv		float32	`json:""`
	GustMph		float32	`json:""`
	GustKph		float32	`json:""`
}

type TransformedData struct {

}

func main() {

}
