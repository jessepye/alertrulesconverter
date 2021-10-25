package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

var (
	hw1File string
	hw2File string
)

func init() {
	flag.StringVar(&hw1File, "hw1-file", "hw1_alerting_rules.json", "The HW1 Alerting Rules File")
	flag.StringVar(&hw2File, "hw2-file", "hw2_alerting_rules.yml", "The HW2 Alerting Rules File")
	flag.Parse()
}

func main() {
	hw1, err := LoadHW1Alerts(hw1File)
	if err != nil {
		fmt.Println(err)
	}
	hw1.Display()

	hw2, err := LoadHW2AlertRules(hw2File)
	if err != nil {
		fmt.Println(err)
	}
	hw2.Display()
}

type HW1 struct {
	ConfigFile string
	HW1Alerts  []*HW1Alert
}

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

func (hw *HW1) Display() {
	fmt.Printf("HW1 File: %s\n", hw.ConfigFile)
	for _, alert := range hw.HW1Alerts {
		fmt.Printf("Query: %s\n", alert.Query)
	}
}

func LoadHW1Alerts(filename string) (*HW1, error) {
	hw1 := &HW1{}

	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	hw1.ConfigFile = filename

	err = json.Unmarshal(jsonFile, &hw1.HW1Alerts)
	if err != nil {
		return nil, err
	}
	return hw1, nil
}

type HW2 struct {
	ConfigFile    string          `yaml:"-"`
	AlertingRules []*HW2AlertRule `yaml:"groups"`
}

type HW2AlertRule struct {
	Name  string   `yaml:"name"`
	Rules []*Alert `yaml:"rules,flow"`
}

type Alert struct {
	Alert       string      `yaml:"alert"`
	Expression  string      `yaml:"expr"`
	For         string      `yaml:"for"`
	Annotations Annotations `yaml:",flow"`
}

type Annotations struct {
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
}

func (hw *HW2) Display() {
	fmt.Printf("HW2 File: %s\n", hw.ConfigFile)
	for _, alertRule := range hw.AlertingRules {
		fmt.Printf("- Name: %s\n", alertRule.Name)
		for _, rule := range alertRule.Rules {
			fmt.Printf("  - Alert: %s\n", rule.Alert)
			fmt.Printf("    Expr: %s\n", rule.Expression)
		}
	}
}

func LoadHW2AlertRules(filename string) (*HW2, error) {
	hw2 := &HW2{}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	hw2.ConfigFile = filename

	err = yaml.Unmarshal(yamlFile, &hw2)
	if err != nil {
		return nil, err
	}
	return hw2, nil
}

func ProcessHW2Alert(alert HW1Alert) (*HW2AlertRule, error) {
	hw2alertRule := &HW2AlertRule{}
	query := alert.Query

	fmt.Printf(query)
	// "origin == 'mysql' and name == '/mysql/available' and deployment == 'cf-df12622fdcf2a31199ae' and job == 'mysql,database'"
	origin, err := regexp.Match("origin == .* and",query)
	if err != nil {
		return nil, err
	}

	fmt.Printf(origin)

	return hw2alertRule, nil
}
