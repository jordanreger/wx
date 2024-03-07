package warnings_test

import (
	"testing"

	"github.com/fjalldev/wx/products/warnings"
)

func TestLatest_Tornado(t *testing.T) {
	latest := warnings.Latest("Tornado warning")
	t.Log(latest)
}

func TestLatest_SevereTstorm(t *testing.T) {
	latest := warnings.Latest("Severe thunderstorm warning")
	t.Log(latest)
}

func TestLatest_FlashFlood(t *testing.T) {
	latest := warnings.Latest("Flash flood warning")
	t.Log(latest)
}

func TestLatest_All(t *testing.T) {
	TestLatest_Tornado(t)
	TestLatest_SevereTstorm(t)
	TestLatest_FlashFlood(t)
}
