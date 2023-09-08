package main

import (
	"encoding/json"
	"fmt"
)

type address struct {
	Street     string      `json:"street"`
	Ste        string      `json:"suite,omitempty"`
	City       string      `json:"city"`
	State      string      `json:"state"`
	Zipcode    string      `json:"zipcode"`
	Coordinate *coordinate `json:"coordinate,omitempty"`
}
type coordinate struct {
	Lat *float64 `json:"latitude,omitempty"`
	Lng *float64 `json:"longitude,omitempty"`
}

func main() {
	data := `{
		"street": "200 Larkin St",
		"city": "San Francisco",
		"state": "CA",
		"zipcode": "94102"
	}`
	cData := `{
		"latitude": 0.0,
		"longitude": 0.0
	}`
	c := new(coordinate)
	addr := new(address)
	json.Unmarshal([]byte(data), &addr)
	json.Unmarshal([]byte(cData), &c)
	addressByte, _ := json.MarshalIndent(addr, "", "    ")
	coordinateBytes, _ := json.MarshalIndent(c, "", "    ")
	fmt.Printf("%s\n", string(addressByte))
	fmt.Printf("%s\n", string(coordinateBytes))
}
