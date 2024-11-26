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

func ProductUpdate(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> ProductUpdateService.go On ", now.Format("2006-01-02 15:04:05"))


	id, err := utils.ReadParam(request)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	body := utils.ProductBody(request)
	body.Id = id
	bodyRepo := productRepository(body)
	db := confiq.Connect()

	detailProduct, err := dao.ProductDAO.GetDetailProduct(db, id)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if detailProduct.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Produk tidak diketahui")
		return
	}

	if bodyRepo.ProductName.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Nama Produk tidak boleh kosong")
		return
	}

	if bodyRepo.Qty.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Qty tidak boleh kosong")
		return
	}

	if bodyRepo.Price.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Harga tidak boleh kosong")
		return
	}

	if bodyRepo.IdBrand.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "IdBrand tidak boleh kosong")
		return
	}

	detail, err := dao.BrandDAO.GetDetailBrand(db, bodyRepo.IdBrand.Int64)
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	if detail.ID.Int64 == 0{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Id Brand tidak diketahui")
		return
	}

	err = dao.ProductDAO.UpdateProduct(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessEditData)
	return nil
}
