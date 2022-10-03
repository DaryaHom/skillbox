package db

import (
	"encoding/csv"
	"interim_attestation/internal/city"
	"io"
	"os"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jszwec/csvutil"
)

type StoreDB interface {
	CreateCity(id string, name string, region string, district string, population int, foundation int) error
	ReadInfo(reader *csv.Reader) error
	GetInfo(id string) (*city.City, error)
	GetByRegion(regionName string) ([]string, error)
	GetByDistrict(districtName string) ([]string, error)
	GetByPopulation(minVal, maxVal int) ([]string, error)
	GetByFoundation(minVal, maxVal int) ([]string, error)
	UpdatePopulation(id string, population int) error
	DeleteCity(id string) error
	WriteInfo(file *os.File) error
}

type DB struct {
	Conn redis.Conn
}

// CreateCity - create a new key in Redis and writes the passed values,
// or overwrites them if the key already exists
func (rd *DB) CreateCity(id string, name string, region string, district string, population int, foundation int) error {
	_, err := rd.Conn.Do(
		"HMSET",
		id,
		"id",
		id,
		"name",
		name,
		"region",
		region,
		"district",
		district,
		"population",
		population,
		"foundation",
		foundation,
	)
	if err != nil {
		return err
	}
	return nil
}

// ReadInfo - reads csv-file to the Redis data store
func (rd *DB) ReadInfo(reader *csv.Reader) error {
	// Check if there is data in Redis
	keys, err := redis.Strings(rd.Conn.Do("KEYS", "*"))
	if err != nil {
		return err
	}

	// If there is data, stop reading so as not to lose the changes made by another service
	if len(keys) > 0 {
		return nil
	}

	// Create header for file without it
	header, err := csvutil.Header(city.NewCity(), "csv")
	if err != nil {
		return err
	}

	dec, err := csvutil.NewDecoder(reader, header...)
	if err != nil {
		return err
	}

	// Decode line to the city-structure & add structure fields to Redis
	for {
		var c *city.City
		if err = dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		err = rd.CreateCity(c.ID, c.Name, c.Region, c.District, c.Population, c.Foundation)
		if err != nil {
			return err
		}
	}
	return nil
}

//
//func (c *City) GetInfo() ([]byte, error) {
//	res, err := json.Marshal(c)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}

//GetInfo - returns from Redis info about city by id
func (rd *DB) GetInfo(id string) (*city.City, error) {
	var city city.City

	// Get all fields of a key that is equal to id
	reply, err := redis.Values(rd.Conn.Do("HGETALL", id))
	if err != nil {
		return &city, err
	}

	// Scan info from reply to the new City struct
	err = redis.ScanStruct(reply, &city)
	if err != nil {
		return &city, err
	}

	return &city, nil
}

// GetByRegion - returns from Redis info about cities by region
func (rd *DB) GetByRegion(regionName string) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(rd.Conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the region name matches the received name.
	// If the region name matches, get the city name & add it to returned slice
	for _, key := range keys {
		region, err := redis.String(rd.Conn.Do("HGET", key, "region"))
		if err != nil {
			return cities, err
		}
		if region == regionName {
			city, err := redis.String(rd.Conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// GetByDistrict - returns from Redis info about cities by district
func (rd *DB) GetByDistrict(districtName string) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(rd.Conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the district name matches the received name.
	// If the district name matches, get the city name & add it to returned slice
	for _, key := range keys {
		district, err := redis.String(rd.Conn.Do("HGET", key, "district"))
		if err != nil {
			return cities, err
		}
		if district == districtName {
			city, err := redis.String(rd.Conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// GetByPopulation - returns from Redis a list of names of all cities with the specified population range
func (rd *DB) GetByPopulation(minVal, maxVal int) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(rd.Conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the population is in the specified range.
	// If the population is in the range, get the city name & add it to returned slice
	for _, key := range keys {
		population, err := redis.Int(rd.Conn.Do("HGET", key, "population"))
		if err != nil {
			return cities, err
		}
		if population >= minVal && population <= maxVal {
			city, err := redis.String(rd.Conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// GetByFoundation - returns from Redis a list of names of all cities with the specified foundation range
func (rd *DB) GetByFoundation(minVal, maxVal int) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(rd.Conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the foundation is in the specified range.
	// If the foundation is in the range, get the city name & add it to returned slice
	for _, key := range keys {
		foundation, err := redis.Int(rd.Conn.Do("HGET", key, "foundation"))
		if err != nil {
			return cities, err
		}
		if foundation >= minVal && foundation <= maxVal {
			city, err := redis.String(rd.Conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// UpdatePopulation - sets new population value for the city with specified id
func (rd *DB) UpdatePopulation(id string, population int) error {
	_, err := rd.Conn.Do(
		"HSET",
		id,
		"population",
		population,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rd *DB) DeleteCity(id string) error {
	_, err := rd.Conn.Do("DEL", id)
	if err != nil {
		return err
	}
	return nil
}

// WriteInfo - writes changes to the csv-file before service stop
func (rd *DB) WriteInfo(file *os.File) error {
	mutex := &sync.Mutex{}
	writer := csv.NewWriter(file)

	keys, err := redis.Strings(rd.Conn.Do("KEYS", "*"))
	if err != nil {
		return err
	}

	var cities [][]string
	for _, key := range keys {
		city, err := redis.Strings(rd.Conn.Do("HVALS", key))
		if err != nil {
			return err
		}
		cities = append(cities, city)
	}

	mutex.Lock()
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	err = writer.WriteAll(cities)
	if err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}
