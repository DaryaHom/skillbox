package data

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetFromAPI - reads data from simulator API.
// Writes data to an array of structures that's passed as result interface{}
func GetFromAPI(host, simulatorAddr, addr string, result interface{}) error {

	resp, err := http.Get(host + simulatorAddr + addr)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)

	if err != nil {
		return err
	}

	return nil
}
