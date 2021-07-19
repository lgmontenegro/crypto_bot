package application

import (
	"lgmontenegro/crypto_bot/internal/config"
	"testing"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Bootstrap(tt.args.config)
		})
	}
}

func TestApplication_Start(t *testing.T) {
	type args struct {
		verbose bool
	}
	tests := []struct {
		name string
		a    *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Start(tt.args.verbose)
		})
	}
}
