package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Car struct {
	ID     int    `json:"id"`
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Status string `json:"status"` //status is to check the car is  Available, In Service, or Rented
}

var cars []Car
var nextID int

// GET all cars handler
func getCarsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// GET car by ID handler
func getCarByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Getting ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	for _, car := range cars {
		if car.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// POST handler for creating a new car
func createCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var newCar Car
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// analyze the structure of the JSON
	err = json.Unmarshal(body, &newCar)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	// Set new ID and adding the car
	newCar.ID = nextID
	nextID++
	cars = append(cars, newCar)
	// return back with the created car
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCar)
}

// PUT handler is for updating an existing car
func updateCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// getting ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var updatedCar Car
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// analyze the structure of the JSON
	err = json.Unmarshal(body, &updatedCar)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	for i, car := range cars {
		if car.ID == id {
			cars[i].Make = updatedCar.Make
			cars[i].Model = updatedCar.Model
			cars[i].Year = updatedCar.Year
			cars[i].Status = updatedCar.Status
			// It will respond with the updated car
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cars[i])
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// DELETE handler for deleting a car by ID
func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// getting ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	for i, car := range cars {
		if car.ID == id {
			cars = append(cars[:i], cars[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func main() {
	// calling the main function for CRUD operations
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createCarHandler(w, r)
		} else if r.Method == http.MethodGet {
			getCarsHandler(w, r)
		}
	})

	http.HandleFunc("/cars/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCarByIDHandler(w, r)
		} else if r.Method == http.MethodPut {
			updateCarHandler(w, r)
		} else if r.Method == http.MethodDelete {
			deleteCarHandler(w, r)
		}
	})

	fmt.Println("Server running on port: 4455") //server is running on port 4455
	log.Fatal(http.ListenAndServe(":4455", nil))  //starts HTTP server on port 4455 if any error it terminate program
}
