package model

import "database/sql"

type ProductModel struct {
	ID           sql.NullInt64
	ProductName  sql.NullString
	Price        sql.NullInt64
	Qty          sql.NullInt64
	IdBrand      sql.NullInt64
	Brand        sql.NullString
}
