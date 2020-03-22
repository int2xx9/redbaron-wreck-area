#!/bin/sh

basedir="$(cd "$(dirname "$0")"; pwd)"
cd "$basedir/fixtures"
curl -o tokyo_station.xml "https://www.geocoding.jp/api/?q=%E6%9D%B1%E4%BA%AC%E9%A7%85"
curl -o hokkaido.html "https://www01.redbaron.co.jp/shop/hokkaido/hokkaido/"
curl -o search.html "https://www01.redbaron.co.jp/shop/search/"

