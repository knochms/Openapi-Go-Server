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

type User struct {

	Id int64 `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	Phone string `json:"phone,omitempty"`

	// User Status
	UserStatus int32 `json:"userStatus,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	return nil
}

// AssertRecurseUserRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of User (e.g. [][]User), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUser, ok := obj.(User)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserRequired(aUser)
	})
}
