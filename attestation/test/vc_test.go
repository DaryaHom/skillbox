package test

import (
	"attestation/internal/vc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsValid(t *testing.T) {
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
			data:      []string{"U5", "32", "96", "E-Voice", "0.88", "845", "48", "47"},
			alphaCode: alphaCode,
			want:      false,
		},
		{
			name:      "2",
			data:      []string{"RU", "91", "1843", "TransparentCalls", "0.62", "562", "26", "39"},
			alphaCode: alphaCode,
			want:      true,
		},
		{
			name:      "3",
			data:      []string{"GB", "98", "903", "TransparentCalls", "0.67", "25", "23"},
			alphaCode: alphaCode,
			want:      false,
		},
		{
			name:      "4",
			data:      []string{"CA", "51", "813", "JustPhon", "0.61", "349", "46", "25"},
			alphaCode: alphaCode,
			want:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := vc.IsValid(tc.data, tc.alphaCode)

			assert.Equal(t, tc.want, res)
		})
	}
}
