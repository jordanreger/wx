package social

import "strings"

func GetHashtags(area string) string {
	var states []string

	areas := strings.Split(area, "; ")
	for _, area := range areas {
		states = append(states, "#"+strings.Split(area, ", ")[1]+"WX")
	}

	hashtags := strings.Join(states, " ")
	return hashtags
}
