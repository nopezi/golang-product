package structs

type Product struct {
	Id int `json:"id"`
	NameProduct string `json:"nama_product"`
	Price int `json:"price"`
	Image string `json:"image"`
	CreatedAt string `json:"created_at"`
}

type UpdateProduct struct {
	Id int `json:"id"`
	NameProduct string `json:"nama_produk"`
	Price int `json:"harga"`
}

type InputProduct struct {
	NameProduct string `json:"nama_produk"`
	Price int `json:"harga"`
	CreatedAt string
}

type Images struct {
	IdProduct int64 `json:"id_produk"`
	Name string `json:"image"`
}

type ImageProduct struct {
	Id int64 `json:"id"`
	IdProduct int64 `json:"id_produk"`
	Name string `json:"image"`
}

type UpdateImages struct {
	Id int64 `json:"id"`
	Name string `json:"image"`
}

type DeleteImageProduct struct {
	Id int64
}

type CekTableImage struct {
	TableSchema string
}