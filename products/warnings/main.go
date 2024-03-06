// https://fjall.dev/products/warnings/main.go

package warnings

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type body struct {
	Type     string    `json:"type"`
	Features []Warning `json:"features"`
	Title    string    `json:"title"`
	Updated  time.Time `json:"updated"`
}

type Warning struct {
	ID         string       `json:"id"`
	Properties w_properties `json:"properties"`
}

type w_properties struct {
	Area        string     `json:"areaDesc"`
	Effective   time.Time  `json:"effective"`
	Onset       time.Time  `json:"onset"`
	Expires     time.Time  `json:"expires"`
	Ends        time.Time  `json:"ends"`
	MessageType string     `json:"messageType"`
	Severity    string     `json:"severity"`
	Certainty   string     `json:"certainty"`
	Urgency     string     `json:"urgency"`
	Event       string     `json:"event"`
	Sender      string     `json:"senderName"`
	Headline    string     `json:"headline"`
	Description string     `json:"description"`
	Parameters  parameters `json:"parameters"`
}

type parameters struct {
	AWIPSidentifier []string `json:"AWIPSidentifier"`
	WMOidentifier   []string `json:"WMOidentifier"`
	NWSheadline     []string `json:"NWSheadline"`
}

type tbody struct {
	EventTypes []string `json:"eventTypes"`
}

func All(wt string) []string {
	res, err := http.Get("https://api.weather.gov/alerts/active?status=actual&event=" + wt + "&limit=1")
	if err != nil {
		panic(err)
	}

	var wb body
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &wb)

	features := wb.Features

	var warnings []string

	if len(features) == 0 {
		wtxt := "No current " + strings.ToLower(strings.ReplaceAll(wt, "%20", " ")) + "s"
		warnings = append(warnings, wtxt)
	}

	for _, warning := range features {
		area := warning.Properties.Area
		station := warning.Properties.Parameters.AWIPSidentifier[0][3:6]
		headline := warning.Properties.Headline

		wtxt := station + " - " + area + "\n\n" + headline + "\n"
		warnings = append(warnings, wtxt)
	}

	return warnings
}

func Latest(wt string) string {
	res, err := http.Get("https://api.weather.gov/alerts/active?status=actual&event=" + wt + "&limit=1")
	if err != nil {
		panic(err)
	}

	var wb body
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &wb)

	features := wb.Features

	var warning string

	if len(features) == 0 {
		warning = "No current " + strings.ToLower(strings.ReplaceAll(wt, "%20", " ")) + "s"
		return warning
	}

	latest := features[0]
	messageType := latest.Properties.MessageType
	event := latest.Properties.Event
	area := latest.Properties.Area
	state := strings.Split(area, ", ")
	// sender := latest.Properties.Sender
	// headline := latest.Properties.Parameters.NWSheadline[0]
	ends := latest.Properties.Ends.UTC().Format(time.RFC1123)

	if messageType == "Alert" {
		warning = event + " including " + area + " until " + ends + "\n\n#" + state[len(state)-1] + "WX"
	} else if messageType == "Update" {
		warning = event + " continues for " + area + " until " + ends + "\n\n#" + state[len(state)-1] + "WX"
	} else {
		warning = "No current " + strings.ToLower(strings.ReplaceAll(wt, "%20", " ")) + "s"
	}

	return warning
}

func Raw(wt string) string {
	res, err := http.Get("https://api.weather.gov/alerts/active?status=actual&event=" + wt + "&limit=1")
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(res.Body)

	return string(b)
}

func Types() string {
	res, err := http.Get("https://api.weather.gov/alerts/types")
	if err != nil {
		panic(err)
	}

	var tb tbody
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &tb)

	var types string
	ets := tb.EventTypes
	for _, t := range ets {
		types += t + "\n"
	}

	return types
}
