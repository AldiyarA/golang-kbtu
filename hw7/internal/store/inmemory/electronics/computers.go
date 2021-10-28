package electronics

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hw7/api"
	"sync"
)

type ComputerRepo struct {
	data map[int64]*api.Computer
	api.UnimplementedComputerServiceServer
	mu *sync.RWMutex
}

func (c *ComputerRepo) All(ctx context.Context, req *api.Empty) (*api.Computers, error){
	c.mu.RLock()
	defer c.mu.RUnlock()
	computers := &api.Computers{
		Computers: make([]*api.Computer, 0, len(c.data)),
	}
	for _, computer := range c.data{
		computers.Computers = append(computers.Computers, computer)
	}
	return computers, nil
}
func (c *ComputerRepo) Get(ctx context.Context, req *api.Id) (*api.Computer, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	id := req.Id
	fmt.Println(id)
	computer, ok := c.data[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("comruter with id %d does not exist", req.Id))
	}
	return computer, nil
}
func (c* ComputerRepo) Create(ctx context.Context, req *api.Computer) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("here")
	c.data[req.Id] = req
	return &api.Empty{}, nil
}

func (c *ComputerRepo) Update (ctx context.Context, req *api.Computer) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[req.Id] = req
	return &api.Empty{}, nil
}
func (c *ComputerRepo) Delete(ctx context.Context, req *api.Id) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, req.Id)
	return &api.Empty{}, nil
}
