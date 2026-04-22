package vin

import (
	"errors"
	"testing"
)

func Test_parse_lengthCheck(t *testing.T) {
	// given
	vin := "1234567890"

	// when
	_, err := NewVINFromString(vin)

	// then
	if !errors.Is(err, ErrInvalidVINLength) {
		t.Errorf("expected ErrInvalidVINLength, got %v", err)
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		name string
		vin string
		wantErr bool
		expectedRegion string
	} {
		{
            name:           "valid VIN",
            vin:            "WBADT43452G915187",
            wantErr:        false,
            expectedRegion: "W",
        },
        {
            name:           "too short",
            vin:            "1234567890",
            wantErr:        true,
            expectedRegion: "",
        },
        {
            name:           "too long",
            vin:            "WBADT43452G915187EXTRA",
            wantErr:        true,
            expectedRegion: "",
        },
        {
            name:           "empty string",
            vin:            "",
            wantErr:        true,
            expectedRegion: "",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultingVIN, err := NewVINFromString(tt.vin)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewVINFromString() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && !errors.Is(err, ErrInvalidVINLength) {
				t.Errorf("expected ErrInvalidVINLength, got %v", err)
			}

			if !tt.wantErr && resultingVIN.region != tt.expectedRegion {
				t.Errorf("expected region %s, got %s", tt.expectedRegion, resultingVIN.region)
			}
		})
	}
}