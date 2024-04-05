/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstoreserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// StoreAPIController binds http requests to an api service and writes the service results to the http response
type StoreAPIController struct {
	service StoreAPIServicer
	errorHandler ErrorHandler
}

// StoreAPIOption for how the controller is set up.
type StoreAPIOption func(*StoreAPIController)

// WithStoreAPIErrorHandler inject ErrorHandler into controller
func WithStoreAPIErrorHandler(h ErrorHandler) StoreAPIOption {
	return func(c *StoreAPIController) {
		c.errorHandler = h
	}
}

// NewStoreAPIController creates a default api controller
func NewStoreAPIController(s StoreAPIServicer, opts ...StoreAPIOption) Router {
	controller := &StoreAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the StoreAPIController
func (c *StoreAPIController) Routes() Routes {
	return Routes{
		"DeleteOrder": Route{
			strings.ToUpper("Delete"),
			"/v2/store/order/{orderId}",
			c.DeleteOrder,
		},
		"GetInventory": Route{
			strings.ToUpper("Get"),
			"/v2/store/inventory",
			c.GetInventory,
		},
		"GetOrderById": Route{
			strings.ToUpper("Get"),
			"/v2/store/order/{orderId}",
			c.GetOrderById,
		},
		"PlaceOrder": Route{
			strings.ToUpper("Post"),
			"/v2/store/order",
			c.PlaceOrder,
		},
	}
}

// DeleteOrder - Delete purchase order by ID
func (c *StoreAPIController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderIdParam := chi.URLParam(r, "orderId")
	if orderIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"orderId"}, nil)
		return
	}
	result, err := c.service.DeleteOrder(r.Context(), orderIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetInventory - Returns pet inventories by status
func (c *StoreAPIController) GetInventory(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetInventory(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetOrderById - Find purchase order by ID
func (c *StoreAPIController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderIdParam, err := parseNumericParameter[int64](
		chi.URLParam(r, "orderId"),
		WithRequire[int64](parseInt64),
		WithMinimum[int64](1),
		WithMaximum[int64](5),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetOrderById(r.Context(), orderIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PlaceOrder - Place an order for a pet
func (c *StoreAPIController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	orderParam := Order{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&orderParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOrderRequired(orderParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertOrderConstraints(orderParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PlaceOrder(r.Context(), orderParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
