package mail

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func NewData() *EmailData {
	return &EmailData{}
}

func (d *EmailData) SetCountry(country string) {
	d.Country = country
}

func (d *EmailData) SetProvider(provider string) {
	d.Provider = provider
}

func (d *EmailData) SetDeliveryTime(deliveryTime int) {
	d.DeliveryTime = deliveryTime
}

// IsValid - checks the data validity
func IsValid(data []string, alphaCodes map[string]string) (bool, error) {

	// Each line must contain 3 fields
	if len(data) < 3 {
		return false, nil
	}

	// Check the validity of the country data
	if _, ok := alphaCodes[data[0]]; !ok {
		return false, nil
	}

	// Create set of valid email providers
	type Void struct{}
	void := Void{}
	providers := map[string]Void{
		"Gmail": void, "Yahoo": void, "Hotmail": void, "MSN": void, "Orange": void, "Comcast": void,
		"AOL": void, "Live": void, "RediffMail": void, "GMX": void, "Protonmail": void, "Yandex": void, "Mail.ru": void,
	}

	// Check the validity of the provider data
	if _, ok := providers[data[1]]; !ok {
		return false, nil
	}

	// Check the validity of the DeliveryTime indicator
	if delTime, err := strconv.Atoi(data[2]); err != nil || delTime == 0 {
		return false, nil
	}

	return true, nil
}

// GetStatus - function to get valid sms-system status data from CSV file
func GetStatus(alphaCodes map[string]string) ([]EmailData, error) {
	fmt.Println()
	fmt.Println("****************")
	fmt.Println("Email-Status:")

	// Create a store to keep valid values
	store := make([]EmailData, 0)

	// Read data from file
	data, err := ioutil.ReadFile("../attestation/assets/email.data")
	if err != nil {
		return store, err
	}

	// Split data & check the validity of each line
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		e := NewData()
		d := strings.Split(l, ";")

		// If the line isn't valid, check the next one.
		// Otherwise, enter the data into the EmailData-structure (e) and add it to the store
		if b, err := IsValid(d, alphaCodes); !b || err != nil {
			continue
		}

		e.SetCountry(d[0])
		e.SetProvider(d[1])

		deliveryTime, _ := strconv.Atoi(d[2])
		e.SetDeliveryTime(deliveryTime)

		store = append(store, *e)
	}

	// Testing the function
	for _, m := range store {
		fmt.Printf("%v\n", m)
	}

	return store, nil
}

// SortByTimeAsc - sorts the slice of structures by field "DeliveryTime"
// Returns copy of received slice
func SortByTimeAsc(mailData []EmailData) []EmailData {
	mailSortedByTime := make([]EmailData, len(mailData), len(mailData))
	copy(mailSortedByTime, mailData)

	sort.Slice(mailSortedByTime, func(i, j int) (less bool) {
		return mailSortedByTime[i].DeliveryTime < mailSortedByTime[j].DeliveryTime
	})
	return mailSortedByTime
}

// GetStructuredData - returns map with 2 slices inside.
// The first slice contains the 3 fastest providers, the second - the 3 slowest
func GetStructuredData(alphaCodes map[string]string) (map[string][][]EmailData, error) {
	mailData, err := GetStatus(alphaCodes)
	if err != nil {
		return nil, err
	}

	// Sort data by delivery time
	mailSortedByTime := SortByTimeAsc(mailData)

	// Create map to store sorted values
	m := make(map[string][][]EmailData)

	// Run a loop on sorted email-data
	for _, datum := range mailSortedByTime {

		// If the country is not yet in the storage, then 2 slices are formed.
		// The first one for list of providers with the minimum delivery time, the second one - for the maximum one
		if _, ok := m[datum.Country]; !ok {
			arr := make([][]EmailData, 2)

			arrAsc := make([]EmailData, 0)
			for i, counter := 0, 0; i < len(mailSortedByTime) && counter < 3; i++ {
				if mailSortedByTime[i].Country == datum.Country {
					arrAsc = append(arrAsc, mailSortedByTime[i])
					counter++
				}
			}

			arrDesc := make([]EmailData, 0)
			for i, counter := len(mailSortedByTime)-1, 0; i >= 0 && counter < 3; i-- {
				if mailSortedByTime[i].Country == datum.Country {
					arrDesc = append(arrDesc, mailSortedByTime[i])
					counter++
				}
			}

			// Add results to the store
			arr[0], arr[1] = arrAsc, arrDesc
			m[datum.Country] = arr
		}
	}
	return m, nil
}
