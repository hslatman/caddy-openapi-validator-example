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
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/go-chi/chi"

	api "github.com/deepmap/oapi-codegen/examples/petstore-expanded/chi/api"
)

func init() {
	caddy.RegisterModule(ExpandedPetStore{})
}

// ExpandedPetStore struct keeps module data
type ExpandedPetStore struct {
	handler http.Handler
}

// CaddyModule defines the ExpandedPetStore module
func (ExpandedPetStore) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.expanded_petstore_api_example",
		New: func() caddy.Module { return new(ExpandedPetStore) },
	}
}

// Provision sets up the Petstore API
func (p *ExpandedPetStore) Provision(ctx caddy.Context) error {

	// Setup an Expanded Pet Store implementation based on an implementation in
	// deepmap/oapi-codegen that's based on the Chi framework.
	ps := api.NewPetStore()
	handler := api.Handler(ps)

	r := chi.NewRouter()
	r.Mount("/api", handler)

	p.handler = r

	// Add some pets
	ps.NextId = 1
	var one api.Pet
	one.Name = "Pet One"
	one.Tag = nil
	one.Id = ps.NextId
	ps.NextId = ps.NextId + 1
	ps.Pets[one.Id] = one

	var two api.Pet
	two.Name = "Pet Two"
	two.Tag = nil
	two.Id = ps.NextId
	ps.NextId = ps.NextId + 1
	ps.Pets[two.Id] = two

	return nil

}

// ServeHTTP serves a simple (and currently incomplete) Pet Store API
func (p *ExpandedPetStore) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	// Set the default response content type
	w.Header().Set("Content-Type", "application/json")

	// Call the Chi Server(Interface) to execute the request
	p.handler.ServeHTTP(w, r)

	// Continue to the next handler in the Caddy stack (if it exists)
	return next.ServeHTTP(w, r)
}
