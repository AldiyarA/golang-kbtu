package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"hw8/internal/models"
	"hw8/internal/store"
	"net/http"
	"strconv"
)

type TitleHandler struct {
	Repository store.TitleRepository
}

func NewTitleHandler(repository store.TitleRepository) *TitleHandler {
	return &TitleHandler{Repository: repository}
}

func (t* TitleHandler) GenRoutes(r chi.Router){
	r.Get(`/`, func(w http.ResponseWriter, r *http.Request) {
		titles, err := t.Repository.All(r.Context())
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
		title, err := t.Repository.ByID(r.Context(), id)
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
		if err := t.Repository.Create(r.Context(), title); err!=nil{
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
		if err := t.Repository.Update(r.Context(), title); err != nil{
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
		if err := t.Repository.Delete(r.Context(), id); err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}
		render.JSON(w, r, "OK")
	})

}
