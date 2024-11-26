package out

type ProductResponse struct {
	Id             int64    `json:"-"`
	ProductName    string   `json:"nama_produk"`
	Price          int64    `json:"harga"`
	Qty            int64    `json:"qty"`
	BrandId        int64    `json:"id_brand"`
	Brand          string   `json:"brand"`
}
