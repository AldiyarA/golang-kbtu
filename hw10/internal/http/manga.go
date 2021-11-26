package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-redis/redis/v8"
	"hw8/internal/models"
	"hw8/internal/store"
	"log"
	"net/http"
	"strconv"
)

type MangaResource struct {
	store store.Store
	rdb *redis.Client
}

func NewMangaResource(store store.Store, rdb *redis.Client) *MangaResource {
	return &MangaResource{
		store: store,
		rdb: rdb,
	}
}

func (mr *MangaResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", mr.CreateManga)
	r.Get("/", mr.AllManga)
	r.Get("/{id}", mr.ByID)
	r.Put("/", mr.UpdateManga)
	r.Delete("/{id}", mr.DeleteManga)

	return r
}

func (mr *MangaResource) CreateManga(w http.ResponseWriter, r *http.Request) {
	title := new(models.Title)
	if err := json.NewDecoder(r.Body).Decode(title); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	fmt.Println(title)
	if err := mr.store.Manga().Create(r.Context(), title); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, "OK")
	mr.UpdateCaches()
}

func (mr *MangaResource) AllManga(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.TitleFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		titlesFromRDB, err := mr.rdb.Get(context.Background(), "manga#"+searchQuery).Result()
		fmt.Printf("rdb = %s\n", titlesFromRDB)
		if err == nil {
			titles := make([]*models.Title, 0)
			err := json.Unmarshal([]byte(titlesFromRDB), &titles)
			if err != nil {
				return
			}
			render.JSON(w, r, titles)
			return
		}
		filter.Query = &searchQuery
	}
	titles, err := mr.store.Manga().All(r.Context(), filter)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	if searchQuery != "" {
		fmt.Println(searchQuery)
		titlesMarshal, _ := json.Marshal(titles)
		mr.rdb.Set(context.Background(), "manga#"+searchQuery, titlesMarshal, 0)
	}
	render.JSON(w, r, titles)
}

func (mr *MangaResource) ByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	title, err := mr.store.Manga().ByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	render.JSON(w, r, title)
}


func (mr *MangaResource) UpdateManga(w http.ResponseWriter, r *http.Request) {
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
	if err := mr.store.Manga().Update(r.Context(), title); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
	render.JSON(w, r, "OK")
	mr.UpdateCaches()
}
func (mr *MangaResource) DeleteManga(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	if err := mr.store.Manga().Delete(r.Context(), id); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
	render.JSON(w, r, "OK")
	mr.UpdateCaches()
}

func (mr *MangaResource) UpdateCaches () {
	keys := mr.rdb.Keys(context.Background(), "manga#*").Val()
	for _, keyVal := range keys{
		key := keyVal[6:]
		fmt.Println(key)
		filter := &models.TitleFilter{Query: &key}
		titles, err := mr.store.Manga().All(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
		titlesMarshal, _ := json.Marshal(titles)
		mr.rdb.Set(context.Background(), keyVal, titlesMarshal, 0)
	}
}