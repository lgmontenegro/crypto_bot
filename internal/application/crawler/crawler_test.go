package crawler

import (
	"reflect"
	"testing"
)

func TestCrawler_DataRetriever(t *testing.T) {
	tests := []struct {
		name     string
		c        *Crawler
		wantBody []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, err := tt.c.DataRetriever()
			if (err != nil) != tt.wantErr {
				t.Errorf("Crawler.DataRetriever() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("Crawler.DataRetriever() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}
