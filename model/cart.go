package model

type Detail struct {
	ID_product int64 `json: "id_product"`
	Quantity   int8  `json: "quantity"`
}

type Cart struct {
	Products []Detail
}
