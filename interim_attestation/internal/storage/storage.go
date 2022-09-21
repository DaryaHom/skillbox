package storage

import (
	"encoding/csv"
	"interim_attestation/internal/city"
	"os"
)

type Storage struct {
	store map[int]*city.City
}

func NewStorage() *Storage {
	return &Storage{
		make(map[int]*city.City),
	}
}

func (s *Storage) ReadCitiesInfo(reader *csv.Reader) error {
	cities, err := city.ReadInfo(reader)
	if err != nil {
		return err
	}
	for _, v := range cities {
		s.store[v.Id()] = v
	}
	return nil
}

func (s *Storage) WriteCitiesInfo(file *os.File) error {
	cities := make([]city.City, len(s.store), len(s.store))
	counter := 0
	for _, c := range s.store {
		cities[counter] = *c
		counter++
	}

	err := city.WriteInfo(file, cities)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetCityInfo(id int) ([]byte, error) {
	res, err := s.store[id].GetInfo()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Storage) CreateCity(id int, name string, region string, district string, population int, foundation int) {
	c := city.NewCity()
	c.SetID(id)
	c.SetName(name)
	c.SetRegion(region)
	c.SetDistrict(district)
	c.SetPopulation(population)
	c.SetFoundation(foundation)
	s.store[id] = c
}

func (s *Storage) UpdatePopulation(id int, population int) {
	s.store[id].SetPopulation(population)
}

func (s *Storage) GetByRegion(name string) []string {
	var cities []string
	for _, v := range s.store {
		if v.GetRegion() == name {
			cities = append(cities, v.GetName())
		}
	}
	return cities
}

func (s *Storage) GetByDistrict(name string) []string {
	var cities []string
	for _, v := range s.store {
		if v.GetDistrict() == name {
			cities = append(cities, v.GetName())
		}
	}
	return cities
}

func (s *Storage) GetByPopulation(minVal, maxVal int) []string {
	var cities []string
	for _, v := range s.store {
		population := v.GetPopulation()
		if population >= minVal && population <= maxVal {
			cities = append(cities, v.GetName())
		}
	}
	return cities
}

func (s *Storage) GetByFoundation(minVal, maxVal int) []string {
	var cities []string
	for _, v := range s.store {
		foundation := v.GetFoundation()
		if foundation >= minVal && foundation <= maxVal {
			cities = append(cities, v.GetName())
		}
	}
	return cities
}

func (s *Storage) DeleteCity(id int) {
	delete(s.store, id)
}
