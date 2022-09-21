package handler

import (
	"encoding/json"
	"fmt"
	"interim_attestation/internal/req"
	"interim_attestation/internal/storage"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
)

// Get - a hello function
func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Welcome")); err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
	}
}

// GetCity - returns information about the city by its id
func GetCity(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, "id")
		id, err := strconv.Atoi(val)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		res, err := s.GetCityInfo(id)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		req.WriteOkResponse(w, res)
	}
}

//CreateCity - creates new city in storage & fill it with request info
func CreateCity(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]interface{}{}

		// Decoding JSON fields to the map
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		defer req.CloseRequestBody(r.Body)

		// Casting values to the types corresponding to the fields of the city-structure
		id, err := strconv.Atoi(fmt.Sprintf("%v", request["id"]))
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		name := fmt.Sprintf("%v", request["name"])
		region := fmt.Sprintf("%v", request["region"])
		district := fmt.Sprintf("%v", request["district"])
		population, err := strconv.Atoi(fmt.Sprintf("%v", request["population"]))
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		foundation, err := strconv.Atoi(fmt.Sprintf("%v", request["foundation"]))
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}

		// New structure with received values
		s.CreateCity(id, name, region, district, population, foundation)
		response := []byte(fmt.Sprintf("New city %s is created\n", name))
		req.WriteCreateResponse(w, response)
	}
}

//UpdatePopulation - updates city population in the storage
func UpdatePopulation(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, "id")
		if val != "" {
			id, err := strconv.Atoi(val)
			if err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}

			request := map[string]interface{}{}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}
			defer req.CloseRequestBody(r.Body)

			population, err := strconv.Atoi(fmt.Sprintf("%v", request["new population"]))
			if err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}

			s.UpdatePopulation(id, population)
			response := []byte(fmt.Sprintf("Population of city %d has been updated to %d\n", id, population))
			req.WriteOkResponse(w, response)
			return
		}
		req.WriteBadRequest(w, []byte(fmt.Sprintf("Input id")))
	}
}

// GetByRegion - returns list of all cities in the region
func GetByRegion(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		res := []byte(fmt.Sprintf("%v", s.GetByRegion(name)))
		req.WriteOkResponse(w, res)
	}
}

// GetByDistrict - returns list of all cities in the district
func GetByDistrict(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		res := []byte(fmt.Sprintf("%v", s.GetByDistrict(name)))
		req.WriteOkResponse(w, res)
	}
}

// GetByPopulation - returns list of all cities with the specified population range
func GetByPopulation(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(chi.URLParam(r, "range"), "-")

		// If the minimum value is not specified in the url, for example "/-3000", then it sets to 0
		minVal := 0
		var err error
		if params[0] != "" {
			minVal, err = strconv.Atoi(params[0])
			if err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}
		}

		// If the maximum value is not specified in the url, for example "/3000-", then it sets to maximum int value
		maxVal := math.MaxInt
		if len(params) > 1 && params[1] != "" {
			maxVal, err = strconv.Atoi(params[1])
			if err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}
		}
		res := []byte(fmt.Sprintf("%v", s.GetByPopulation(minVal, maxVal)))
		req.WriteOkResponse(w, res)
	}
}

// GetByFoundation - returns list of all cities with the specified foundation range
func GetByFoundation(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(chi.URLParam(r, "range"), "-")

		//if the minimum value is not specified in the url, for example "/-1997", then it is set to 0
		minVal := 0
		var err error
		if params[0] != "" {
			minVal, err = strconv.Atoi(params[0])
			if err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}
		}

		//if the maximum value is not specified in the url, for example "/1997-", then it is set to this year
		maxVal := time.Now().Year()
		if len(params) > 1 && params[1] != "" {
			maxVal, err = strconv.Atoi(params[1])
			if err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}
		}

		res := []byte(fmt.Sprintf("%v", s.GetByFoundation(minVal, maxVal)))
		req.WriteOkResponse(w, res)
	}
}

// DeleteCity - removes city with the specified id from the storage
func DeleteCity(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		defer req.CloseRequestBody(r.Body)

		id, err := strconv.Atoi(request["target_id"])
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}

		s.DeleteCity(id)
		response := []byte(fmt.Sprintf("City %d is deleted\n", id))
		req.WriteOkResponse(w, response)
	}
}
