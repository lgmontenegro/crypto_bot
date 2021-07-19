package data_processor

import (
	"sync"
	"testing"
)

func TestDataProcessor_Process(t *testing.T) {
	type args struct {
		exec    Data
		pair    string
		wg      sync.WaitGroup
		verbose bool
	}
	tests := []struct {
		name    string
		d       *DataProcessor
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Process(tt.args.exec, tt.args.pair, tt.args.wg, tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("DataProcessor.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
