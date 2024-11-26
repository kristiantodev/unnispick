package ProdukService

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"unnispick/confiq"
	"unnispick/constanta"
	"unnispick/dao"
	"unnispick/dto/in"
	"unnispick/dto/out"
	"unnispick/model"
	"unnispick/utils"
)

func ProductInsert(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> ProductInsertService.go On ", now.Format("2006-01-02 15:04:05"))


	body := utils.ProductBody(request)
	bodyRepo := productRepository(body)
	db := confiq.Connect()

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

	err = dao.ProductDAO.InsertProduct(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessAddData)
	return nil
}

func productRepository(body in.ProductRequest) model.ProductModel  {

	return model.ProductModel{
		ID     :  sql.NullInt64{Int64: body.Id},
		ProductName: sql.NullString{String: body.ProductName},
		Price:  sql.NullInt64{Int64: body.Price},
		Qty:  sql.NullInt64{Int64: body.Qty},
		IdBrand:  sql.NullInt64{Int64: body.BrandId},
	}

}
