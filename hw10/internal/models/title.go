package models

import (
	"time"
)

type Title struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	NameEnglish string `json:"name_english" db:"name_english"`
	Release time.Time `json:"release" db:"release"`
	Final  time.Time `json:"final" db:"final"`
	Status int       `json:"status" db:"status"`
	Type int       `json:"type" db:"type"`
}

type TitleFilter struct {
	Query *string `json:"query"`
}