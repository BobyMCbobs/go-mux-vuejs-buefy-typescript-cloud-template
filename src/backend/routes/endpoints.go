package routes

import (
	"net/http"

	"gitlab.com/bobymcbobs/go-mux-vuejs-buefy-typescript-cloud-template/src/backend/types"
)

func GetEndpoints(endpointPrefix string) types.Endpoints {
	return types.Endpoints{
		{
			EndpointPath: endpointPrefix + "/hello",
			HandlerFunc:  APIgetHello,
			HttpMethod:   http.MethodGet,
		},
	}
}
