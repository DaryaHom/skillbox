package city

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/jszwec/csvutil"
)

type City struct {
	ID         int    `csv`
	Name       string `csv:`
	Region     string `csv:`
	District   string `csv:`
	Population int    `csv:`
	Foundation int    `csv:`
}

func NewCity() *City {
	return &City{}
}

func (c *City) Id() int {
	return c.ID
}

func (c *City) SetID(id int) {
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

func (c *City) GetInfo() ([]byte, error) {
	res, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ReadInfo - reads csv-file to the array of city-structures
func ReadInfo(reader *csv.Reader) ([]*City, error) {

	// Creating header for file without it
	userHeader, err := csvutil.Header(NewCity(), "csv")
	if err != nil {
		return nil, err
	}

	dec, err := csvutil.NewDecoder(reader, userHeader...)
	if err != nil {
		return nil, err
	}

	// Decoding line to the structure & appending to the array
	var cities []*City
	for {
		var c *City
		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}
	return cities, nil
}

// WriteInfo - reads lines from an array and overwrites them to the file
func WriteInfo(file *os.File, c []City) error {
	mutex := &sync.Mutex{}
	writer := csv.NewWriter(file)

	var cities [][]string
	for _, city := range c {
		var row []string
		row = append(row, fmt.Sprintf("%v", city.ID))
		row = append(row, city.Name)
		row = append(row, city.Region)
		row = append(row, city.District)
		row = append(row, fmt.Sprintf("%v", city.Population))
		row = append(row, fmt.Sprintf("%v", city.Foundation))
		cities = append(cities, row)
	}

	mutex.Lock()
	_, err := file.Seek(0, io.SeekStart)
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
