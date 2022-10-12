package inc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // possible statuses active and closed
}

// GetData - gets accident data from simulator API
func GetData(host, simulatorAddr string) ([]IncidentData, error) {
	fmt.Println()
	fmt.Println("****************")
	fmt.Println("Accident Status:")

	var store []IncidentData

	resp, err := http.Get(host + simulatorAddr + "/accendent")
	if err != nil {
		return store, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return store, err
		}

		err = json.Unmarshal(body, &store)
		if err != nil {
			return store, err
		}
	}

	// Testing function
	for _, d := range store {
		fmt.Printf("%v\n", d)
	}

	return store, nil
}

// SortByStatus - sorts the slice of structures by field "Status"
// Returns copy of received slice
func SortByStatus(incData []IncidentData) []IncidentData {
	incSortedByStatus := make([]IncidentData, len(incData), len(incData))
	copy(incSortedByStatus, incData)

	sort.Slice(incSortedByStatus, func(i, j int) (less bool) {
		return incSortedByStatus[i].Status < incSortedByStatus[j].Status
	})
	return incSortedByStatus
}

// GetStructuredData - orders the data according to the task condition:
func GetStructuredData(host, simulatorAddr string) ([]IncidentData, error) {

	incData, err := GetData(host, simulatorAddr)
	if err != nil {
		return nil, err
	}

	sortedIncData := SortByStatus(incData)
	return sortedIncData, nil
}
