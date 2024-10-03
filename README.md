
# Car Management CRUD API

This is a simple  API built with Go programming for managing a collection of cars. 
The API allows users to perform basic  **CRUD (Create, Read, Update, Delete)** operations on car data. Each car has fields like `ID`, `Make`, `Model`, `Year`, and `Status` which indicates if the car is "Available", "In Service", or "Rented". The server listens on port **4455**.

## Features

- **Create** a new car: Adds a new car to the collection by sending a `POST` request.
- **Read** all cars or a specific car by ID: Retrieves the list of all cars or a particular car using `GET`.
- **Update** car information: Modifies the details of an existing car using a `PUT` request.
- **Delete** a car: Removes a car from the collection by sending a `DELETE` request.

## Endpoints

- `GET /cars`: Retrive all cars.
- `GET /cars/{id}`: retrive a specific car by its ID.
- `POST /cars`: Creates a new car.
- `PUT /cars/{id}`: Updates the details of a car by ID.
- `DELETE /cars/{id}`: Deletes a car by ID.

## Running the API
   go run main.go
