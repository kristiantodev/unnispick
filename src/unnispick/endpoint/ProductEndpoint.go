package endpoint

import (
	"net/http"
	"unnispick/service/ProdukService"
)

func ProductEndpointWithId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "PUT":
		ProdukService.ProductUpdate(response, request)
		break
	case "DELETE":
		ProdukService.DeleteProduct(response, request)
		break
	}
}

func ProductEndpointWithoutId(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		ProdukService.ProductInsert(response, request)
		break
	case "GET":
		ProdukService.GetProductList(response, request)
		break
	}
}
