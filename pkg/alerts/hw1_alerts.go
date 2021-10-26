package alerts

import (
	"encoding/json"
	"io/ioutil"
)

type HW1Alerts []*HW1Alert

type HW1Alert struct {
	Query     string    `json:"query"`
	Enabled   bool      `json:"enabled"`
	Threshold Threshold `json:"threshold"`
}

type Threshold struct {
	Type     string  `json:"type"`
	Critical float32 `json:"critical"`
	Warning  float32 `json:"warning,omitempty"`
}

func LoadHW1File(filename string) (*HW1Alerts, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	hw1 := HW1Alerts{}

	err = json.Unmarshal(content, &hw1)
	if err != nil {
		return nil, err
	}

	return &hw1, nil
}
