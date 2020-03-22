import React, { useState, useEffect } from 'react';
import { render } from 'react-dom';
import { Map, Circle, Marker, Popup, TileLayer, Viewport } from 'react-leaflet';

type ShopData = {
  name: string;
  address: string;
  longitude: number;
  latitude: number;
}

type RedbaronMapProps = {
  centerLatitude: number;
  centerLongitude: number;
  initialZoom: number;
  shops: ShopData[];
  wreckerDistance: number;
  pin: boolean;
}

const RedbaronMap = (props: RedbaronMapProps) => {
  const [zoom, setZoom] = useState<number>(props.initialZoom);
  const viewportChangedHandler = (v: Viewport) => {
    if (!v.zoom) return;
    setZoom(v.zoom);
  };

  const circles = props.shops.map(shop => (
    <Circle
      key={`circle${shop.name}`}
      center={[shop.latitude, shop.longitude]}
      fillColor="blue"
      radius={props.wreckerDistance} />
  ));

  const markers = props.shops.map(shop => (
    <Marker key={`marker${shop.name}`} position={[shop.latitude, shop.longitude]}>
      <Popup>Name: {shop.name}<br />Address: {shop.address}</Popup>
    </Marker>
  ));

  return (
    <Map
      style={{width: '100vw', height: '100vh'}}
      center={[props.centerLatitude, props.centerLongitude]}
      onViewportChanged={viewportChangedHandler}
      zoom={props.initialZoom}>
      <TileLayer
        attribution='&amp;copy <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      {circles}
      {props.pin && zoom >= 10 && markers}
    </Map>
  );
}

type AppQueryHash = {
  distance: string;
}

const parseQueryHash = <T extends unknown>() => {
  const rawParams = location.hash.indexOf('#') === 0 ? location.hash.slice(1) : "";
  const entries = new URLSearchParams(rawParams).entries();
  const query = {} as T;
  for (const [key, value] of entries) {
    (query as any)[key] = value;
  }
  return query;
};

const App = () => {
  const queryHash = parseQueryHash<AppQueryHash>();
  const distance = queryHash.distance ? parseInt(queryHash.distance) : 50;
  const [shops, setShops] = useState<ShopData[]>([]);
  useEffect(() => {
    (async () => {
      var data = await fetch('/shops.json');
      var json = await data.json();
      setShops(json.map((record: any) => ({
        name: record.name,
        address: record.address,
        longitude: parseFloat(record.longitude),
        latitude: parseFloat(record.latitude),
      })));
    })();
  }, []);

  return (
    <div>
      <RedbaronMap
        centerLatitude={38.554826}
        centerLongitude={135.858479}
        initialZoom={6}
        pin={true}
        wreckerDistance={distance * 1000}
        shops={shops} />
    </div>
  );
}

window.addEventListener('DOMContentLoaded', () => {
  render(<App />, document.getElementById('app'));
});
