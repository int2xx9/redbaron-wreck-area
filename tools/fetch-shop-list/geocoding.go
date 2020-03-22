package main

import (
	"encoding/xml"
)

type GeocodingCoordinate struct {
	Latitude     string `xml:"lat"`
	Longitude    string `xml:"lng"`
	LatitudeDMS  string `xml:"lat_dms"`
	LongitudeDMS string `xml:"lng_dms"`
}

type Geocoding struct {
	Version          string              `xml:"version"`
	Address          string              `xml:"address"`
	Coordinate       GeocodingCoordinate `xml:"coordinate"`
	OpenLocationCode string              `xml:"open_location_code"`
	URL              string              `xml:"url"`
	NeedsToVerify    string              `xml:"needs_to_verify"`
	GoogleMaps       string              `xml:"google_maps"`
}

func ParseGeocoding(xmlData string) (*Geocoding, error) {
	parsed := Geocoding{}
	if err := xml.Unmarshal([]byte(xmlData), &parsed); err != nil {
		return nil, err
	}
	return &parsed, nil
}
