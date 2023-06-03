package eleven

import (
	"testing"
)

func TestEleven_ListVoices(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "ListVoices",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			_, err := e.ListVoices()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListVoices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
