package main

import (
	"encoding/json"
	"log"
)

type InputData struct {
	Location	struct	{
	Name		string	`json:"name"`
	Region		string	`json:"region"`
	Country		string	`json:"country"`
	Latitude	float32	`json:"lat"`
	Longitude	float32	`json:"lon"`
	TimeZone	string	`json:"tz_id"`
	LocalTimeEpoch	int32	`json:"localtime_epoch"`
	LocalTime	string	`json:"localtime"`
	} `json:"location"`

	Current		struct	{
	LastUpdateEpoch	int32	`json:"last_updated_epoch"`
	LastUpdat	string	`json:"last_updated"`
	TempC		float32	`json:"temp_c"`
	TempF		float32	`json:"temp_f"`
	IsDay		int32	`json:"is_day"`
	Condition	struct	{
		Text		string	`json:"text"`
		Icon		string	`json:"icon"`
		Code		int32	`json:"code"`
		} `json:"condition"`	
	WindMph		float32	`json:"wind_mph"`
	WindKph		float32	`json:"wind_kph"`
	WindDegree	float32	`json:"wind_degree"`
	WindDirection	string	`json:"wind_dir"`
	PressureMb	float32	`json:"pressure_mb"`
	PressureIn	float32	`json:"pressure_in"`
	PrecipMm	float32	`json:"precip_mm"`
	PrecipIn	float32	`json:"precip_in"`
	Humidity	float32	`json:"humidity"`
	Cloud		float32	`json:"cloud"`
	FeelsLikeC	float32	`json:"feelslike_C"`
	FeelsLikeF	float32	`json:"feelslike_F"`
	VisKm		float32	`json:"vis_km"`
	VisMiles	float32	`json:"vis_miles"`
	Uv		float32	`json:"uv"`
	GustMph		float32	`json:"gust_mph"`
	GustKph		float32	`json:"gust_kph"`
	} `json:"current"`

}

type AstroData struct {
	Location struct {
	Name		string	`json:"name"`
	Region		string	`json:"region"`
	Country		string	`json:"country"`
	Latitude	float32	`json:"lat"`
	Longitude	float32	`json:"lon"`
	TimeZone	string	`json:"tz_id"`
	LocalTimeEpoch	int32	`json:"localtime_epoch"`
	LocalTime	string	`json:"localtime"`
	} `json:"location"`
	Astronomy struct {
	    Astro struct {
		Sunrise		string	`json:"sunrise"`
		Sunset		string	`json:"sunset"`
		Moonrise	string	`json:"moonrise"`
		Moonset		string	`json:"moonset"`
		MoonPhase	string	`json:"moon_phase"`
		MoonIllum	string	`json:"moon_illumination"`
		IsMoonUp	int8	`json:"is_moon_up"`
		IsSunUp		int8	`json:"is_sun_up"`
	    } `json:"astro"`
	} `json:"astronomy"`
}

type TransformedData struct {
	City		string	`json:"City"`
	Region		string	`json:"Region"`
	Country		string	`json:"Country"`
	LocalTime	string	`json:"Local Time"`
	TempF		float32	`json:"Temp F"`
	Condition	string	`json:"Condition"`
	WindMph		float32	`json:"Wind MPH"`
	WindDegree	float32	`json:"Wind Degree"`
	WindDir		string	`json:"Wind Direction"`
	PressureIn	float32	`json:"Pressure(inches)"`
	PrecipIn	float32	`json:"Precipitaion(inches)"`
	Humidity	float32	`json:"Humidity"`
	Cloud		float32	`json:"Cloud"`
	FeelsLikeF	float32	`json:"Feels Like F"`
	VisMiles	float32	`json:"Visibility(Miles)"`
	Uv		float32	`json:"UV"`
	GustMph		float32	`json:"Gust(MPH)"`
	Sunrise		string	`json:"Sunrise"`
	Sunset		string	`json:"Sunset"`
	Moonrise	string	`json:"Moonrise"`
	Moonset		string	`json:"Moonset"`
	MoonPhase	string	`json:"Moon Phase"`
	MoonIllum	string	`json:"Moon Illumination"`
	IsMoonUp	int8	`json:"Is Moon Up"`
	IsSunUp		int8	`json:"Is Sun Up"`


}

func Transform(input *InputData, astroData *AstroData) string {
	transformed := TransformedData {
		City: input.Location.Name,
		Region: input.Location.Region,
		Country: input.Location.Country,
		LocalTime: input.Location.LocalTime,
		TempF: input.Current.TempF,
		Condition: input.Current.Condition.Text,
		WindMph: input.Current.WindMph,
		WindDegree: input.Current.WindDegree,
		WindDir: input.Current.WindDirection,
		PressureIn: input.Current.PressureIn,
		PrecipIn: input.Current.PrecipIn,
		Humidity: input.Current.Humidity,
		Cloud: input.Current.Cloud,
		FeelsLikeF: input.Current.FeelsLikeF,
		VisMiles: input.Current.VisMiles,
		Uv: input.Current.Uv,
		GustMph: input.Current.GustMph,
		Sunrise: astroData.Astronomy.Astro.Sunrise,
		Sunset: astroData.Astronomy.Astro.Sunset,
		Moonrise: astroData.Astronomy.Astro.Moonrise,
		Moonset: astroData.Astronomy.Astro.Moonset,
		MoonPhase: astroData.Astronomy.Astro.MoonPhase,
		MoonIllum: astroData.Astronomy.Astro.MoonIllum,
		IsMoonUp: astroData.Astronomy.Astro.IsMoonUp,
		IsSunUp: astroData.Astronomy.Astro.IsSunUp}
	
	formattedBody, err := json.MarshalIndent(transformed, "", "    ")

	if err != nil {
		log.Fatal(err)
	}
	return string(formattedBody)
}
