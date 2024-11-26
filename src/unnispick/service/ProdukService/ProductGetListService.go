package ProdukService

import (
	"fmt"
	"net/http"
	"time"
	"unnispick/confiq"
	"unnispick/constanta"
	"unnispick/dao"
	"unnispick/dto/out"
	"unnispick/model"
)

func GetProductList(response http.ResponseWriter, request *http.Request) (err error) {
	now := time.Now()
	fmt.Println("HIT -> GetProductListService.go On ", now.Format("2006-01-02 15:04:05"))
	db := confiq.Connect()

	params := request.URL.Query()
	if err != nil {
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	datas, err := dao.ProductDAO.GetProductList(db, dao.CustomQueryModel{
		Page:  params.Get("page"),
		Limit: params.Get("limit"),
	})

	if err != nil {
		fmt.Println("Error Apa ya", err.Error())
		out.ResponseOut(response, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}


	db.Close()
	out.ResponseOut(response, convertRepoToDTO(datas), true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

func convertRepoToDTO(datas []model.ProductModel) (output []out.ProductResponse) {
	for i:=0;i<len(datas);i++ {
		output = append(output, out.ProductResponse{
			Id:            datas[i].ID.Int64,
			ProductName:   datas[i].ProductName.String,
			Price:         datas[i].Price.Int64,
			Qty:           datas[i].Qty.Int64,
			BrandId:       datas[i].IdBrand.Int64,
			Brand:         datas[i].Brand.String,
		})
	}
	return output
}
