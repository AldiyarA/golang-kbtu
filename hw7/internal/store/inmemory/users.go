package inmemory

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hw7/api"
	"sync"
)

type UserRepo struct {
	data map[int64]*api.User
	api.UnimplementedUserServiceServer
	mu *sync.RWMutex
}

func (c *UserRepo) All(ctx context.Context, req *api.Empty) (*api.Users, error){
	c.mu.RLock()
	defer c.mu.RUnlock()
	users := &api.Users{
		Users: make([]*api.User, 0, len(c.data)),
	}
	for _, user := range c.data{
		users.Users = append(users.Users, user)
	}
	return users, nil
}
func (c *UserRepo) Get(ctx context.Context, req *api.Id) (*api.User, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	id := req.Id
	fmt.Println(id)
	user, ok := c.data[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("user with id %d does not exist", req.Id))
	}
	return user, nil
}
func (c* UserRepo) Create(ctx context.Context, req *api.User) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[req.Id] = req
	return &api.Empty{}, nil
}

func (c *UserRepo) Update (ctx context.Context, req *api.User) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[req.Id] = req
	return &api.Empty{}, nil
}
func (c *UserRepo) Delete(ctx context.Context, req *api.Id) (*api.Empty, error){
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, req.Id)
	return &api.Empty{}, nil
}

