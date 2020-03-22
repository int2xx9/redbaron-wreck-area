package main

import (
	"io/ioutil"
	"testing"
)

func TestAreaPage(t *testing.T) {
	bytes, err := ioutil.ReadFile("./fixtures/hokkaido.html")
	if err != nil {
		t.FailNow()
	}
	fixture := string(bytes)
	page, err := ParseAreaPage(fixture)
	if err != nil {
		t.FailNow()
	}
	t.Run("ParseAreaPage", func(t *testing.T) {
		page, err := ParseAreaPage(fixture)
		if err != nil {
			t.FailNow()
		}

		if page == nil {
			t.FailNow()
		}
	})
	t.Run("Shops", func(t *testing.T) {
		if len(page.Shops) != 12 {
			t.Fail()
		}
		if page.Shops[0].Name != "レッドバロン北見" {
			t.Errorf("expected: %s, actual: %s",
				"レッドバロン北見",
				page.Shops[0].Name,
			)
		}
		if page.Shops[0].Address != "北海道北見市西三輪6丁目1-14" {
			t.Errorf("expected: %s, actual: %s",
				"北海道北見市西三輪6丁目1-14",
				page.Shops[0].Address,
			)
		}
	})
}

func TestAreaListPage(t *testing.T) {
	bytes, err := ioutil.ReadFile("./fixtures/search.html")
	if err != nil {
		t.FailNow()
	}
	fixture := string(bytes)
	page, err := ParseAreaListPage(fixture)
	if err != nil {
		t.FailNow()
	}
	t.Run("ParseAreaListPage", func(t *testing.T) {
		page, err := ParseAreaListPage(fixture)
		if err != nil {
			t.FailNow()
		}

		if page == nil {
			t.FailNow()
		}
	})
	t.Run("Areas", func(t *testing.T) {
		if len(page.Areas) != 46 {
			t.Fail()
		}
		if page.Areas[0].Name != "北海道" {
			t.Errorf("expected: %s, actual: %s",
				"北海道",
				page.Areas[0].Name,
			)
		}
		if page.Areas[0].URL != "https://www01.redbaron.co.jp/shop/hokkaido/hokkaido/" {
			t.Errorf("expected: %s, actual: %s",
				"https://www01.redbaron.co.jp/shop/hokkaido/hokkaido/",
				page.Areas[0].URL,
			)
		}
	})
}
