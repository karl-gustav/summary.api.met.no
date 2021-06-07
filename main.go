package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.met.no/weatherapi/locationforecast/2.0/complete?altitude=81&lat=59.41371205213798&lon=5.339655847143129", nil)
	if err != nil {
		fmt.Printf("err was %v", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36")

	r, err := client.Do(req)
	if err != nil {
		fmt.Printf("err was %v", err)
		return
	}

	var weather Weather
	err = json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		fmt.Printf("err was %v", err)
		return
	}
	defer r.Body.Close()
	for _, ts := range weather.Properties.Timeseries {
		fmt.Printf("%v\n", ts.Time)
	}
}
