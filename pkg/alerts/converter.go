package alerts

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`origin == '.*?' and name == '(.*?)'`)

// KPI Key Performance Indicators
type KPI struct {
	Name    string
	Metrics map[string]*Rule
}

func NewKPI(name string) *KPI {
	kpi := KPI{}
	kpi.Name = name
	return &kpi
}

type Converter struct {
	KPIs map[string]*KPI
}

func NewConverter() *Converter {
	k := NewKPI("DiegoAuctioneerMetrics")
	k.Metrics = map[string]*Rule{
		"AuctioneerLRPAuctionsFailed": {
			Alert: "AuctioneerLRPAuctionsFailed",
			Expr:  "avg(AuctioneerLRPAuctionsFailed) >= 1",
			For:   "5m",
		},
		"AuctioneerFetchStatesDuration": {
			Alert: "AuctioneerFetchStatesDuration",
			Expr:  "max_over_time(AuctioneerFetchStatesDuration[5m])/1000000000 >= 5",
			For:   "5m",
		},
		"AuctioneerLRPAuctionsStarted": {
			Alert: "AuctioneerLRPAuctionsStarted",
			Expr:  "avg(AuctioneerLRPAuctionsStarted) > 0",
			For:   "5m",
		},
		"AuctioneerTaskAuctionsFailed": {
			Alert: "AuctioneerTaskAuctionsFailed",
			Expr:  "avg_over_time(AuctioneerTaskAuctionsFailed[5m]) >= 1",
			For:   "5m",
		},
	}

	c := Converter{
		KPIs: map[string]*KPI{
			"auctioneer": k,
		},
	}
	return &c
}

func (c *Converter) ConvertHW1Alerts(hw1 *HW1Alerts) (*HW2Alerts, error) {
	alertRule := c.createHW2AlertingRules("auctioneer", *hw1)

	if alertRule == nil {
		return nil, fmt.Errorf("alerting rules found: %d", 0)
	}

	hw2 := HW2Alerts{
		Alerts: []*HW2Alert{alertRule},
	}
	return &hw2, nil
}

func (c *Converter) createHW2AlertingRules(origin string, alerts []*HW1Alert) *HW2Alert {
	if alerts == nil {
		return nil
	}

	k, ok := c.KPIs[origin]
	if !ok {
		return nil
	}

	r := make([]Rule, 0)

	for _, alert := range alerts {
		result := re.FindStringSubmatch(alert.Query)
		if len(result) > 1 {
			name := result[1]
			rule, ok := k.Metrics[name]
			if ok {
				r = append(r, *rule)
			}
		}
	}

	if len(r) < 1 {
		return nil
	}

	return &HW2Alert{
		Name:  k.Name,
		Rules: r,
	}
}
