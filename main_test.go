package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeatherDescriptions(t *testing.T) {
	tests := []struct {
		code string
		want string
	}{
		{"113", "晴れ"},
		{"119", "曇り"},
		{"305", "大雨"},
		{"329", "雪"},
	}
	for _, tt := range tests {
		got, ok := weatherDescriptions[tt.code]
		if !ok {
			t.Errorf("weatherDescriptions[%q] が存在しない", tt.code)
			continue
		}
		if got != tt.want {
			t.Errorf("weatherDescriptions[%q] = %q, want %q", tt.code, got, tt.want)
		}
	}
}

func TestFetchWeather(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		wantResult  string
		wantErr     bool
	}{
		{
			name: "正常系",
			body: `{"current_condition":[{"temp_C":"20","weatherCode":"113","humidity":"50"}]}`,
			wantResult: "東京の天気: 晴れ、気温: 20°C、湿度: 50%",
		},
		{
			name: "未知の天気コード",
			body: `{"current_condition":[{"temp_C":"15","weatherCode":"999","humidity":"60"}]}`,
			wantResult: "東京の天気: 不明、気温: 15°C、湿度: 60%",
		},
		{
			name:    "空のレスポンス",
			body:    `{"current_condition":[]}`,
			wantErr: true,
		},
		{
			name:    "不正なJSON",
			body:    `invalid`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tt.body))
			}))
			defer srv.Close()

			original := wttrBaseURL
			wttrBaseURL = srv.URL
			defer func() { wttrBaseURL = original }()

			got, err := fetchWeather("Tokyo")
			if tt.wantErr {
				if err == nil {
					t.Errorf("エラーを期待したが nil だった")
				}
				return
			}
			if err != nil {
				t.Errorf("予期しないエラー: %v", err)
				return
			}
			if got != tt.wantResult {
				t.Errorf("got %q, want %q", got, tt.wantResult)
			}
		})
	}
}
