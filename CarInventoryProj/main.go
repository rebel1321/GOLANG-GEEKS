package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Car struct {
	Id    int64
	Name  string
	Model string
	Brand string
	Year  int
	Price float64
}

var Cars = make(map[int64]Car)
var mu sync.Mutex

func carHandler( w http.ResponseWriter,r *http.Request) {
	// url := r.URL.String()
	path := r.URL.Path
	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			createCar(w, r)
		} else {
			http.Error(w, "Incorrect post request", http.StatusBadRequest)
		}
	case "GET":
		if entity == "" {
			http.Error(w,"We don't support this API",http.StatusBadRequest)
		} else {
			id, err := strconv.ParseInt(entity, 10, 64)
			if err != nil {
				http.Error(w, "Invalid car ID", http.StatusBadRequest)
				return
			}
			getCar(w, r, int64(id))
		}
	case "DELETE":
		if entity == "" {
			http.Error(w,"We don't support this API",http.StatusBadRequest)
		} else {
			id, err := strconv.ParseInt(entity, 10, 64)
			if err != nil {
				http.Error(w, "Invalid car ID", http.StatusBadRequest)
				return
			}
			deleteCar(w, r,id) // Assuming PUT is similar to POST for this example
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	
}
// {
// 		"id": 1,
		// "name": "Toyota",
		// "model": "Corolla",
		// "brand": "Toyota",
		// "year": 2020,
		// "price": 2000000.0,
// 	}
func createCar(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

id := rand.Intn(10000)
car.Id = int64(id)
Cars[car.Id] = car
fmt.Printf("Car created: %+v\n", car) 
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(car)
}
func getCar(w http.ResponseWriter, r *http.Request,id int64) {
	mu.Lock()
	defer mu.Unlock()

car,ok := Cars[id]
if !ok {
	http.Error(w, "Car not found", http.StatusNotFound)
	return
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(car)
}
func deleteCar(w http.ResponseWriter, r *http.Request,id int64) {
	mu.Lock()
	defer mu.Unlock()

_,ok := Cars[id]
if !ok {
	http.Error(w, "Car not found", http.StatusNotFound)
	return
}
delete(Cars,id)
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/cars", carHandler)
	http.HandleFunc("/cars/", carHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
