package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = config.GetString("data.dsn")

func CreateTable() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		query := `
							CREATE TABLE weather (
								id INT AUTO_INCREMENT,
								city VARCHAR(255),
								region VARCHAR(255),
								country VARCHAR(255),
								currentTime VARCHAR(255),
								tempF FLOAT(8,4),
								weatherConditions VARCHAR(255),
								windMph FLOAT(8,4),
								windDegree FLOAT(8,4),
								windDir VARCHAR(255),
								pressureIn FLOAT(8,4),
								precipIn FLOAT(8,4),
								humidity FLOAT(8,4),
								cloud FLOAT(8,4),
								feelsLikeF FLOAT(8,4),
								visMiles FLOAT(8,4),
								uv FLOAT(8,4),
								gustMph FLOAT(8,4),
								sunrise VARCHAR(255),
								sunset VARCHAR(255),
								moonrise VARCHAR(255),
								moonset VARCHAR(255),
								moonPhase VARCHAR(255),
								moonIllum TINYINT,
								isMoonUp TINYINT,
								isSunUp TINYINT,
								PRIMARY KEY (id)
						);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}

func Insert(input TransformedData) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	city := input.City
	region := input.Region
	country := input.Country
	currentTime := input.LocalTime
	tempF := input.TempF
	weatherConditions := input.Condition
	windMph := input.WindMph
	windDegree := input.WindDegree
	windDir := input.WindDir
	pressureIn := input.PressureIn
	precipIn := input.PrecipIn
	humidity := input.Humidity
	cloud := input.Cloud
	feelslikeF := input.FeelsLikeF
	visMiles := input.VisMiles
	uv := input.Uv
	gustMph := input.GustMph
	sunrise := input.Sunrise
	sunset := input.Sunset
	moonrise := input.Moonrise
	moonset := input.Moonset
	moonPhase := input.MoonPhase
	moonIllum := input.MoonIllum
	isMoonUp := input.IsMoonUp
	isSunUp := input.IsSunUp

	result, err := db.Exec(`INSERT INTO weather (city, region, country, currentTime, tempF, weatherConditions, windMph, 
		windDegree, windDir, pressureIn, precipIn, humidity, cloud, feelsLikeF, visMiles, uv, gustMph, sunrise, sunset,
		moonrise, moonset, moonPhase, moonIllum, isMoonUp, isSunUp) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		city, region, country, currentTime, tempF, weatherConditions, windMph, windDegree, windDir, pressureIn, precipIn, humidity, cloud, feelslikeF,
		visMiles, uv, gustMph, sunrise, sunset, moonrise, moonset, moonPhase, moonIllum, isMoonUp, isSunUp)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}

func QueryAll() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	type weather struct {
		id    int
		city  string
		tempF float32
	}

	rows, err := db.Query(`SELECT id, city, tempF FROM weather`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var weatherResults []weather
	for rows.Next() {
		var result weather

		err := rows.Scan(&result.id, &result.city, &result.tempF)
		if err != nil {
			log.Fatal(err)
		}
		weatherResults = append(weatherResults, result)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", weatherResults)
}
