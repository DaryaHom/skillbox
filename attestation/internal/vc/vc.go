package vc

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"Country"`
	Bandwidth           int     `json:"Bandwidth"`
	ResponseTime        int     `json:"response_time"`
	Provider            string  `json:"Provider"`
	ConnectionStability float64 `json:"connection_stability"`
	Ttfb                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func NewData() *VoiceCallData {
	return &VoiceCallData{}
}

func (d *VoiceCallData) SetCountry(country string) {
	d.Country = country
}

func (d *VoiceCallData) SetBandwidth(bandwidth int) {
	d.Bandwidth = bandwidth
}

func (d *VoiceCallData) SetResponseTime(responseTime int) {
	d.ResponseTime = responseTime
}

func (d *VoiceCallData) SetProvider(provider string) {
	d.Provider = provider
}

func (d *VoiceCallData) SetConnectionStability(connectionStability float64) {
	d.ConnectionStability = connectionStability
}

func (d *VoiceCallData) SetTTFB(ttfb int) {
	d.Ttfb = ttfb
}

func (d *VoiceCallData) SetVoicePurity(voicePurity int) {
	d.VoicePurity = voicePurity
}

func (d *VoiceCallData) SetMedianOfCallsTime(medianOfCallsTime int) {
	d.MedianOfCallsTime = medianOfCallsTime
}

// IsValid - checks the data validity
func IsValid(data []string, alphaCodes map[string]string) (bool, error) {

	// Each line must contain 8 fields
	if len(data) < 8 {
		return false, nil
	}

	// Check the validity of the Country data
	if _, ok := alphaCodes[data[0]]; !ok {
		return false, nil
	}

	// Check the validity of the Bandwidth indicator
	if _, err := strconv.Atoi(data[1]); err != nil {
		return false, nil
	}

	// Check the validity of the ResponseTime indicator
	if _, err := strconv.Atoi(data[2]); err != nil {
		return false, nil
	}

	//  Check the validity of the Provider data
	if data[3] != "TransparentCalls" && data[3] != "E-Voice" && data[3] != "JustPhone" {
		return false, nil
	}

	// Check the validity of the ConnectionStability indicator
	if _, err := strconv.ParseFloat(data[4], 64); err != nil {
		return false, nil
	}

	// Check the validity of the TTFB indicator
	if _, err := strconv.Atoi(data[5]); err != nil {
		return false, nil
	}

	// Check the validity of the VoicePurity indicator
	if _, err := strconv.Atoi(data[6]); err != nil {
		return false, nil
	}

	// Check the validity of the MedianOfCallsTime indicator
	if _, err := strconv.Atoi(data[7]); err != nil {
		return false, nil
	}

	return true, nil
}

// GetStatus - function to get VoiceCall system status data from CSV file
func GetStatus(alphaCodes map[string]string) ([]VoiceCallData, error) {
	fmt.Println()
	fmt.Println("****************")
	fmt.Println("VoiceCall-Status:")

	store := make([]VoiceCallData, 0)

	data, err := ioutil.ReadFile("./assets/voice.data")
	if err != nil {
		return store, err
	}

	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		v := NewData()
		d := strings.Split(l, ";")

		if b, err := IsValid(d, alphaCodes); !b || err != nil {
			continue
		}

		v.SetCountry(d[0])

		bandwidth, _ := strconv.Atoi(d[1])
		v.SetBandwidth(bandwidth)

		responseTime, _ := strconv.Atoi(d[2])
		v.SetResponseTime(responseTime)

		v.SetProvider(d[3])

		connectionStability, _ := strconv.ParseFloat(d[4], 64)
		v.SetConnectionStability(connectionStability)

		ttfb, _ := strconv.Atoi(d[5])
		v.SetTTFB(ttfb)

		voicePurity, _ := strconv.Atoi(d[6])
		v.SetVoicePurity(voicePurity)

		median, _ := strconv.Atoi(d[7])
		v.SetMedianOfCallsTime(median)

		store = append(store, *v)
	}

	// Testing the function
	for _, v := range store {
		fmt.Printf("%v\n", v)
	}

	return store, nil
}
