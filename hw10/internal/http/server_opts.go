package http

import (
	lru "github.com/hashicorp/golang-lru"
	"hw8/internal/message_broker"
	"hw8/internal/store"
)

type ServerOption func(srv *Server)

func WithAddress(address string) ServerOption {
	return func(srv *Server) {
		srv.Address = address
	}
}

func WithStore(store store.Store) ServerOption {
	return func(srv *Server) {
		srv.store = store
	}
}

//func WithRedis(redis *redis.Client) ServerOption {
//	return func(srv *Server) {
//		srv.rdb = redis
//	}
//}

func WithCache(cache *lru.TwoQueueCache) ServerOption {
	return func(srv *Server) {
		srv.cache = cache
	}
}

func WithBroker(broker message_broker.MessageBroker) ServerOption {
	return func(srv *Server) {
		srv.broker = broker
	}
}