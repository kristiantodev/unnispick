package dao

import (
	"database/sql"
	"fmt"
	"unnispick/model"
)

type brandDAO struct {
	AbstractDAO
}

var BrandDAO = brandDAO{}.New()

func (input brandDAO) New() (output brandDAO) {
	output.TableName = "brand"
	output.FileName = "BrandDAO.go"
	return
}

func (input brandDAO) InsertBrand(db *sql.DB, inputStruct model.BrandModel) (err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" nama_brand)  " +
		"VALUES (?) "
	params := []interface{}{
		inputStruct.BrandName.String,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (input brandDAO) DeleteBrand(db *sql.DB, id int64) (err error) {
	query := "DELETE FROM " + input.TableName + " " +
		"WHERE id = ? "
	params := []interface{}{
		id,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (input brandDAO) GetDetailBrand(db *sql.DB, id int64) (model.BrandModel, error) {
	query := "SELECT id, nama_brand " +
		"FROM " + input.TableName + " WHERE id = ?"

	row := db.QueryRow(query, id)
	var brand model.BrandModel

	err := row.Scan(&brand.ID, &brand.BrandName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return brand, nil
		}
		return brand, err
	}
	return brand, nil
}
