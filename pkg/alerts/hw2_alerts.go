package alerts

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type HW2Alerts struct {
	Alerts []*HW2Alert `yaml:"groups"`
}

type HW2Alert struct {
	Name  string `yaml:"name"`
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Alert       string `yaml:"alert"`
	Expr        string `yaml:"expr"`
	For         string `yaml:"for"`
	Annotations struct {
		Summary     string `yaml:"summary"`
		Description string `yaml:"description"`
	}
}

func WriteHW2File(filename string, hw2 *HW2Alerts) error {
	if hw2.Alerts == nil {
		return fmt.Errorf("alerts is length %d", 0)
	}

	data, err := yaml.Marshal(hw2)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
