package BrandService

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

func BrandInsert(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> BrandInsertService.go On ", now.Format("2006-01-02 15:04:05"))


	body := utils.BrandBody(request)
	bodyRepo := brandRepository(body)
	db := confiq.Connect()

	if bodyRepo.BrandName.String == ""{
		out.ResponseOut(response, nil, false, constanta.CodeBadRequestResponse, "Nama Brand tidak boleh kosong")
		return
	}

	err = dao.BrandDAO.InsertBrand(db, bodyRepo)
	if err != nil{
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(response, nil, true, constanta.CodeSuccessResponse, constanta.SuccessAddData)
	return nil
}

func brandRepository(body in.BrandRequest) model.BrandModel  {

	return model.BrandModel{
		ID     :  sql.NullInt64{Int64: body.Id},
		BrandName: sql.NullString{String: body.Brand},
	}

}
