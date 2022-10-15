package mms

import (
	"attestation/internal/data"
	"fmt"
	"sort"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

// GetData - gets mms-system status data from simulator API
func GetData(host, simulatorAddr string) ([]MMSData, error) {
	fmt.Println()
	fmt.Println("****************")
	fmt.Println("MMS-Status:")

	var store []MMSData

	err := data.GetFromAPI(host, simulatorAddr, "/mms", &store)
	if err != nil {
		return store, err
	}

	return store, nil
}

// IsValid - checks the data validity
func (d MMSData) IsValid(alphaCodes map[string]string) (bool, error) {

	// Only countries that have passed the alpha-2 code existence check are allowed in the result
	if _, ok := alphaCodes[d.Country]; !ok {
		return false, nil
	}

	// Only valid providers are allowed in the result
	if d.Provider != "Topolo" && d.Provider != "Rond" && d.Provider != "Kildy" {
		return false, nil
	}
	return true, nil
}

// ValidData - makes a slice from data checked by IsValid-function
func ValidData(store []MMSData, alphaCodes map[string]string) []MMSData {
	var validStore []MMSData
	for _, d := range store {
		if b, err := d.IsValid(alphaCodes); !b || err != nil {
			continue
		}
		validStore = append(validStore, d)
	}
	return validStore
}

// SortByProvider - sorts the slice of structures by field "provider"
// Returns copy of received slice
func SortByProvider(mmsData []MMSData) []MMSData {
	mmsSortedByProvider := make([]MMSData, len(mmsData), len(mmsData))
	copy(mmsSortedByProvider, mmsData)

	sort.Slice(mmsSortedByProvider, func(i, j int) (less bool) {
		return mmsSortedByProvider[i].Provider < mmsSortedByProvider[j].Provider
	})
	return mmsSortedByProvider
}

// SortByCountry - sorts the slice of structures by field "country"
// Returns copy of received slice
func SortByCountry(mmsData []MMSData) []MMSData {
	mmsSortedByCountry := make([]MMSData, len(mmsData), len(mmsData))
	copy(mmsSortedByCountry, mmsData)

	sort.Slice(mmsSortedByCountry, func(i, j int) (less bool) {
		return mmsSortedByCountry[i].Country < mmsSortedByCountry[j].Country
	})
	return mmsSortedByCountry
}

// SetFullCountryName - set full country name from alphaCodes instead of country code
func SetFullCountryName(mmsData []MMSData, alphaCodes map[string]string) []MMSData {
	for i := 0; i < len(mmsData); i++ {
		mmsData[i].Country = alphaCodes[mmsData[i].Country]
	}
	return mmsData
}

// GetStructuredData - orders the data according to the task condition:
//	- changes the country code to the full country name
//	- sorts a copy of the data by provider
//	- sorts a copy of the data by country
func GetStructuredData(alphaCodes map[string]string, host, simulatorAddr string) ([]MMSData, []MMSData, error) {

	mmsData, err := GetData(host, simulatorAddr)
	if err != nil {
		return nil, nil, err
	}

	validMmsData := ValidData(mmsData, alphaCodes)

	// Testing the function (stage 3 of the task)
	for _, d := range validMmsData {
		fmt.Printf("%v\n", d)
	}

	// Replace short country codes to full country names
	mmsData = SetFullCountryName(mmsData, alphaCodes)

	// Sort data by provider & add sorted slice to result-structure
	mmsSortedByProvider := SortByProvider(mmsData)

	// Sort data by country & add sorted slice to result-structure
	mmsSortedByCountry := SortByCountry(mmsData)

	return mmsSortedByProvider, mmsSortedByCountry, nil
}
