package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"golang.org/x/xerrors"
)

const (
	areaListUrl = "https://www01.redbaron.co.jp/shop/search/"
)

var (
	ErrHttp = xerrors.New("status code isn't 200")
)

func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", ErrHttp
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func main() {
	areaListRaw, err := Fetch(areaListUrl)
	if err != nil {
		panic(err)
	}

	areaList, err := ParseAreaListPage(areaListRaw)
	if err != nil {
		panic(err)
	}

	shops := []ShopData{}
	for _, area := range areaList.Areas {
		areaPageHtml, err := Fetch(area.URL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: Fetch: %s\n", err)
			continue
		}

		areaPage, err := ParseAreaPage(areaPageHtml)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: ParseAreaPage: %s\n", err)
			continue
		}

		for _, shop := range areaPage.Shops {
			fmt.Fprintf(os.Stderr, "shop: %s / %s\n", shop.Name, shop.Address)

			geoHtml, err := Fetch("https://www.geocoding.jp/api/?q=" + shop.Address)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: Fetch: %s\n", err)
				continue
			}

			geo, err := ParseGeocoding(geoHtml)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: ParseGeocoding: %s\n", err)
				continue
			}

			shops = append(shops, ShopData{
				Name:      shop.Name,
				Address:   shop.Address,
				Longitude: geo.Coordinate.Longitude,
				Latitude:  geo.Coordinate.Latitude,
			})

			time.Sleep(10 * time.Second)
		}
	}

	jsonBytes, err := json.Marshal(shops)
	fmt.Println(string(jsonBytes))
}

type ShopData struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}
