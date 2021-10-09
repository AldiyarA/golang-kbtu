package models

import (
	"sync"
	"time"
)

type Runable interface {
	Run()
}

type Factory struct {
	Name     string     `json:"name"`
	Out      int        `json:"result"`
	OutValue int        `json:"-"`
	Duration int        `json:"-"`
	Status   string     `json:"status"`
	mu       sync.Mutex `json:"-"`
}

type Factory2 struct {
	Factory `json:""`
	Name    string `json:"name"`
	In      []int  `json:"in"`
	NeedIn  []int  `json:"-"`
}

func (f *Factory) Run() {
	f.Status = "producing"
	for {
		time.Sleep(time.Duration(f.Duration) * time.Second)
		f.mu.Lock()
		f.Out += f.OutValue
		f.mu.Unlock()
	}
}

func (f *Factory2) Run() {
	f.Status = "waiting"
	n := len(f.In)
LOOP:
	for {
		for i := 0; i < n; i++ {
			if f.In[i] < f.NeedIn[i] {
				f.Status = "waiting"
				time.Sleep(1 * time.Second)
				goto LOOP
			}
		}
		f.mu.Lock()
		for i := 0; i < n; i++ {
			f.In[i] -= f.NeedIn[i]
		}
		f.mu.Unlock()
		f.Status = "producing"
		time.Sleep(time.Duration(f.Duration) * time.Second)
		f.mu.Lock()
		f.Out += f.OutValue
		f.mu.Unlock()
	}
}

type Worker struct {
	Duration int       `json:"-"`
	In       *Factory  `json:"-"`
	Out      *Factory2 `json:"-"`
	OutIndex int       `json:"-"`
	Power    int       `json:"-"`
	Taken    int       `json:"now dragging"`
	Status   string    `json:"status"`
}

func (w *Worker) Run() {
	w.Status = "waiting on " + w.In.Name
LOOP:
	for {
		if w.In.Out == 0 {
			w.Status = "waiting on " + w.In.Name
			time.Sleep(400 * time.Millisecond)
			goto LOOP
		}
		w.In.mu.Lock()
		if w.In.Out < w.Power {
			w.Taken = w.In.Out
			w.In.Out = 0
		} else {
			w.Taken = w.Power
			w.In.Out -= w.Power
		}
		w.In.mu.Unlock()
		w.Status = "On the way from " + w.In.Name + " to " + w.Out.Name
		time.Sleep(time.Duration(w.Duration) * time.Second)
		w.Out.mu.Lock()
		w.Out.In[w.OutIndex] += w.Taken
		w.Taken = 0
		w.Out.mu.Unlock()
		w.Status = "On the way from " + w.Out.Name + " to " + w.In.Name
		time.Sleep(time.Duration(w.Duration) * time.Second)
	}
}
