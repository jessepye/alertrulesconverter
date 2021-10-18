package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type HW1Alert struct {
	Query     string    `json:"query"`
	Enabled   bool      `json:"enabled"`
	Threshold Threshold `json:"threshold"`
}

type Threshold struct {
	Critical float32 `json:"critical"`
	Type     string  `json:"type"`
	Warning  float32 `json:"warning,omitempty"`
}

type HW2AlertRules struct {
	AlertingRules []HW2AlertRule `yaml:"groups"`
}

type HW2AlertRule struct {
	Name  string  `yaml:"name"`
	Rules []Alert `yaml:"rules,flow"`
}

type Alert struct {
	Alert       string `yaml:"alert"`
	Expression  string `yaml:"expr"`
	For         string `yaml:"for"`
	Annotations struct {
		Summary     string
		Description string
	}
}

func main() {
	jsonFile, err := ioutil.ReadFile("hw1_alerting_rules.json")
	if err != nil {
		fmt.Println(err)
	}

	var hw1Alerts []HW1Alert
	err = json.Unmarshal(jsonFile, &hw1Alerts)
	if err != nil {
		fmt.Println(err)
	}

	yamlFile, err := ioutil.ReadFile("hw2_alerting_rules.yml")
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range hw1Alerts {
		s := v.Query
		fmt.Printf("Query: %s\n", s)
	}

	var hw2 HW2AlertRules
	err = yaml.Unmarshal([]byte(yamlFile), &hw2)
	if err != nil {
		fmt.Println(err)
	}

	var hw2Map = make(map[string]map[string]string)
	for _, alertRule := range hw2.AlertingRules {
		hw2Map[alertRule.Name] = map[string]string{}
		for _, rule := range alertRule.Rules {
			hw2Map[alertRule.Name][rule.Alert] = rule.Expression
		}
	}
	for k, rules := range hw2Map {
		fmt.Printf("- %s\n", k)
		for alert, expr := range rules {
			fmt.Printf("  Alert: %s\n", alert)
			fmt.Printf("  Expr: %s\n", expr)
		}
	}
}
