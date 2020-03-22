package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseGeocoding(t *testing.T) {
	bytes, err := ioutil.ReadFile("./fixtures/tokyo_station.xml")
	if err != nil {
		t.FailNow()
	}
	fixture := string(bytes)
	geocoding, err := ParseGeocoding(fixture)
	if err != nil {
		t.FailNow()
	}

	expected := Geocoding{
		Version: "1.2",
		Address: "東京駅",
		Coordinate: GeocodingCoordinate{
			Latitude:     "35.681236",
			Longitude:    "139.767125",
			LatitudeDMS:  "35,40,52.45",
			LongitudeDMS: "139,46,1.649",
		},
		OpenLocationCode: "8Q7XMQJ8+FR",
		URL:              "https://www.geocoding.jp/?q=%E6%9D%B1%E4%BA%AC%E9%A7%85",
		NeedsToVerify:    "no",
		GoogleMaps:       "東京駅",
	}
	if !reflect.DeepEqual(expected, *geocoding) {
		t.Errorf("expected: %+v, actual: %+v", expected, *geocoding)
	}
}
