package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Vehicle - testing comment
type Vehicle struct {
	Vehicles []Fleetlist
}

// Fleetlist - testing comment
type Fleetlist struct {
	Name           string
	Id             int
	Vin            string
	OdometerMeters int
	EngineHours    int
}

// Configuration - Importing access token from config.json
type Configuration struct {
	Accesstoken string
}

func main() {

	// Read in the access token from an untracked local file
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("File failed: ", err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var config Configuration
	json.Unmarshal(byteValue, &config)

	// Generate the API query
	url := "https://api.samsara.com/v1/fleet/list?access_token=" + config.Accesstoken
	postdata := []byte(`{"groupId": 3991}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(postdata))

	client := &http.Client{}

	if err != nil {
		fmt.Println("improper request", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Response invalid: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var vehicle Vehicle
	json.Unmarshal(body, &vehicle)
	for _, v := range vehicle.Vehicles {
		fmt.Printf("id: %d  Name: %s  vin: %s  odometer: %d  enginehrs: %d \n", v.Id, v.Name, v.Vin, v.OdometerMeters, v.EngineHours)
	}
	fmt.Println(vehicle.Vehicles[4].Id)
}
