// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.12.0

package querytest

import (
	"database/sql"
	"time"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type WeatherMetric struct {
	Time            time.Time
	TimezoneShift   sql.NullInt32
	CityName        sql.NullString
	TempC           interface{}
	FeelsLikeC      interface{}
	TempMinC        interface{}
	TempMaxC        interface{}
	PressureHpa     interface{}
	HumidityPercent interface{}
	WindSpeedMs     interface{}
	WindDeg         sql.NullInt32
	Rain1hMm        interface{}
	Rain3hMm        interface{}
	Snow1hMm        interface{}
	Snow3hMm        interface{}
	CloudsPercent   sql.NullInt32
	WeatherTypeID   sql.NullInt32
}
