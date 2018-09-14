package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {

	url := "https://api.samsara.com/v1/fleet/list?access_token=RN8Lb4AbjogbP6jQd9EO6AZ2UdPafQ"
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
