// This package contains all the models needed for each endpoint
// Each endpoint has its own unique structure that represents
// all the parameters and the services needed to oparate
package request

import "github.com/go-redis/redis"

// All requests should have access to a specific "service"
// This generic request contains all logic our endpoints neetd to
// interact with.
//
// @see createSecretRequest
// @see findSecretRequest
type GenericSecretRequest struct {
	Redis redis.Client
}