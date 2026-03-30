package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type wttrResponse struct {
	CurrentCondition []struct {
		TempC       string `json:"temp_C"`
		WeatherCode string `json:"weatherCode"`
		Humidity    string `json:"humidity"`
	} `json:"current_condition"`
}

var weatherDescriptions = map[string]string{
	"113": "晴れ",
	"116": "一部曇り",
	"119": "曇り",
	"122": "曇り",
	"143": "霧",
	"176": "小雨",
	"179": "小雪",
	"182": "みぞれ",
	"185": "着氷性の霧雨",
	"200": "雷雨",
	"227": "地吹雪",
	"230": "吹雪",
	"248": "霧",
	"260": "着氷性の霧",
	"263": "霧雨",
	"266": "霧雨",
	"281": "着氷性の霧雨",
	"284": "着氷性の霧雨",
	"293": "小雨",
	"296": "小雨",
	"299": "雨",
	"302": "雨",
	"305": "大雨",
	"308": "大雨",
	"311": "着氷性の雨",
	"314": "着氷性の雨",
	"317": "みぞれ",
	"320": "みぞれ",
	"323": "小雪",
	"326": "小雪",
	"329": "雪",
	"332": "雪",
	"335": "大雪",
	"338": "大雪",
	"350": "あられ",
	"353": "にわか雨",
	"356": "強いにわか雨",
	"359": "豪雨",
	"362": "みぞれのにわか雨",
	"365": "みぞれのにわか雨",
	"368": "にわか雪",
	"371": "強いにわか雪",
	"374": "あられのにわか雨",
	"377": "あられのにわか雨",
	"386": "雷雨",
	"389": "激しい雷雨",
	"392": "雷雪",
	"395": "激しい雷雪",
}

func fetchWeather(location string) (string, error) {
	resp, err := http.Get("https://wttr.in/" + location + "?format=j1")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data wttrResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if len(data.CurrentCondition) == 0 {
		return "", fmt.Errorf("天気データが取得できませんでした")
	}

	cond := data.CurrentCondition[0]
	desc, ok := weatherDescriptions[cond.WeatherCode]
	if !ok {
		desc = "不明"
	}

	return fmt.Sprintf("東京の天気: %s、気温: %s°C、湿度: %s%%", desc, cond.TempC, cond.Humidity), nil
}

func main() {
	now := time.Now()
	weekdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	fmt.Printf("おはよう\n今日は%d年%d月%d日(%s曜日)です\n", now.Year(), now.Month(), now.Day(), weekdays[now.Weekday()])

	weather, err := fetchWeather("Tokyo")
	if err != nil {
		fmt.Printf("天気の取得に失敗しました: %v\n", err)
		return
	}
	fmt.Println(weather)
}
