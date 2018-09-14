package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var modelyear = "2010"
var maker = "ford"
var model = "fusion"

func main() {
	fmt.Println("Starting the search...")
	var url bytes.Buffer
	url.WriteString("https://one.nhtsa.gov/webapi/api/Recalls/vehicle/modelyear/")
	url.WriteString(modelyear)
	url.WriteString("/make/")
	url.WriteString(maker)
	url.WriteString("/model/")
	url.WriteString(model)
	url.WriteString("?format=json")
	fmt.Println(url.String())
	response, err := http.Get(url.String())
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("Terminating the application...")
}
