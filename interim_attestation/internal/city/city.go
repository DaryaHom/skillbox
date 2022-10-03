package city

import (
	"encoding/csv"

	"github.com/gomodule/redigo/redis"
)

type Getter interface {
	GetAll() []City
}

type Adder interface {
	Add(City)
}

type Reader interface {
	ReadInfo(conn *redis.Conn, reader *csv.Reader) error
}

type Repo struct {
	City []City
}

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
