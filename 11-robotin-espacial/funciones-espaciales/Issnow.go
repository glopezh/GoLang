package funciones_espaciales

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Traer información del API iss-now

type IssnowResponse struct {
	Message     string `json:"message"`
	Timestamp   int    `json:"timestamp"`
	IssPosition struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	} `json:"iss_position"`
}

func Issnow() [2]string {

	// llamando al otro API

	response, err := http.Get("http://api.open-notify.org/iss-now.json")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	IssnowResponse := &IssnowResponse{}
	json.Unmarshal(body, IssnowResponse)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latitude: %s\n", IssnowResponse.IssPosition.Latitude)
	fmt.Printf("Longitude: %s\n", IssnowResponse.IssPosition.Longitude)
	result  := [2]string{IssnowResponse.IssPosition.Latitude,IssnowResponse.IssPosition.Longitude}
	return result
}

