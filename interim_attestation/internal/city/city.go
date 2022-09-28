package city

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jszwec/csvutil"
)

type City struct {
	ID         string `redis:"id"`
	Name       string `redis:"name"`
	Region     string `redis:"region"`
	District   string `redis:"district"`
	Population int    `redis:"population"`
	Foundation int    `redis:"foundation"`
}

func NewCity() *City {
	return &City{}
}

func (c *City) Id() string {
	return c.ID
}

func (c *City) SetID(id string) {
	c.ID = id
}

func (c *City) GetName() string {
	return c.Name
}

func (c *City) SetName(name string) {
	c.Name = name
}

func (c *City) GetRegion() string {
	return c.Region
}

func (c *City) SetRegion(region string) {
	c.Region = region
}

func (c *City) GetDistrict() string {
	return c.District
}

func (c *City) SetDistrict(district string) {
	c.District = district
}

func (c *City) GetPopulation() int {
	return c.Population
}

func (c *City) SetPopulation(population int) {
	c.Population = population
}

func (c *City) GetFoundation() int {
	return c.Foundation
}

func (c *City) SetFoundation(foundation int) {
	c.Foundation = foundation
}

// CreateCity - create a new key in Redis and writes the passed values,
// or overwrites them if the key already exists
func CreateCity(conn redis.Conn, id string, name string, region string, district string, population int, foundation int) error {
	_, err := conn.Do(
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
func ReadInfo(conn redis.Conn, reader *csv.Reader) error {
	// Check if there is data in Redis
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		return err
	}

	// If there is data, stop reading so as not to lose the changes made by another service
	if len(keys) > 0 {
		return nil
	}

	// Create header for file without it
	header, err := csvutil.Header(NewCity(), "csv")
	if err != nil {
		return err
	}

	dec, err := csvutil.NewDecoder(reader, header...)
	if err != nil {
		return err
	}

	// Decode line to the city-structure & add structure fields to Redis
	for {
		var c *City
		if err = dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		err = CreateCity(conn, c.ID, c.Name, c.Region, c.District, c.Population, c.Foundation)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *City) GetInfo() ([]byte, error) {
	res, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//GetInfo - returns from Redis info about city by id
func GetInfo(conn redis.Conn, id string) (*City, error) {
	var city City

	// Get all fields of a key that is equal to id
	reply, err := redis.Values(conn.Do("HGETALL", id))
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
func GetByRegion(conn redis.Conn, regionName string) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the region name matches the received name.
	// If the region name matches, get the city name & add it to returned slice
	for _, key := range keys {
		region, err := redis.String(conn.Do("HGET", key, "region"))
		if err != nil {
			return cities, err
		}
		if region == regionName {
			city, err := redis.String(conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// GetByDistrict - returns from Redis info about cities by district
func GetByDistrict(conn redis.Conn, districtName string) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the district name matches the received name.
	// If the district name matches, get the city name & add it to returned slice
	for _, key := range keys {
		district, err := redis.String(conn.Do("HGET", key, "district"))
		if err != nil {
			return cities, err
		}
		if district == districtName {
			city, err := redis.String(conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// GetByPopulation - returns from Redis a list of names of all cities with the specified population range
func GetByPopulation(conn redis.Conn, minVal, maxVal int) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the population is in the specified range.
	// If the population is in the range, get the city name & add it to returned slice
	for _, key := range keys {
		population, err := redis.Int(conn.Do("HGET", key, "population"))
		if err != nil {
			return cities, err
		}
		if population >= minVal && population <= maxVal {
			city, err := redis.String(conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// GetByFoundation - returns from Redis a list of names of all cities with the specified foundation range
func GetByFoundation(conn redis.Conn, minVal, maxVal int) ([]string, error) {
	var cities []string

	// Get a list of keys from Redis
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		return cities, err
	}

	// For each key, check if the foundation is in the specified range.
	// If the foundation is in the range, get the city name & add it to returned slice
	for _, key := range keys {
		foundation, err := redis.Int(conn.Do("HGET", key, "foundation"))
		if err != nil {
			return cities, err
		}
		if foundation >= minVal && foundation <= maxVal {
			city, err := redis.String(conn.Do("HGET", key, "name"))
			if err != nil {
				return cities, err
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}

// UpdatePopulation - sets new population value for the city with specified id
func UpdatePopulation(conn redis.Conn, id string, population int) error {
	_, err := conn.Do(
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

func DeleteCity(conn redis.Conn, id string) error {
	_, err := conn.Do("DEL", id)
	if err != nil {
		return err
	}
	return nil
}

// WriteInfo - writes changes to the csv-file before service stop
func WriteInfo(conn redis.Conn, file *os.File) error {
	mutex := &sync.Mutex{}
	writer := csv.NewWriter(file)

	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		return err
	}

	var cities [][]string
	for _, key := range keys {
		city, err := redis.Strings(conn.Do("HVALS", key))
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
