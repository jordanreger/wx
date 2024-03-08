// https://fjall.dev/products/warnings/main.go

package warnings

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Body struct {
	Type     string    `json:"type"`
	Warnings []Warning `json:"features"`
	Title    string    `json:"title"`
	Updated  time.Time `json:"updated"`
}

type Warning struct {
	ID         string      `json:"id,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
}

type Properties struct {
	Area        string      `json:"areaDesc,omitempty"`
	Effective   *time.Time  `json:"effective,omitempty"`
	Onset       *time.Time  `json:"onset,omitempty"`
	Expires     *time.Time  `json:"expires,omitempty"`
	Ends        *time.Time  `json:"ends,omitempty"`
	MessageType string      `json:"messageType,omitempty"`
	Severity    string      `json:"severity,omitempty"`
	Certainty   string      `json:"certainty,omitempty"`
	Urgency     string      `json:"urgency,omitempty"`
	Event       string      `json:"event,omitempty"`
	Sender      string      `json:"senderName,omitempty"`
	Headline    string      `json:"headline,omitempty"`
	Description string      `json:"description,omitempty"`
	Parameters  *Parameters `json:"parameters,omitempty"`
}

type Parameters struct {
	AWIPSidentifier []string `json:"AWIPSidentifier"`
	WMOidentifier   []string `json:"WMOidentifier"`
	NWSheadline     []string `json:"NWSheadline"`
}

type tbody struct {
	EventTypes []string `json:"eventTypes"`
}

func All(wt string) []Warning {
	res, err := http.Get("https://api.weather.gov/alerts/active?status=actual&event=" + wt + "&limit=1")
	if err != nil {
		panic(err)
	}

	var body Body
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &body)

	var all []Warning

	if len(body.Warnings) == 0 {
		all = append(all, Warning{
			Properties: &Properties{
				Headline: "No current " + strings.ToLower(strings.ReplaceAll(wt, "%20", " ")) + "s",
			},
		})
	} else {
		all = body.Warnings
	}

	return all
}

func Latest(wt string) Warning {
	res, err := http.Get("https://api.weather.gov/alerts/active?status=actual&event=" + wt + "&limit=1")
	if err != nil {
		panic(err)
	}

	var body Body
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &body)

	var latest Warning

	if len(body.Warnings) == 0 {
		latest = Warning{
			Properties: &Properties{
				Headline: "No current " + strings.ToLower(strings.ReplaceAll(wt, "%20", " ")) + "s",
			},
		}
	} else {
		latest = body.Warnings[0]
	}

	return latest
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
