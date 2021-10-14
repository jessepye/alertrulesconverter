package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type HW1Alerts struct {
    HW1Alerts []HW1Alert `map[string]interface{} `
}

type HW1Alert struct {
    Query string `json:"query"`
    Enabled bool `json:"enabled"`
    Threshold Threshold `json:"threshold"`
}

type Threshold struct {
    Critical float32 `json:"critical"`
    Warning float32 `json:"warning"`
    Type string `json:"type"`
}

func main() {
	jsonFile, err := ioutil.ReadFile("hw1_alerting_rules.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened json")

	var hw1Alerts []HW1Alert

    err = json.Unmarshal([]byte(jsonFile), &hw1Alerts)

	if err != nil {
	    fmt.Println(err)
    }

    fmt.Printf("%+v\n", hw1Alerts)        // [{AutoCAD Autodesk} {Firefox Mozilla} {Chrome Google}]
}
