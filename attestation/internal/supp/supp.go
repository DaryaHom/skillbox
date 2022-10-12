package supp

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// GetData - gets support-system status data from simulator API
func GetData(host, simulatorAddr string) ([]SupportData, error) {
	fmt.Println()
	fmt.Println("****************")
	fmt.Println("Support Status:")

	var store []SupportData

	resp, err := http.Get(host + simulatorAddr + "/support")
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

// GetOpenedTickets - counts the number of active tickets
func GetOpenedTickets(suppData []SupportData) int {
	sum := 0
	for _, datum := range suppData {
		sum += datum.ActiveTickets
	}
	return sum
}

// CheckLoad - returns the status of the load on the support based on:
// - up to 9 tickets - the support is not loaded,
// - 9 - 16 - the support is moderately loaded,
// - > 16 - the support is overloaded.
func CheckLoad(suppData []SupportData) int {
	sum := GetOpenedTickets(suppData)

	load := 0
	switch {
	case sum < 9:
		load = 1
	case sum >= 9 && sum <= 16:
		load = 2
	case sum > 16:
		load = 3
	}
	return load
}

// GetWaitingTime - returns potential response time for a new ticket
func GetWaitingTime(suppData []SupportData) int {
	timeForTicket := 60.0 / 18.0
	ticketsCount := GetOpenedTickets(suppData)
	waitingTime := math.Round(timeForTicket * float64(ticketsCount))
	return int(waitingTime)
}

// GetStructuredData - orders the data according to the task condition:
func GetStructuredData(host, simulatorAddr string) (int, int, error) {

	suppData, err := GetData(host, simulatorAddr)

	if err != nil {
		return -1, -1, err
	}

	load := CheckLoad(suppData)
	waitingTime := GetWaitingTime(suppData)

	return load, waitingTime, nil
}
