package store

import (
	"hw7/api"
)

type Store interface {
	Users() api.UserServiceServer
	Electronics() ElectronicsRepository
}

type ElectronicsRepository interface {
	Computers() api.ComputerServiceServer
	Phones() api.PhoneServiceServer
}