package model

import "database/sql"

type BrandModel struct {
	ID          sql.NullInt64
	BrandName   sql.NullString
}
