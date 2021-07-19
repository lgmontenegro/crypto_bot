package alert

import (
	ticker "lgmontenegro/crypto_bot/internal/application/ticker_info"
	"testing"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAlarmTicker := tt.a.Watch(tt.args.firstTicker, tt.args.ticker); gotAlarmTicker != tt.wantAlarmTicker {
				t.Errorf("Alert.Watch() = %v, want %v", gotAlarmTicker, tt.wantAlarmTicker)
			}
		})
	}
}
