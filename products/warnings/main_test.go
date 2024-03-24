package warnings_test

import (
	"encoding/json"
	"testing"

	"github.com/jordanreger/wx/products/warnings"
)

/* warnings.All */
func TestAll_Tornado(t *testing.T) {
	all := warnings.All("Tornado warning")
	res, _ := json.MarshalIndent(all, "", "	")
	t.Logf("%+v", string(res))
}

func TestAll_SevereTstorm(t *testing.T) {
	all := warnings.All("Severe thunderstorm warning")
	res, _ := json.MarshalIndent(all, "", "	")
	t.Logf("%+v", string(res))
}
func TestAll_FlashFlood(t *testing.T) {
	all := warnings.All("Flash flood warning")
	res, _ := json.MarshalIndent(all, "", "	")
	t.Logf("%+v", string(res))
}

func TestAll_All(t *testing.T) {
	TestAll_Tornado(t)
	TestAll_SevereTstorm(t)
	TestAll_FlashFlood(t)
}

/* warnings.Latest */
func TestLatest_Tornado(t *testing.T) {
	latest := warnings.Latest("Tornado warning")
	res, _ := json.MarshalIndent(latest, "", "	")
	t.Logf("%+v", string(res))
}

func TestLatest_SevereTstorm(t *testing.T) {
	latest := warnings.Latest("Severe thunderstorm warning")
	res, _ := json.MarshalIndent(latest, "", "	")
	t.Logf("%+v", string(res))
}

func TestLatest_FlashFlood(t *testing.T) {
	latest := warnings.Latest("Flash flood warning")
	res, _ := json.MarshalIndent(latest, "", "	")
	t.Logf("%+v", string(res))
}

func TestLatest_All(t *testing.T) {
	TestLatest_Tornado(t)
	TestLatest_SevereTstorm(t)
	TestLatest_FlashFlood(t)
}

/* warnings.Raw */
func TestRaw_Tornado(t *testing.T) {
	raw := warnings.Raw("Tornado warning")
	t.Log(raw)
}

func TestRaw_SevereTstorm(t *testing.T) {
	raw := warnings.Raw("Severe thunderstorm warning")
	t.Log(raw)
}

func TestRaw_FlashFlood(t *testing.T) {
	raw := warnings.Raw("Flash flood warning")
	t.Log(raw)
}

func TestRaw_All(t *testing.T) {
	TestRaw_Tornado(t)
	TestRaw_SevereTstorm(t)
	TestRaw_FlashFlood(t)
}

/* warnings.Types */
func TestTypes(t *testing.T) {
	types := warnings.Types()
	t.Log(types)
}
