package electronics

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hw7/api"
	"sync"
)

type PhoneRepo struct {
	data map[int64]*api.Phone
	api.UnimplementedPhoneServiceServer
	mu *sync.RWMutex
}

func (c *PhoneRepo) All(ctx context.Context, req *api.Empty) (*api.Phones, error){
	c.mu.RLock()
	defer c.mu.RUnlock()
	phones := &api.Phones{
		Phones: make([]*api.Phone, 0, len(c.data)),
	}
	for _, computer := range c.data{
		phones.Phones = append(phones.Phones, computer)
	}
	return phones, nil
}
func (c *PhoneRepo) Get(ctx context.Context, req *api.Id) (*api.Phone, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	id := req.Id
	fmt.Println(id)
	phone, ok := c.data[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("phone with id %d does not exist", req.Id))
	}
	return phone, nil
}
func (c* PhoneRepo) Create(ctx context.Context, req *api.Phone) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[req.Id] = req
	return &api.Empty{}, nil
}

func (c *PhoneRepo) Update (ctx context.Context, req *api.Phone) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[req.Id] = req
	return &api.Empty{}, nil
}
func (c *PhoneRepo) Delete(ctx context.Context, req *api.Id) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, req.Id)
	return &api.Empty{}, nil
}
