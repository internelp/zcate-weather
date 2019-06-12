package main

// http://www.weather.com.cn/textFC/hb.shtml
// 这里找城市 id

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Weather 天气信息
type Weather struct {
	Weatherinfo struct {
		City    string `json:"city"`
		Cityid  string `json:"cityid"`
		Temp1   string `json:"temp1"`
		Temp2   string `json:"temp2"`
		Weather string `json:"weather"`
		Img1    string `json:"img1"`
		Img2    string `json:"img2"`
		Ptime   string `json:"ptime"`
	} `json:"weatherinfo"`
}

var weather Weather
var cityID string

func main() {

	flag.StringVar(&cityID, "i", "101010300", "需要查询的 City ID。")
	flag.Parse()
	getWeather()
	badWeather := []string{"雨", "雷", "尘", "霾", "雹", "雪", "雾", "沙"}
	for index := 0; index < len(badWeather); index++ {
		// fmt.Println(badWeather[index])
		if strings.Contains(weather.Weatherinfo.Weather, badWeather[index]) {
			fmt.Printf("%v，今日%v，%v-%v。", weather.Weatherinfo.City, weather.Weatherinfo.Weather, weather.Weatherinfo.Temp1, weather.Weatherinfo.Temp2)
			os.Exit(0)
		}
	}
	log.Fatalf("%v，今日%v，%v-%v。", weather.Weatherinfo.City, weather.Weatherinfo.Weather, weather.Weatherinfo.Temp1, weather.Weatherinfo.Temp2)
}

func getWeather() {
	client := &http.Client{}

	resp, err := client.Get(fmt.Sprintf("http://www.weather.com.cn/data/cityinfo/%s.html", cityID))
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("StatusCode = %v", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalln(err)
	}
}
