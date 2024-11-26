package endpoint

import (
	"net/http"
	"unnispick/service/BrandService"
)

func BrandEntpointWithId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "DELETE":
		BrandService.DeleteBrand(response, request)
		break
	}
}

func BrandEndpointWithoutId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		BrandService.BrandInsert(response, request)
		break
	}
}
