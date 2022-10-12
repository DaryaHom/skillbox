package test

import (
	"attestation/internal/mms"
	"testing"
)

func TestData_IsValid(t *testing.T) {
	alphaCodes := map[string]string{"AD": "Andorra", "AE": "United Arab Emirates", "AM": "Armenia", "BM": "Bermuda",
		"CA": "Canada", "CN": "China", "CO": "Colombia", "EE": "Estonia", "EG": "Egypt", "ES": "Spain",
		"GB": "United Kingdom", "GE": "Georgia", "GR": "Greece", "ID": "Indonesia", "IL": "Israel", "JP": "Japan",
		"KE": "Kenya", "KG": "Kyrgyzstan", "LV": "Latvia", "RU": "Russian Federation", "TW": "Taiwan",
		"US": "United States", "UZ": "Uzbekistan", "VA": "Holy See (Vatican City State)", "ZM": "Zambia", "ZW": "Zimbabwe"}

	d := []mms.MMSData{
		{
			Country:      "RU",
			Provider:     "Topolo",
			Bandwidth:    "79",
			ResponseTime: "450",
		},
		{
			Country:      "US",
			Provider:     "Rond",
			Bandwidth:    "36",
			ResponseTime: "1607",
		},
		{
			Country:      "GB",
			Provider:     "Topolo GB",
			Bandwidth:    "76",
			ResponseTime: "115",
		},
		{
			Country:      "U5",
			Provider:     "Topolo",
			Bandwidth:    "43",
			ResponseTime: "1448",
		},
	}

	tests := []struct {
		name       string
		data       mms.MMSData
		alphaCodes map[string]string
		want       bool
		wantErr    bool
	}{
		{
			name:       "1",
			data:       d[0],
			alphaCodes: alphaCodes,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "2",
			data:       d[1],
			alphaCodes: alphaCodes,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "3",
			data:       d[2],
			alphaCodes: alphaCodes,
			want:       false,
			wantErr:    false,
		},
		{
			name:       "4",
			data:       d[3],
			alphaCodes: alphaCodes,
			want:       false,
			wantErr:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.data.IsValid(tc.alphaCodes)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidData() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("IsValid() got = %v, want %v", got, tc.want)
			}
		})
	}
}
