package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"hw8/internal/models"
	"hw8/internal/store"
	"log"
	"net/http"
	"strconv"
	"time"
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
	r.Route(`/anime`, func(r chi.Router) {
		r.Get(`/`, func(w http.ResponseWriter, r *http.Request) {
			titles, err := s.store.Anime().All(r.Context())
			if err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			render.JSON(w, r, titles)
		})
		r.Get(`/{id}`, func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			title, err := s.store.Anime().ByID(r.Context(), id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			render.JSON(w, r, title)
		})
		r.Post(`/`, func(w http.ResponseWriter, r *http.Request) {
			title := new(models.Title)
			if err := json.NewDecoder(r.Body).Decode(title); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			fmt.Println(title)
			if err := s.store.Anime().Create(r.Context(), title); err!=nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			render.JSON(w, r, "OK")
		})
		r.Put(`/`, func(w http.ResponseWriter, r *http.Request) {
			title := new(models.Title)
			if err := json.NewDecoder(r.Body).Decode(title); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			fmt.Println(title)

			if err := validation.ValidateStruct(title, validation.Field(&title.Id, validation.Required)); err != nil{
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			if err := s.store.Anime().Update(r.Context(), title); err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			render.JSON(w, r, "OK")
		})
		r.Delete(`/{id}`, func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			if err := s.store.Anime().Delete(r.Context(), id); err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			render.JSON(w, r, "OK")
		})
	})
	r.Route(`/manga`, func(r chi.Router) {
		r.Get(`/`, func(w http.ResponseWriter, r *http.Request) {
			titles, err := s.store.Manga().All(r.Context())
			if err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			render.JSON(w, r, titles)
		})
		r.Get(`/{id}`, func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			title, err := s.store.Manga().ByID(r.Context(), id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			render.JSON(w, r, title)
		})
		r.Post(`/`, func(w http.ResponseWriter, r *http.Request) {
			title := new(models.Title)
			if err := json.NewDecoder(r.Body).Decode(title); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			fmt.Println(title)
			if err := s.store.Manga().Create(r.Context(), title); err!=nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			render.JSON(w, r, "OK")
		})
		r.Put(`/`, func(w http.ResponseWriter, r *http.Request) {
			title := new(models.Title)
			if err := json.NewDecoder(r.Body).Decode(title); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			fmt.Println(title)

			if err := validation.ValidateStruct(title, validation.Field(&title.Id, validation.Required)); err != nil{
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			if err := s.store.Manga().Update(r.Context(), title); err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			render.JSON(w, r, "OK")
		})
		r.Delete(`/{id}`, func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			if err := s.store.Manga().Delete(r.Context(), id); err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			render.JSON(w, r, "OK")
		})
	})
	r.Route(`/ranobe`, func(r chi.Router) {
		r.Get(`/`, func(w http.ResponseWriter, r *http.Request) {
			titles, err := s.store.Ranobe().All(r.Context())
			if err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			render.JSON(w, r, titles)
		})
		r.Get(`/{id}`, func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			title, err := s.store.Ranobe().ByID(r.Context(), id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			render.JSON(w, r, title)
		})
		r.Post(`/`, func(w http.ResponseWriter, r *http.Request) {
			title := new(models.Title)
			if err := json.NewDecoder(r.Body).Decode(title); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			fmt.Println(title)
			if err := s.store.Ranobe().Create(r.Context(), title); err!=nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			render.JSON(w, r, "OK")
		})
		r.Put(`/`, func(w http.ResponseWriter, r *http.Request) {
			title := new(models.Title)
			if err := json.NewDecoder(r.Body).Decode(title); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			fmt.Println(title)

			if err := validation.ValidateStruct(title, validation.Field(&title.Id, validation.Required)); err != nil{
				w.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			if err := s.store.Ranobe().Update(r.Context(), title); err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			render.JSON(w, r, "OK")
		})
		r.Delete(`/{id}`, func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Unknown err: %v", err)
				return
			}
			if err := s.store.Ranobe().Delete(r.Context(), id); err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB err: %v", err)
				return
			}
			render.JSON(w, r, "OK")
		})
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
