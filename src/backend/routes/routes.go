/*
	route related
*/

package routes

import (
	"net/http"

	"gitlab.com/bobymcbobs/go-mux-vuejs-buefy-typescript-cloud-template/src/backend/common"
	"gitlab.com/bobymcbobs/go-mux-vuejs-buefy-typescript-cloud-template/src/backend/types"
)

func APIroot(w http.ResponseWriter, r *http.Request) {
	// root of API
	JSONresp := types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "Hit root of webserver",
		},
	}
	common.JSONResponse(r, w, 200, JSONresp)
}

func APIgetHello(w http.ResponseWriter, r *http.Request) {
	// root of API
	helloMsg := common.SayHello()
	JSONresp := types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: helloMsg,
		},
	}
	common.JSONResponse(r, w, 200, JSONresp)
}

func APIUnknownEndpoint(w http.ResponseWriter, r *http.Request) {
	common.JSONResponse(r, w, 404, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "This endpoint doesn't seem to exist.",
		},
	})
}

