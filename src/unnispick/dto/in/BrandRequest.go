package in

type BrandRequest struct {
	Id       int64    `json:"-"`
	Brand    string   `json:"nama_brand"`
}
