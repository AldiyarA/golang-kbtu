package store

import (
	"hw7/api"
)

type Store interface {
	Users() api.UserServiceServer
	Electronics() ElectronicsRepository
	Tools() ToolsRepository
}

type ElectronicsRepository interface {
	Computers() api.ComputerServiceServer
	Phones() api.PhoneServiceServer
}

type ToolsRepository interface {

}