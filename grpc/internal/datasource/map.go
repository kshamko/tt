package datasource

import (
	"context"
	"errors"
	"sync"
)

type Data struct {
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []Coordinate
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}

type Coordinate struct {
	Deg int32
	Min int32
}

// ErrRoutes error.
var ErrNotFound = errors.New("data not found")

type Map struct {
	data map[string]Data
	mu   sync.RWMutex
}

func NewMap() *Map {
	return &Map{
		data: map[string]Data{},
	}
}

func (m *Map) GetEntry(ctx context.Context, id string) (Data, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if data, ok := m.data[id]; ok {
		return data, nil
	}

	return Data{}, ErrNotFound
}

func (m *Map) AddEntry(ctx context.Context, id string, data Data) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	m.data[id] = data
	return nil
}

func (m *Map) Status() (interface{}, error) {
	return "ok", nil
}
