package main

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
	}

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
		}

	}	

	WindMph		float32	`json:"wind_mph"`
	WindKph		float32	`json:"wind_kph"`
	WindDegree	int8	`json:"wind_degree"`
	WindDirection	string	`json:"wind_dir"`
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

func Transform(input *InputData) TransformedData {
	transformed := TransformedData {
		City: input.Location.Name,
		Region: input.Location.Region,
		Country: input.Location.Country,
		LocalTime: input.Location.LocalTime,
		TempF: input.Current.TempF,
		Condition: input.Current.Condition.Text,
		WindMph: input.WindMph,
		WindDegree: input.WindDegree,
		WindDir: input.WindDirection,
		PressureIn: input.PressureIn,
		PrecipIn: input.PrecipIn,
		Humidity: input.Humidity,
		Cloud: input.Cloud,
		FeelsLikeF: input.FeelsLikeF,
		VisMiles: input.VisMiles,
		Uv: input.Uv,
		GustMph: input.GustMph}
	return transformed
}
