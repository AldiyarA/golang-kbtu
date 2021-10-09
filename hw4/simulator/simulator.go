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

var workers []models.Runable
var factories []models.Runable

var (
	f1 = NewFactory("Ore mine", 1, 3)
	f2 = NewFactory2("Smelting factory", 1, 3, []int{2})
	f3 = NewFactory("Wood factory", 1, 2)
	f4 = NewFactory2("Tool Factory", 1, 5, []int{3, 2})
	w1 = NewWorker(5, 0, 3, &f1, &f2)
	w2 = NewWorker(5, 0, 3, &f2.Factory, &f4)
	w3 = NewWorker(5, 1, 3, &f3, &f4)
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAll)
	http.HandleFunc("/workers", returnWorkers)
	http.HandleFunc("/factories", returnFactories)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func returnWorkers(w http.ResponseWriter, r *http.Request) {
	for _, r := range workers {
		json.NewEncoder(w).Encode(r)
	}
}
func returnFactories(w http.ResponseWriter, r *http.Request) {
	for _, r := range factories {
		json.NewEncoder(w).Encode(r)
	}
}
func returnAll(w http.ResponseWriter, r *http.Request) {
	returnWorkers(w, r)
	returnFactories(w, r)
}
func runAll() {
	for _, r := range workers {
		go r.Run()
	}
	for _, r := range factories {
		go r.Run()
	}
}
func Simulate() {
	go handleRequests()
	workers = []models.Runable{&w1, &w2, &w3}
	factories = []models.Runable{&f1, &f2, &f3, &f4}
	runAll()
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
