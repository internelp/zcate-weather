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
	Data struct {
		Yesterday struct {
			Date string `json:"date"`
			High string `json:"high"`
			Fx   string `json:"fx"`
			Low  string `json:"low"`
			Fl   string `json:"fl"`
			Type string `json:"type"`
		} `json:"yesterday"`
		City     string `json:"city"`
		Forecast []struct {
			Date      string `json:"date"`
			High      string `json:"high"`
			Fengli    string `json:"fengli"`
			Low       string `json:"low"`
			Fengxiang string `json:"fengxiang"`
			Type      string `json:"type"`
		} `json:"forecast"`
		Ganmao string `json:"ganmao"`
		Wendu  string `json:"wendu"`
	} `json:"data"`
	Status int    `json:"status"`
	Desc   string `json:"desc"`
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
		if strings.Contains(weather.Data.Forecast[0].Type, badWeather[index]) {
			fmt.Printf("%v，今日%v，%v，%v。", weather.Data.City, weather.Data.Forecast[0].Type, weather.Data.Forecast[0].Low, weather.Data.Forecast[0].High)
			os.Exit(0)
		}
	}
	log.Fatalf("%v，今日%v，%v，%v。", weather.Data.City, weather.Data.Forecast[0].Type, weather.Data.Forecast[0].Low, weather.Data.Forecast[0].High)
}

func getWeather() {
	client := &http.Client{}

	resp, err := client.Get(fmt.Sprintf("http://wthrcdn.etouch.cn/weather_mini?citykey=%s", cityID))
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
