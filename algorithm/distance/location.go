package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// URL 緯度経度サービス
const URL = "http://geoapi.heartrails.com/api/json"

// GeoResponse レスポンス構造体
type GeoResponse struct {
	Response struct {
		Location []Location `json:"location"`
	} `json:"response"`
}

// Location ロケーション
type Location struct {
	City       string `json:"city"`
	CityKana   string `json:"city_kana"`
	Town       string `json:"town"`
	TownKana   string `json:"town_kana"`
	Longitude  string `json:"x"`
	Latitude   string `json:"y"`
	Prefecture string `json:"prefecture"`
	Postal     string `json:"postal"`
}

// GeoLocation キーワード検索
func GeoLocation(keyword string) []Location {

	req, err := http.NewRequest(
		"GET",
		URL,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	params := req.URL.Query()

	params.Add("method", "suggest")
	params.Add("keyword", keyword)
	params.Add("matching", "like")

	req.URL.RawQuery = params.Encode()

	timeout := time.Duration(5 * time.Second)
	clienct := &http.Client{
		Timeout: timeout,
	}

	res, err := clienct.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var response GeoResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err.Error())
	}

	return response.Response.Location
}
