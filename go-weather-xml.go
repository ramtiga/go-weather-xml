package main

import (
    "fmt"
    "encoding/xml"
    "net/http"
    "io/ioutil"
    "log"
)
var FEED_URL string = "http://weather.livedoor.com/forecast/rss/area/130010.xml"

type WeatherHack struct {
    Title string `xml:"channel>title"`
    Description []string `xml:"channel>item>description"`
}

func main() {
    wh, err := getWeather(FEED_URL)

    if err != nil {
        log.Fatalf("Log: %v", err)
        return
    }

    fmt.Println(wh.Title)
    for n, v := range wh.Description {
        if n > 0 {
            fmt.Printf("%s \n", v)
        }
    }
}


func getWeather(feed string) (p *WeatherHack, err error) {

    res, err := http.Get(feed)
    if err != nil {
        log.Fatalf("Log: %v", err)
        return
    }

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("Log: %v", err)
        return
    }
    wh := new(WeatherHack)
    err = xml.Unmarshal(b, &wh)

    return wh, err
}
