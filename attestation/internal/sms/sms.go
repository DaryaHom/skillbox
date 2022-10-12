package sms

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type SMSData struct {
	Country      string `json:"Country"`
	Bandwidth    string `json:"Bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"Provider"`
}

func NewData() *SMSData {
	return &SMSData{}
}

func (d *SMSData) SetCountry(country string) {
	d.Country = country
}

func (d *SMSData) SetBandwidth(bandwidth string) {
	d.Bandwidth = bandwidth
}

func (d *SMSData) SetResponseTime(responseTime string) {
	d.ResponseTime = responseTime
}

func (d *SMSData) SetProvider(provider string) {
	d.Provider = provider
}

func (d *SMSData) SetFullCountryName(fullCountryName string) {
	d.Country = fullCountryName
}

// IsValid - checks the data validity
func IsValid(data []string, alphaCodes map[string]string) (bool, error) {

	// Each line must contain 4 fields
	if len(data) < 4 {
		return false, nil
	}

	// Only countries that have passed the alpha-2 code existence check are allowed in the result
	if _, ok := alphaCodes[data[0]]; !ok {
		return false, nil
	}

	// Only valid providers are allowed in the result
	if data[3] != "Topolo" && data[3] != "Rond" && data[3] != "Kildy" {
		return false, nil
	}
	return true, nil
}

// GetData - reads sms-system status data from CSV file
func GetData(alphaCodes map[string]string) ([]SMSData, error) {
	fmt.Println("****************")
	fmt.Println("SMS-Status:")

	store := make([]SMSData, 0)

	data, err := ioutil.ReadFile("./assets/sms.data")
	if err != nil {
		return store, err
	}

	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		s := NewData()
		d := strings.Split(l, ";")

		if b, err := IsValid(d, alphaCodes); !b || err != nil {
			continue
		}

		// Use valid data to fill structure fields
		s.SetCountry(d[0])
		s.SetBandwidth(d[1])
		s.SetResponseTime(d[2])
		s.SetProvider(d[3])

		store = append(store, *s)
	}

	// Testing the function (stage 2, point 4 of the task)
	for _, v := range store {
		fmt.Printf("%v\n", v)
	}

	return store, nil
}

// SortByProvider - sorts the slice of structures by field "provider"
// Returns copy of received slice
func SortByProvider(smsData []SMSData) []SMSData {
	smsSortedByProvider := make([]SMSData, len(smsData), len(smsData))
	copy(smsSortedByProvider, smsData)

	sort.Slice(smsSortedByProvider, func(i, j int) (less bool) {
		return smsSortedByProvider[i].Provider < smsSortedByProvider[j].Provider
	})
	return smsSortedByProvider
}

// SortByCountry - sorts the slice of structures by field "country"
// Returns copy of received slice
func SortByCountry(smsData []SMSData) []SMSData {
	smsSortedByCountry := make([]SMSData, len(smsData), len(smsData))
	copy(smsSortedByCountry, smsData)

	sort.Slice(smsSortedByCountry, func(i, j int) (less bool) {
		return smsSortedByCountry[i].Country < smsSortedByCountry[j].Country
	})
	return smsSortedByCountry
}

// SetFullCountryName - set full country name from alphaCodes instead of country code
func SetFullCountryName(smsData []SMSData, alphaCodes map[string]string) []SMSData {
	for i := 0; i < len(smsData); i++ {
		smsData[i].SetFullCountryName(alphaCodes[smsData[i].Country])
	}
	return smsData
}

// GetStructuredData - orders the data according to the task condition:
//	- changes the country code to the full name
//	- sorts a copy of the data by provider
//	- sorts a copy of the data by country
func GetStructuredData(alphaCodes map[string]string) ([]SMSData, []SMSData, error) {
	// Read valid sms-data from csv-file
	smsData, err := GetData(alphaCodes)
	if err != nil {
		return nil, nil, err
	}

	// Replace short Country codes to full Country names
	smsData = SetFullCountryName(smsData, alphaCodes)

	// Sort data by Provider & add sorted slice to result-structure
	smsSortedByProvider := SortByProvider(smsData)

	// Sort data by Country & add sorted slice to result-structure
	smsSortedByCountry := SortByCountry(smsData)

	return smsSortedByProvider, smsSortedByCountry, nil
}
