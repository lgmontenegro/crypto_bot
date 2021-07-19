package alert

import (
	"fmt"
	ticker "lgmontenegro/crypto_bot/internal/application/ticker_info"
	"testing"
	"time"
)

func TestAlert_Watch(t *testing.T) {
	type args struct {
		firstTicker ticker.Ticker
		ticker      ticker.Ticker
	}
	tests := []struct {
		name            string
		a               *Alert
		args            args
		wantAlarmTicker bool
	}{
		{
			name: "success true",
			a: &Alert{
				Pair: "BTC-USD",
				Perc: 0.5,
				Times: 5 * time.Second,
			},
			args: args{
				firstTicker: ticker.Ticker{
					Ask: 0.100,
					Bid: 0.100,
					Currency: "USD",
					CreatedAt: time.Now().Add(time.Duration(-500 * time.Second)),
				},
				ticker: ticker.Ticker{
					Ask: 0.050,
					Bid: 0.050,
					Currency: "USD",
					CreatedAt: time.Now().Add(time.Duration(-500 * time.Second)),
				},
			},
			wantAlarmTicker: true,
		},		
		{
			name: "success false",
			a: &Alert{
				Pair: "BTC-USD",
				Perc: 0.5,
				Times: 5 * time.Second,
			},
			args: args{
				firstTicker: ticker.Ticker{
					Ask: 0.100,
					Bid: 0.100,
					Currency: "USD",
					CreatedAt: time.Now().Add(time.Duration(-500 * time.Second)),
				},
				ticker: ticker.Ticker{
					Ask: 0.080,
					Bid: 0.080,
					Currency: "USD",
					CreatedAt: time.Now().Add(time.Duration(-500 * time.Second)),
				},
			},
			wantAlarmTicker: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAlarmTicker := tt.a.Watch(tt.args.firstTicker, tt.args.ticker); gotAlarmTicker != tt.wantAlarmTicker {
				fmt.Println(tt.a.Warning.Oscilation)
				t.Errorf("Alert.Watch() = %v, want %v", gotAlarmTicker, tt.wantAlarmTicker)
			}
		})
	}
}
