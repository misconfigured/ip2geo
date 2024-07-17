# Serve MaxMind DB as a JSON endpoint

## Build
    go get ./...
    go build

## Run
    ./ip2geo

## Tests
    go get -t ./...
    go test

## Usage
    curl -s -d '{"ip": "4.16.85.243"}' -X POST http://localhost:8080/location/resolve | jq .

## Notes
If you pass bad data in, it will return you an empty JSON object. Otherwise your results will look like the following

```{
  "City": {
    "GeoNameID": 5367440,
    "Names": {
      "de": "Livermore",
      "en": "Livermore",
      "fr": "Livermore",
      "ja": "リバモア",
      "ru": "Ливермор",
      "zh-CN": "利佛摩"
    }
  },
  "Continent": {
    "Code": "NA",
    "GeoNameID": 6255149,
    "Names": {
      "de": "Nordamerika",
      "en": "North America",
      "es": "Norteamérica",
      "fr": "Amérique du Nord",
      "ja": "北アメリカ",
      "pt-BR": "América do Norte",
      "ru": "Северная Америка",
      "zh-CN": "北美洲"
    }
  },
  "Country": {
    "GeoNameID": 6252001,
    "IsInEuropeanUnion": false,
    "IsoCode": "US",
    "Names": {
      "de": "USA",
      "en": "United States",
      "es": "Estados Unidos",
      "fr": "États-Unis",
      "ja": "アメリカ合衆国",
      "pt-BR": "Estados Unidos",
      "ru": "США",
      "zh-CN": "美国"
    }
  },
  "Location": {
    "AccuracyRadius": 50,
    "Latitude": 37.5038,
    "Longitude": -121.5253,
    "MetroCode": 807,
    "TimeZone": "America/Los_Angeles"
  },
  "Postal": {
    "Code": "94550"
  },
  "RegisteredCountry": {
    "GeoNameID": 6252001,
    "IsInEuropeanUnion": false,
    "IsoCode": "US",
    "Names": {
      "de": "USA",
      "en": "United States",
      "es": "Estados Unidos",
      "fr": "États-Unis",
      "ja": "アメリカ合衆国",
      "pt-BR": "Estados Unidos",
      "ru": "США",
      "zh-CN": "美国"
    }
  },
  "RepresentedCountry": {
    "GeoNameID": 0,
    "IsInEuropeanUnion": false,
    "IsoCode": "",
    "Names": null,
    "Type": ""
  },
  "Subdivisions": [
    {
      "GeoNameID": 5332921,
      "IsoCode": "CA",
      "Names": {
        "de": "Kalifornien",
        "en": "California",
        "es": "California",
        "fr": "Californie",
        "ja": "カリフォルニア州",
        "pt-BR": "Califórnia",
        "ru": "Калифорния",
        "zh-CN": "加利福尼亚州"
      }
    }
  ],
  "Traits": {
    "IsAnonymousProxy": false,
    "IsSatelliteProvider": false
  }
}
```
