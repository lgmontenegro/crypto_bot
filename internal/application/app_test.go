package application

import (
	"lgmontenegro/crypto_bot/internal/config"
	"testing"
	"time"
)

func TestApplication_Bootstrap(t *testing.T) {
	type args struct {
		config config.Config
	}
	tests := []struct {
		name string
		a    *Application
		args args
	}{
		{
			name: "return complete app",
			a: &Application{},
			args: args{
				config: config.Config{
					Times: 5*time.Second,
					URL: "http://localhost/",
					Endpoint: "v1/ticker/",
					Pairs: []string{"BTC-ADA", "ADA-USD"},
					Alerts: []config.Alert{
						{
							Pair: "BTC-ADA",
							Perc: 0.01,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.a.Bootstrap(tt.args.config); tt.a.dataProcessor.Alerts[0].Pair != tt.args.config.Alerts[0].Pair {
				t.Errorf("Application.Bootstrap() not settle up correctly")
			}
		})
	}
}
