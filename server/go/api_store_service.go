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
	"context"
	"net/http"
	"errors"
)

// StoreApiService is a service that implements the logic for the StoreApiServicer
// This service should implement the business logic for every endpoint for the StoreApi API.
// Include any external packages or services that will be required by this service.
type StoreApiService struct {
}

// NewStoreApiService creates a default api service
func NewStoreApiService() StoreApiServicer {
	return &StoreApiService{}
}

// DeleteOrder - Delete purchase order by ID
func (s *StoreApiService) DeleteOrder(ctx context.Context, orderId int64) (ImplResponse, error) {
	// TODO - update DeleteOrder with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteOrder method not implemented")
}

// GetInventory - Returns pet inventories by status
func (s *StoreApiService) GetInventory(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetInventory with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]int32{}) or use other options such as http.Ok ...
	//return Response(200, map[string]int32{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetInventory method not implemented")
}

// GetOrderById - Find purchase order by ID
func (s *StoreApiService) GetOrderById(ctx context.Context, orderId int64) (ImplResponse, error) {
	// TODO - update GetOrderById with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Order{}) or use other options such as http.Ok ...
	//return Response(200, Order{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetOrderById method not implemented")
}

// PlaceOrder - Place an order for a pet
func (s *StoreApiService) PlaceOrder(ctx context.Context, order Order) (ImplResponse, error) {
	// TODO - update PlaceOrder with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Order{}) or use other options such as http.Ok ...
	//return Response(200, Order{}), nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("PlaceOrder method not implemented")
}
