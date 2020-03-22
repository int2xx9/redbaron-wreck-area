package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ShopSummary struct {
	Name    string
	Address string
}

type AreaPage struct {
	Shops []ShopSummary
}

func ParseAreaPage(page string) (*AreaPage, error) {
	areaPage := AreaPage{
		Shops: []ShopSummary{},
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		return nil, err
	}
	pref, err := doc.Find("#breadcrumbs ul li").Last().Html()
	if err != nil {
		return nil, err
	}
	doc.Find(".listBox").Each(func(_ int, s *goquery.Selection) {
		name := s.Find("h3 a").First()
		address := s.Find("table tr td").First()
		areaPage.Shops = append(areaPage.Shops, ShopSummary{
			Name:    name.Text(),
			Address: pref + string([]rune(address.Text())[1:]),
		})
	})
	return &areaPage, nil
}

type AreaUrl struct {
	Name string
	URL  string
}

type AreaListPage struct {
	Areas []AreaUrl
}

func ParseAreaListPage(page string) (*AreaListPage, error) {
	areaListPage := AreaListPage{
		Areas: []AreaUrl{},
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		return nil, err
	}
	doc.Find(".searchBlock ul li a").Each(func(_ int, s *goquery.Selection) {
		v, exists := s.Attr("href")
		if !exists {
			return
		}

		baseURL, err := url.Parse(areaListUrl)
		if err != nil {
			return
		}

		relativePath, err := url.Parse(v)
		if err != nil {
			return
		}

		areaListPage.Areas = append(areaListPage.Areas, AreaUrl{
			Name: s.Text(),
			URL:  baseURL.ResolveReference(relativePath).String(),
		})
	})
	return &areaListPage, nil
}
