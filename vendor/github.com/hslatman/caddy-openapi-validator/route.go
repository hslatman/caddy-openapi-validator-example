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

package openapi

import (
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
)

// validateRoute checks whether a route with the right properties (server, path, method) can be found
func (v *Validator) validateRoute(r *http.Request) (*openapi3filter.RequestValidationInput, *oapiError) {

	// Reconstruct the url from the request; makes it work for localhost
	url := r.URL
	url.Host = r.Host // TODO: verify this is an OK thing to do (i.e. what about proxies? Other protocols?)
	if r.TLS == nil {
		url.Scheme = "http"
	} else {
		url.Scheme = "https"
	}

	// TODO: option to cut off a prefix after host? I.e. /api, if that's part of the base url?

	method := r.Method
	route, pathParams, err := v.router.FindRoute(method, url)

	// No route found for the request
	if err != nil {
		switch e := err.(type) {
		case *openapi3filter.RouteError:
			// The requested path doesn't match the server, path or anything else.
			// TODO: switch between cases based on the e.Reason string? Some are not found, some are invalid method, etc.
			return nil, &oapiError{
				Code:    http.StatusNotFound, //http.StatusBadRequest,
				Message: e.Reason,
			}
		default:
			// Fallback for unexpected or unimplemented cases
			return nil, &oapiError{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("error validating route: %s", err.Error()),
			}
		}
	}

	validationInput := &openapi3filter.RequestValidationInput{
		Request:     r,
		PathParams:  pathParams,
		Route:       route,
		QueryParams: r.URL.Query(),
		// QueryParams  url.Values
	}

	if v.options != nil {
		validationInput.Options = &v.options.Options
		validationInput.ParamDecoder = v.options.ParamDecoder
	}

	return validationInput, nil
}
