package out

import (
	"encoding/json"
	"net/http"
	"unnispick/utils"
)

type APIResponse struct {
	Unnispick unnispickMessage `json:"unnispick"`
}

type unnispickMessage struct {
	Success bool     `json:"success"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Content interface{} `json:"content"`
}

func (ar APIResponse) String() string {
	return utils.StructToJSON(ar)
}

func ResponseOut(response http.ResponseWriter, data interface{}, success bool, code int, message string){
	response.Header().Set("Content-type", "application/json")
	var apiResponse APIResponse
	apiResponse.Unnispick.Success = success
	apiResponse.Unnispick.Code = code
	apiResponse.Unnispick.Message = message
	apiResponse.Unnispick.Content = data
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(apiResponse)
}