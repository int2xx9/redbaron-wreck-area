redbaron-wrecker-area
==================================================

A tool to check covered area of roadside assistance.

## How to run

```
$ yarn webpack-dev-server --content-base public
$ open http://localhost:8080/
$ open 'http://localhost:8080/#distance=10'  # set a radius 10km (default 50km)
```

## How to update shop list

```
$ cd tools/fetch-shop-list
$ go build
$ ./fetch-shop-list > ../../public/shops.json
```

