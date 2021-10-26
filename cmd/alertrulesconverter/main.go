package main

import (
	"alertrulesconverter/pkg/alerts"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	hw1File string
	hw2File string
)

func init() {
	flag.StringVar(&hw1File, "hw1-file", "hw1.json", "The HW1 Alerting Rules File")
	flag.StringVar(&hw2File, "hw2-file", "hw2.yml", "The HW2 Alerting Rules File")
	flag.Parse()
}

func main() {
	hw1, err := alerts.LoadHW1File(hw1File)
	if err != nil {
		panic(err)
	}

	conv := alerts.NewConverter()
	hw2, err := conv.ConvertHW1Alerts(hw1)
	if err != nil {
		panic(err)
	}

	err = alerts.WriteHW2File(hw2File, hw2)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(hw2File)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
