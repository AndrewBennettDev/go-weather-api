package main

import (
	"encode/json"
	"fmt"
)

type InputData struct {
	Location	struct	`json:"location"`
	Current		struct	`json:"current"`
	WindMph		float32	`json:"wind_mph"`
	WindKph		float32	`json:"wind_kph"`
	WingDegree	int8	`json:"wind_degree"`
	WindDirection	float32	`json:"wind_dir"`
	PressureMb	float32	`json:"pressure_mb"`
	PressureIn	float32	`json:"pressure_in"`
	PrecipMm	float32	`json:"precip_mm"`
	PrecipIn	float32	`json:"precip_in"`
	Humidity	int8	`json:"humidity"`
	Cloud		int8	`json:"cloud"`
	FeelsLikeC	float32	`json:"feelslike_C"`
	FeelsLikeF	float32	`json:"feelslike_F"`
	VisKm		int8	`json:"vis_km"`
	VisMiles	int8	`json:"vis_miles"`
	Uv		int8	`json:"uv"`
	GustMph		float32	`json:"gust_mph"`
	GustKph		float32	`json:"gust_kph"`
}

type Location struct {
	Name		string	`json:"name"`
	Region		string	`json:"region"`
	Country		string	`json:"country"`
	Latitude	float32	`json:"lat"`
	Longitude	float32	`json:"lon"`
	TimeZone	string	`json:"tz_id"`
	LocalTimeEpoch	int32	`json:"localtime_epoch"`
	LocalTime	string	`json:"localtime"`
}

type Current struct {
	LastUpdateEpoch	int32	`json:"last_updated_epoch"`
	LastUpdat	string	`json:"last_updated"`
	TempC		float32	`json:"temp_c"`
	TempF		float32	`json:"temp_f"`
	IsDay		int32	`json:"is_day"`
	Condition	struct	`json:"condition"`
}

type Condition struct {
	Text		string	`json:"text"`
	Icon		string	`json:"icon"`
	Code		int32	`json:"code"`
}

type TransformedData struct {
	City		string	`json:"city"`
	Region		string	`json:"region"`
	Country		string	`json:"coutry"`
	LocalTime	string	`json:"localtime"`
	TempF		float32	`json:"temp_f"`
	Condition	string	`json:"condition"`
	WindMph		float32	`json:"wind_mph"`
	WindDegree	int8	`json:"wind_degree"`
	WindDir		string	`json:"wind_dir"`
	PressureIn	float32	`json:"press_in"`
	PrecipIn	float32	`json:"precip_in"`
	Humidity	int8	`json:"humidity"`
	Cloud		int8	`json:"cloud"`
	FeelsLikeF	float32	`json:"feelslike_f"`
	VisMiles	int8	`json:"vis_miles"`
	Uv		int8	`json:"uv"`
	GustMph		float32	`json:"gust_mph"`
}

func Transform(input InputData) TransformedData {
	transformed := TransformedData{
		City: input.location.name
		Region: input.location.region
		Country: input.location.country
		LocalTime: input.location.localtime
		TempF: input.current.temp_f
		Condition: input.current.condition.text
		WindMph: input.wind_mph
		WindDegree: input.wind_mph
		WindDir: input.wind_dir
		PressureIn: input.pressure_in
		PrecipIn: input.precip_in
		Humidity: input.humidity
		Cloud: input.cloud
		FeelsLikeF: input.feelslike_F
		VisMiles: input.vis_miles
		Uv: input.uv
		GustMph: input.gust_mph
	}
	return transformed
}
