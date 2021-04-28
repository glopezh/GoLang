package funciones_espaciales

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Trear información del API whereisiss
type WhereisissResponse struct {
	Name       string  `json:"name"`
	ID         int     `json:"id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Altitude   float64 `json:"altitude"`
	Velocity   float64 `json:"velocity"`
	Visibility string  `json:"visibility"`
	Footprint  float64 `json:"footprint"`
	Timestamp  int     `json:"timestamp"`
	Daynum     float64 `json:"daynum"`
	SolarLat   float64 `json:"solar_lat"`
	SolarLon   float64 `json:"solar_lon"`
	Units      string  `json:"units"`
}

func Whereisiss() {
	response, err := http.Get("https://api.wheretheiss.at/v1/satellites/25544")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	WhereisissResponse := &WhereisissResponse{}
	json.Unmarshal(body, WhereisissResponse)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latitude: %s\n", WhereisissResponse.Latitude)
	fmt.Printf("Longitude: %s\n", WhereisissResponse.Longitude)
}
