package met

import "time"

type Forecast struct {
	Type     string `json:"type"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float32 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Meta       Meta        `json:"meta"`
		Timeseries []TimeSerie `json:"timeseries"`
	} `json:"properties"`
}

type TimeSerie struct {
	Time time.Time `json:"time"`
	Data struct {
		Instant     InstantDetails `json:"instant"`
		Next1Hour   *Next1Hour     `json:"next_1_hours"`
		Next6Hours  *Next6Hours    `json:"next_6_hours"`
		Next12Hours *Next12Hours   `json:"next_12_hours"`
	} `json:"data,omitempty"`
}

type InstantDetails struct {
	Details struct {
		AirPressureAtSeaLevel      float32  `json:"air_pressure_at_sea_level"`
		AirTemperature             float32  `json:"air_temperature"`
		AirTemperaturePercentile10 float32  `json:"air_temperature_percentile_10"`
		AirTemperaturePercentile90 float32  `json:"air_temperature_percentile_90"`
		CloudAreaFraction          float32  `json:"cloud_area_fraction"`
		CloudAreaFractionHigh      float32  `json:"cloud_area_fraction_high"`
		CloudAreaFractionLow       float32  `json:"cloud_area_fraction_low"`
		CloudAreaFractionMedium    float32  `json:"cloud_area_fraction_medium"`
		DewPointTemperature        float32  `json:"dew_point_temperature"`
		FogAreaFraction            *float32 `json:"fog_area_fraction"`
		RelativeHumidity           float32  `json:"relative_humidity"`
		UltravioletIndexClearSky   *float32 `json:"ultraviolet_index_clear_sky"`
		WindFromDirection          float32  `json:"wind_from_direction"`
		WindSpeed                  float32  `json:"wind_speed"`
		WindSpeedOfGust            *float32 `json:"wind_speed_of_gust"`
		WindSpeedPercentile10      float32  `json:"wind_speed_percentile_10"`
		WindSpeedPercentile90      float32  `json:"wind_speed_percentile_90"`
	} `json:"details"`
}

type Next1Hour struct {
	Summary struct {
		SymbolCode string `json:"symbol_code"`
	} `json:"summary"`
	Details struct {
		PrecipitationAmount        float32 `json:"precipitation_amount"`
		PrecipitationAmountMax     float32 `json:"precipitation_amount_max"`
		PrecipitationAmountMin     float32 `json:"precipitation_amount_min"`
		ProbabilityOfPrecipitation float32 `json:"probability_of_precipitation"`
		ProbabilityOfThunder       float32 `json:"probability_of_thunder"`
	} `json:"details"`
}

type Next6Hours struct {
	Summary struct {
		SymbolCode string `json:"symbol_code"`
	} `json:"summary"`
	Details struct {
		AirTemperatureMax          float32 `json:"air_temperature_max"`
		AirTemperatureMin          float32 `json:"air_temperature_min"`
		PrecipitationAmount        float32 `json:"precipitation_amount"`
		PrecipitationAmountMax     float32 `json:"precipitation_amount_max"`
		PrecipitationAmountMin     float32 `json:"precipitation_amount_min"`
		ProbabilityOfPrecipitation float32 `json:"probability_of_precipitation"`
	} `json:"details"`
}

type Next12Hours struct {
	Summary struct {
		SymbolCode       string `json:"symbol_code"`
		SymbolConfidence string `json:"symbol_confidence"`
	} `json:"summary"`
	Details struct {
		ProbabilityOfPrecipitation float32 `json:"probability_of_precipitation"`
	} `json:"details"`
}

type Meta struct {
	UpdatedAt time.Time `json:"updated_at"`
	Units     struct {
		AirPressureAtSeaLevel      string `json:"air_pressure_at_sea_level"`
		AirTemperature             string `json:"air_temperature"`
		AirTemperatureMax          string `json:"air_temperature_max"`
		AirTemperatureMin          string `json:"air_temperature_min"`
		AirTemperaturePercentile10 string `json:"air_temperature_percentile_10"`
		AirTemperaturePercentile90 string `json:"air_temperature_percentile_90"`
		CloudAreaFraction          string `json:"cloud_area_fraction"`
		CloudAreaFractionHigh      string `json:"cloud_area_fraction_high"`
		CloudAreaFractionLow       string `json:"cloud_area_fraction_low"`
		CloudAreaFractionMedium    string `json:"cloud_area_fraction_medium"`
		DewPointTemperature        string `json:"dew_point_temperature"`
		FogAreaFraction            string `json:"fog_area_fraction"`
		PrecipitationAmount        string `json:"precipitation_amount"`
		PrecipitationAmountMax     string `json:"precipitation_amount_max"`
		PrecipitationAmountMin     string `json:"precipitation_amount_min"`
		ProbabilityOfPrecipitation string `json:"probability_of_precipitation"`
		ProbabilityOfThunder       string `json:"probability_of_thunder"`
		RelativeHumidity           string `json:"relative_humidity"`
		UltravioletIndexClearSky   string `json:"ultraviolet_index_clear_sky"`
		WindFromDirection          string `json:"wind_from_direction"`
		WindSpeed                  string `json:"wind_speed"`
		WindSpeedOfGust            string `json:"wind_speed_of_gust"`
		WindSpeedPercentile10      string `json:"wind_speed_percentile_10"`
		WindSpeedPercentile90      string `json:"wind_speed_percentile_90"`
	} `json:"units"`
}
