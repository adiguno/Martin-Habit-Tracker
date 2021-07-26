package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Habit entity
type Habit struct {
	ID          uint64      `json:"id"`
	habitType   HabitType   `json:"habitType"`
	habitStatus HabitStatus `json:"habitStatus"`
	description string      `json:"description"`
	userID      uint64      `json:"userId"`
}

type HabitType int

const (
	Daily = iota
	Weekday
	Weekend
)

func (ht HabitType) String() string {
	return [...]string{"Daily", "Weekday", "Weekend"}[ht]
}

type HabitStatus int

const (
	Todo = iota
	InProgress
	Done
)

func (hs HabitStatus) String() string {
	return [...]string{"Todo", "In Progress", "Done"}[hs]
}

func homeLink(writer http.ResponseWriter, request *http.Request) {
	// func homeLink(writer http.ResponseWriter, request *http.Request) Habit {
	fmt.Println("get habits")
	habit1 := Habit{
		ID:          0,
		habitType:   HabitType(Daily),
		habitStatus: HabitStatus(Done),
		description: "test habit",
		userID:      1234,
	}
	fmt.Println(habit1)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// json.NewEncoder(writer).Encode(habit1)

	payload, err := json.Marshal(habit1)
	if err != nil {
		log.Println(err)
	}
	writer.Write(payload)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
