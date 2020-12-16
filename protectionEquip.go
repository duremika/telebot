package main

import "time"

type Equip struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}