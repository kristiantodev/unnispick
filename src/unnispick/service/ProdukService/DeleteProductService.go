package ProdukService

import (
	"fmt"
	"net/http"
	"time"
	"unnispick/confiq"
	"unnispick/constanta"
	"unnispick/dao"
	"unnispick/dto/out"
	"unnispick/utils"
)

func DeleteProduct(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> DeleteProductService.go On ", now.Format("2006-01-02 15:04:05"))

	id, err := utils.ReadParam(request)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}
	db := confiq.Connect()

	detail, err := dao.ProductDAO.GetDetailProduct(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if detail.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Produk tidak diketahui")
		return
	}

	err = dao.ProductDAO.DeleteProduct(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessDeleteData)
	return nil
}
