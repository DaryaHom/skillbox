package test

import "testing"

type MockStoreDB struct{}

func (m *MockStoreDB) CreateCity(id string, name string, region string, district string, population int, foundation int) error {
	return nil
}

func (m *MockStoreDB) CountSales() (int, error) {
	return 333, nil
}

func TestCalculateSalesRate(t *testing.T) {
	// Инициализируем заглушку.
	m := &MockStoreDB{}
	// Передаём заглушку в функцию calculateSalesRate().
	sr := m.

	// Проверяем, соответствует ли возвращаемое значение ожиданиям на основе
	// фальшивых входных данных.
	exp := "0.33"
	if sr != exp {
		t.Fatalf("got %v; expected %v", sr, exp)
	}
}