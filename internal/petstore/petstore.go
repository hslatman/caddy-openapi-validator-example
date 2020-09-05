// Copyright 2020 Herman Slatman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package petstore

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(PetStore{})
}

// PetStore struct keeping module data
type PetStore struct {
	router    *mux.Router
	pets      map[int]*pet
	currentID int
}

type pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag,omitempty"`
	//Additional string `json:"additional"`
}

// CaddyModule defines the PetStore module
func (PetStore) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.petstore_api_example",
		New: func() caddy.Module { return new(PetStore) },
	}
}

// Provision sets up the Petstore API
func (p *PetStore) Provision(ctx caddy.Context) error {

	p.router = mux.NewRouter()
	api := p.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/pets", p.getPetsHandler).Methods(http.MethodGet)
	api.HandleFunc("/pets", p.postPetsHandler).Methods(http.MethodPost)
	api.HandleFunc("/pets/{id}", p.getPetHandler).Methods(http.MethodGet)

	p.currentID = 1

	p.pets = make(map[int]*pet)
	p.pets[p.currentID] = &pet{
		ID:   p.currentID,
		Name: "Pet 1",
	}

	return nil
}

func (p *PetStore) getPetsHandler(w http.ResponseWriter, r *http.Request) {

	pets := []pet{}
	for _, v := range p.pets {
		pets = append(pets, *v)
	}

	json.NewEncoder(w).Encode(pets)

	w.WriteHeader(http.StatusOK)
}

func (p *PetStore) postPetsHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var t pet
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.currentID = p.currentID + 1
	t.ID = p.currentID
	p.pets[p.currentID] = &t

	w.WriteHeader(http.StatusCreated)
}

func (p *PetStore) getPetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	iid, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // TODO: adapt petstore.yaml to allow 404s?
		return
	}

	pet, ok := p.pets[iid]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(pet)
	w.WriteHeader(http.StatusOK)
}

// ServeHTTP serves a simple (and currently incomplete) Pet Store API
func (p *PetStore) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	// Set the default response content type
	w.Header().Set("Content-Type", "application/json")

	// Call the Gorilla Mux ServeHTTP to match and execute a route
	p.router.ServeHTTP(w, r)

	// Continue to the next handler in the Caddy stack (if it exists)
	return next.ServeHTTP(w, r)
}
