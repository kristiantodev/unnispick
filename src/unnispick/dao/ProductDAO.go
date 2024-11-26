package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"unnispick/model"
)

type productDAO struct {
	AbstractDAO
}

var ProductDAO = productDAO{}.New()

func (input productDAO) New() (output productDAO) {
	output.TableName = "produk"
	output.FileName = "ProductDAO.go"
	return
}

func (input productDAO) InsertProduct(db *sql.DB, inputStruct model.ProductModel) (err error) {
	query := "INSERT INTO " + input.TableName + " (" +
		" nama_produk, harga, qty, id_brand)  " +
		"VALUES (?, ?, ?, ?) "
	params := []interface{}{
		inputStruct.ProductName.String, inputStruct.Price.Int64, inputStruct.Qty.Int64, inputStruct.IdBrand.Int64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (input productDAO) UpdateProduct(db *sql.DB, inputStruct model.ProductModel) (err error) {
	query := "UPDATE " + input.TableName + " SET " +
		"nama_produk = ?, harga = ?, qty = ?, id_brand = ? " +
		"WHERE id = ?"

	params := []interface{}{
		inputStruct.ProductName.String, inputStruct.Price.Int64,
		inputStruct.Qty.Int64, inputStruct.IdBrand.Int64,
		inputStruct.ID.Int64,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (input productDAO) GetDetailProduct(db *sql.DB, id int64) (model.ProductModel, error) {
	query := "SELECT id, nama_produk " +
		"FROM " + input.TableName + " WHERE id = ?"

	row := db.QueryRow(query, id)
	var produk model.ProductModel

	err := row.Scan(&produk.ID, &produk.ProductName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return produk, nil
		}
		return produk, err
	}
	return produk, nil
}

func (input productDAO) GetProductList(db *sql.DB, param CustomQueryModel) (results []model.ProductModel, err error) {
	query := "SELECT p.id, p.nama_produk, p.harga, p.qty, p.id_brand, b.nama_brand " +
		" FROM " + input.TableName + " p" +
		" LEFT JOIN brand b ON b.id = p.id_brand "

	var params []interface{}

	if param.Page != "" && param.Limit != "" {
		page, err := strconv.Atoi(param.Page)
		if err != nil {
			return nil, err
		}
		limit, err := strconv.Atoi(param.Limit)
		if err != nil {
			return nil, err
		}
		offset := limit * (page - 1)

		query += " LIMIT ? OFFSET ?"
		params = append(params, limit, offset)
	}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.ProductModel
		if err := rows.Scan(&data.ID,
			&data.ProductName,
			&data.Price,
			&data.Qty,
			&data.IdBrand,
			&data.Brand); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (input productDAO) DeleteProduct(db *sql.DB, id int64) (err error) {
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
