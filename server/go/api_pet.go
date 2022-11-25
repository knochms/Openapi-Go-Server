/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [https://swagger.io](https://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  _If you're looking for the Swagger 2.0/OAS 2.0 version of Petstore, then click [here](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml). Alternatively, you can load via the `Edit > Load Petstore OAS 2.0` menu option!_  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// PetApiController binds http requests to an api service and writes the service results to the http response
type PetApiController struct {
	service      PetApiServicer
	errorHandler ErrorHandler
}

// PetApiOption for how the controller is set up.
type PetApiOption func(*PetApiController)

// WithPetApiErrorHandler inject ErrorHandler into controller
func WithPetApiErrorHandler(h ErrorHandler) PetApiOption {
	return func(c *PetApiController) {
		c.errorHandler = h
	}
}

// NewPetApiController creates a default api controller
func NewPetApiController(s PetApiServicer, opts ...PetApiOption) Router {
	controller := &PetApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PetApiController
func (c *PetApiController) Routes() Routes {
	return Routes{
		{
			"AddPet",
			strings.ToUpper("Post"),
			"/api/v3/pet",
			c.AddPet,
		},
		{
			"DeletePet",
			strings.ToUpper("Delete"),
			"/api/v3/pet/{petId}",
			c.DeletePet,
		},
		{
			"FindPetsByStatus",
			strings.ToUpper("Get"),
			"/api/v3/pet/findByStatus",
			c.FindPetsByStatus,
		},
		{
			"FindPetsByTags",
			strings.ToUpper("Get"),
			"/api/v3/pet/findByTags",
			c.FindPetsByTags,
		},
		{
			"GetPetById",
			strings.ToUpper("Get"),
			"/api/v3/pet/{petId}",
			c.GetPetById,
		},
		{
			"UpdatePet",
			strings.ToUpper("Put"),
			"/api/v3/pet",
			c.UpdatePet,
		},
		{
			"UpdatePetWithForm",
			strings.ToUpper("Post"),
			"/api/v3/pet/{petId}",
			c.UpdatePetWithForm,
		},
		{
			"UploadFile",
			strings.ToUpper("Post"),
			"/api/v3/pet/{petId}/uploadImage",
			c.UploadFile,
		},
	}
}

// AddPet - Add a new pet to the store
func (c *PetApiController) AddPet(w http.ResponseWriter, r *http.Request) {
	petParam := Pet{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&petParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPetRequired(petParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddPet(r.Context(), petParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeletePet - Deletes a pet
func (c *PetApiController) DeletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.DeletePet(r.Context(), petIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// FindPetsByStatus - Finds Pets by status
func (c *PetApiController) FindPetsByStatus(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	statusParam := query.Get("status")
	result, err := c.service.FindPetsByStatus(r.Context(), statusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// FindPetsByTags - Finds Pets by tags
func (c *PetApiController) FindPetsByTags(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tagsParam := strings.Split(query.Get("tags"), ",")
	result, err := c.service.FindPetsByTags(r.Context(), tagsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetPetById - Find pet by ID
func (c *PetApiController) GetPetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetPetById(r.Context(), petIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdatePet - Update an existing pet
func (c *PetApiController) UpdatePet(w http.ResponseWriter, r *http.Request) {
	petParam := Pet{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&petParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPetRequired(petParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdatePet(r.Context(), petParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdatePetWithForm - Updates a pet in the store with form data
func (c *PetApiController) UpdatePetWithForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	nameParam := query.Get("name")
	statusParam := query.Get("status")
	result, err := c.service.UpdatePetWithForm(r.Context(), petIdParam, nameParam, statusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UploadFile - uploads an image
func (c *PetApiController) UploadFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	additionalMetadataParam := query.Get("additionalMetadata")
	fileParam, err := ReadFormFileToTempFile(r, "file")
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	/*bodyParam := *os.File{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}*/
	result, err := c.service.UploadFile(r.Context(), petIdParam, additionalMetadataParam, fileParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
