package http

import (
	"context"
	"encoding/json"
	"fmt"
	"hw7/api"
	"hw7/internal/store"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store

	Address string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
		store:       store,

		Address: address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	r.Post("/electronics/computers", func(w http.ResponseWriter, r *http.Request) {
		computer := new(api.Computer)
		if err := json.NewDecoder(r.Body).Decode(computer); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Electronics().Computers().Create(r.Context(), computer)
		render.JSON(w, r, computer)
	})
	r.Put("/electronics/computers", func(w http.ResponseWriter, r *http.Request) {
		computer := new(api.Computer)
		if err := json.NewDecoder(r.Body).Decode(computer); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Electronics().Computers().Update(r.Context(), computer)
	})
	r.Get("/electronics/computers", func(w http.ResponseWriter, r *http.Request) {
		computers, err := s.store.Electronics().Computers().All(r.Context(), &api.Empty{})
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, computers.Computers)
	})
	r.Get("/electronics/computers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		computer, err := s.store.Electronics().Computers().Get(r.Context(), &api.Id{Id: int64(id)})
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, computer)
	})
	r.Delete("/electronics/computers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Electronics().Computers().Delete(r.Context(), &api.Id{Id: int64(id)})
	})
	// PHONE
	r.Post("/electronics/phones", func(w http.ResponseWriter, r *http.Request) {
		phone := new(api.Phone)
		if err := json.NewDecoder(r.Body).Decode(phone); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Electronics().Phones().Create(r.Context(), phone)
	})
	r.Put("/electronics/phones", func(w http.ResponseWriter, r *http.Request) {
		phone := new(api.Phone)
		if err := json.NewDecoder(r.Body).Decode(phone); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Electronics().Phones().Update(r.Context(), phone)
	})
	r.Get("/electronics/phones", func(w http.ResponseWriter, r *http.Request) {
		phones, err := s.store.Electronics().Phones().All(r.Context(), &api.Empty{})
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, phones.Phones)
	})
	r.Get("/electronics/phones/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		phone, err := s.store.Electronics().Phones().Get(r.Context(), &api.Id{Id: int64(id)})
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, phone)
	})
	r.Delete("/electronics/phones/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Electronics().Phones().Delete(r.Context(), &api.Id{Id: int64(id)})
	})

	// USER

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		user := new(api.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Users().Create(r.Context(), user)
	})

	r.Put("/users", func(w http.ResponseWriter, r *http.Request) {
		user := new(api.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Users().Update(r.Context(), user)
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := s.store.Users().All(r.Context(), &api.Empty{})
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, users.Users)
	})

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		user, err := s.store.Users().Get(r.Context(), &api.Id{Id: int64(id)})
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, user)
	})

	r.Delete("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Users().Delete(r.Context(), &api.Id{Id: int64(id)})
	})

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() // блокируемся, пока контекст приложения не отменен

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// блок до записи или закрытия канала
	<-s.idleConnsCh
}
