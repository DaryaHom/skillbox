package test

import (
	"attestation/internal/sms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	alphaCode := map[string]string{"AD": "Andorra", "AE": "United Arab Emirates", "AM": "Armenia", "BM": "Bermuda",
		"CA": "Canada", "CN": "China", "CO": "Colombia", "EE": "Estonia", "EG": "Egypt", "ES": "Spain",
		"GB": "United Kingdom", "GE": "Georgia", "GR": "Greece", "ID": "Indonesia", "IL": "Israel", "JP": "Japan",
		"KE": "Kenya", "KG": "Kyrgyzstan", "LV": "Latvia", "RU": "Russian Federation", "TW": "Taiwan",
		"US": "United States", "UZ": "Uzbekistan", "VA": "Holy See (Vatican City State)", "ZM": "Zambia", "ZW": "Zimbabwe"}

	testCases := []struct {
		name      string
		data      []string
		alphaCode map[string]string
		want      bool
	}{
		{
			name:      "1",
			data:      []string{"U5", "41910", "Topol"},
			alphaCode: alphaCode,
			want:      false,
		},
		{
			name:      "2",
			data:      []string{"US", "36", "1576", "Rond"},
			alphaCode: alphaCode,
			want:      true,
		},
		{
			name:      "3",
			data:      []string{"GB28495Topolo"},
			alphaCode: alphaCode,
			want:      false,
		},
		{
			name:      "4",
			data:      []string{"BL", "68", "1594", "Kildy US"},
			alphaCode: alphaCode,
			want:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := sms.IsValid(tc.data, tc.alphaCode)

			assert.Equal(t, tc.want, res)
		})
	}
}
