package handler

import (
	"encoding/json"
	"fmt"
	"interim_attestation/internal/db"
	"interim_attestation/internal/req"
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
func GetCity(redisDB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		res, err := redisDB.GetInfo(id)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}

		req.WriteOkResponse(w, []byte(fmt.Sprintf("%v", *res)))
	}
}

//CreateCity - creates new city in storage & fill it with request info
func CreateCity(redisDB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		request := map[string]interface{}{}

		// Decode JSON fields to the map
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		defer req.CloseRequestBody(r.Body)

		// Caste values to the types corresponding to the fields of the city-structure
		id := fmt.Sprintf("%v", request["id"])
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

		// Create new structure with received values
		err = redisDB.CreateCity(id, name, region, district, population, foundation)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		response := []byte(fmt.Sprintf("New city %s is created\n", name))
		req.WriteCreateResponse(w, response)
	}
}

//UpdatePopulation - updates city population in the database
func UpdatePopulation(redisDB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, "id")
		if val != "" {
			id := val

			// Decode JSON fields to the map
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

			// Update the field 'population' with new value
			err = redisDB.UpdatePopulation(id, population)
			if err != nil {
				req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
				return
			}
			response := []byte(fmt.Sprintf("Population of city %s has been updated to %d\n", id, population))
			req.WriteOkResponse(w, response)
			return
		}
		req.WriteBadRequest(w, []byte(fmt.Sprintf("Input id")))
	}
}

// GetByRegion - returns a list of all cities in the region
func GetByRegion(redisDB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		regionName := chi.URLParam(r, "name")
		cities, err := redisDB.GetByRegion(regionName)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		res := []byte(fmt.Sprintf("%v", cities))
		req.WriteOkResponse(w, res)
	}
}

// GetByDistrict - returns a list of all cities in the district
func GetByDistrict(redisDB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		districtName := chi.URLParam(r, "name")
		cities, err := redisDB.GetByDistrict(districtName)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		res := []byte(fmt.Sprintf("%v", cities))
		req.WriteOkResponse(w, res)
	}
}

// GetByPopulation - returns a list of all cities with the specified population range
func GetByPopulation(redisDB *db.DB) http.HandlerFunc {
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

		cities, err := redisDB.GetByPopulation(minVal, maxVal)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		res := []byte(fmt.Sprintf("%v", cities))
		req.WriteOkResponse(w, res)
	}
}

// GetByFoundation - returns a list of all cities with the specified foundation range
func GetByFoundation(redisDB *db.DB) http.HandlerFunc {
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

		cities, err := redisDB.GetByFoundation(minVal, maxVal)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		res := []byte(fmt.Sprintf("%v", cities))
		req.WriteOkResponse(w, res)
	}
}

// DeleteCity - removes city with the specified id from Redis
func DeleteCity(redisDB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		defer req.CloseRequestBody(r.Body)

		id := request["target_id"]

		err := redisDB.DeleteCity(id)
		if err != nil {
			req.WriteBadRequest(w, []byte(fmt.Sprintf("%v", err)))
			return
		}
		response := []byte(fmt.Sprintf("City %s is deleted\n", id))
		req.WriteOkResponse(w, response)
	}
}
