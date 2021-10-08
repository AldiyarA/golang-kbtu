package simulator

import (
	"context"
	"encoding/json"
	"fmt"
	"hw4/models"
	"log"
	"net/http"
	"time"
)

func NewFactory(name string, outValue, duration int) models.Factory {
	return models.Factory{Name: name, OutValue: outValue, Duration: duration}
}

func NewFactory2(name string, outValue, duration int, needIn []int) models.Factory2 {
	return models.Factory2{Factory: NewFactory(name, outValue, duration), Name: name, In: make([]int, len(needIn)), NeedIn: needIn}
}
func NewWorker(duration, outIndex, power int, in *models.Factory, out *models.Factory2) models.Worker {
	return models.Worker{Duration: duration, In: in, Out: out, OutIndex: outIndex, Power: power}
}

type All struct {
	Factories  []models.Factory  `json:"production"`
	Factories2 []models.Factory2 `json:"usage"`
	Workers    []models.Worker   `json:"workers"`
}

var all All

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func returnAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(all)
}
func Simulate() {
	f1 := NewFactory("Ore mine", 1, 3)
	f2 := NewFactory2("Smelting factory", 1, 3, []int{2})
	f3 := NewFactory("Wood factory", 1, 2)
	f4 := NewFactory2("Tool Factory", 1, 5, []int{3, 2})
	w1 := NewWorker(5, 0, 3, &f1, &f2)
	w2 := NewWorker(5, 0, 3, &f2.Factory, &f4)
	w3 := NewWorker(5, 1, 3, &f3, &f4)
	go handleRequests()
	go f1.Run()
	go f2.Run()
	go f3.Run()
	go f4.Run()
	go w1.Run()
	go w2.Run()
	go w3.Run()
	all = All{Factories: []models.Factory{f1, f2.Factory, f3, f4.Factory}, Factories2: []models.Factory2{f2, f4}, Workers: []models.Worker{w1, w2, w3}}
	tick := time.NewTicker(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			select {
			case <-tick.C:
				fmt.Println("Factory 1: ", f1.Name, ",out:", f1.Out)
				fmt.Println("Factory 3: ", f3.Name, ",out:", f1.Out)
				fmt.Println("Factory 2: ", f2.Name, ",in:", f2.In, ",out:", f2.Out)
				fmt.Println("Factory 4: ", f4.Name, ",in:", f4.In, ",out:", f4.Out)
				fmt.Println("Worker 1: Now dragging", w1.Taken, w1.Status)
				fmt.Println("Worker 2: Now dragging", w2.Taken, w2.Status)
				fmt.Println("Worker 3: Now dragging", w3.Taken, w3.Status)
				fmt.Println()
			case <-ctx.Done():
				return
			}
		}
	}()
	fmt.Scanln()
}
