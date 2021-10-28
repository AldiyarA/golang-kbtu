package electronics

import (
	"hw7/api"
	"hw7/internal/store"
	"sync"
)

type ElectronicsRepo struct {
	computers api.ComputerServiceServer
	phones api.PhoneServiceServer
	mu *sync.RWMutex
}

func NewElectronicsRepo() store.ElectronicsRepository {
	return &ElectronicsRepo{mu: new(sync.RWMutex)}
}

func (e *ElectronicsRepo) Computers() api.ComputerServiceServer {
	if e.computers == nil{
		e.computers = &ComputerRepo{
			data: make(map[int64]*api.Computer),
			mu: new(sync.RWMutex),
		}
	}
	return e.computers
}

func (e *ElectronicsRepo) Phones() api.PhoneServiceServer {
	if e.phones == nil{
		e.phones = &PhoneRepo{
			data: make(map[int64]*api.Phone),
			mu: new(sync.RWMutex),
		}
	}
	return e.phones
}
