package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatNumberWithSuffix(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{
			name:  "format in billion",
			value: 2_050_000_000,
			want:  "2.05 billion",
		},
		{
			name:  "format in million",
			value: 500_050_000,
			want:  "500.05 million",
		},
		{
			name:  "format in thousand",
			value: 40_050,
			want:  "40.05 thousand",
		},
		{
			name:  "format in number",
			value: 40,
			want:  "40",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatNumberWithSuffix(tt.value)

			assert.Equal(t, tt.want, got)
		})
	}

}
