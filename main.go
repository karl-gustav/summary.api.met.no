package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/karl-gustav/api.met.coverter/met"
)

const (
	ue = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"
)

type MySeries struct {
	TimeSeries                  []time.Time `json:"time_series"`
	AirTemperatureMax           float32     `json:"air_temperature_max"`
	AirTemperatureMin           float32     `json:"air_temperature_min"`
	UltravioletIndexClearSkyMax float32     `json:"ultraviolet_index_clear_sky"`
	WindSpeedMax                float32     `json:"wind_speed_max"`
	WindSpeedOfGustMax          float32     `json:"wind_speed_of_gust_max"`
	ProbabilityOfPrecipitation  float32     `json:"probability_of_precipitation"`
	CloudAreaFractionAvreage    float32     `json:"cloud_area_fraction_avreage"`
}

var loc *time.Location

func init() {
	var err error

	loc, err = time.LoadLocation("Europe/Oslo")
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		client := &http.Client{}
		now := time.Now()
		from, err := queryToRFC3339(req.URL.Query(), "from", now)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		to, err := queryToRFC3339(req.URL.Query(), "to", endOfDay(now))
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		metRequest, err := http.NewRequest("GET", "https://api.met.no/weatherapi/locationforecast/2.0/complete?altitude=81&lat=59.41371205213798&lon=5.339655847143129", nil)
		if err != nil {
			fmt.Printf("err was %v", err)
			return
		}

		metRequest.Header.Set("User-Agent", ue)

		r, err := client.Do(metRequest)
		if err != nil {
			http.Error(res, "failed to get data from api.met.no: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if r.StatusCode != http.StatusOK {
			http.Error(
				res,
				fmt.Sprintf("failed to get data from api.met.no, status code was %d", r.StatusCode),
				http.StatusInternalServerError,
			)
			return
		}

		var forecast met.Forecast
		err = json.NewDecoder(r.Body).Decode(&forecast)
		if err != nil {
			http.Error(res, "failed to decode json into met.Forecast struct: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		timeSeries := filter(forecast.Properties.Timeseries, *from, *to)
		if len(timeSeries) == 0 {
			http.Error(res, "no time series found for today", http.StatusInternalServerError)
			return
		}
		oneDay := generateSeries(timeSeries)

		res.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(res).Encode(oneDay)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Serving http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func queryToRFC3339(query url.Values, queryKey string, defaultTime time.Time) (*time.Time, error) {
	queryString := query.Get(queryKey)
	if queryString == "" {
		return &defaultTime, nil
	}
	queryTime, err := time.Parse(time.RFC3339, queryString)
	if err != nil {
		return nil, fmt.Errorf("could not parse %s as RFC3339: %w", queryString, err)
	}
	return &queryTime, nil
}

func filter(timeSeries []met.TimeSerie, from, to time.Time) (out []met.TimeSerie) {
	for _, serie := range timeSeries {
		if serie.Time.After(to) {
			break
		}
		if serie.Time.Before(from) {
			continue
		}
		out = append(out, serie)
	}
	return out
}

func generateSeries(timeSeries []met.TimeSerie) MySeries {
	oneDay := MySeries{
		AirTemperatureMin: math.MaxFloat32,
		AirTemperatureMax: -math.MaxFloat32,
	}
	var cloudAreaFraction float32
	for _, series := range timeSeries {
		details := series.Data.Instant.Details
		oneDay.TimeSeries = append(oneDay.TimeSeries, series.Time.In(loc))
		if details.AirTemperature > oneDay.AirTemperatureMax {
			oneDay.AirTemperatureMax = details.AirTemperature
		}
		if details.AirTemperature < oneDay.AirTemperatureMin {
			oneDay.AirTemperatureMin = details.AirTemperature
		}
		if details.UltravioletIndexClearSky != nil &&
			*details.UltravioletIndexClearSky > oneDay.UltravioletIndexClearSkyMax {
			oneDay.UltravioletIndexClearSkyMax = *details.UltravioletIndexClearSky
		}
		if details.WindSpeed > oneDay.WindSpeedMax {
			oneDay.WindSpeedMax = details.WindSpeed
		}
		if details.WindSpeedOfGust != nil && *details.WindSpeedOfGust > oneDay.WindSpeedOfGustMax {
			oneDay.WindSpeedOfGustMax = *details.WindSpeedOfGust
		}
		if next1Hour := series.Data.Next1Hour; next1Hour != nil {
			if next1Hour.Details.ProbabilityOfPrecipitation > oneDay.ProbabilityOfPrecipitation {
				oneDay.ProbabilityOfPrecipitation = next1Hour.Details.ProbabilityOfPrecipitation
			}
		}
		cloudAreaFraction += details.CloudAreaFraction
	}
	oneDay.CloudAreaFractionAvreage = round(cloudAreaFraction/float32(len(timeSeries)), 2)
	return oneDay
}

func round(number float32, decimalPlaces int) float32 {
	d := math.Pow10(decimalPlaces)
	return float32(math.Round(float64(number)*d) / d)
}

func endOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999, loc)
}
