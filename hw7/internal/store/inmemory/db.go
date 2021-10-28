package inmemory

import (
	"hw7/api"
	"hw7/internal/store"
	"hw7/internal/store/inmemory/electronics"

	"sync"
)

type DB struct {
	users api.UserServiceServer
	electronicsRepo store.ElectronicsRepository
	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Electronics() store.ElectronicsRepository {
	if db.electronicsRepo == nil{
		db.electronicsRepo = electronics.NewElectronicsRepo()
	}
	return db.electronicsRepo
}
func (db *DB) Users() api.UserServiceServer {
	if db.users == nil{
		db.users = &UserRepo{
			data: make(map[int64]*api.User),
			mu: new(sync.RWMutex),
		}
	}
	return db.users
}