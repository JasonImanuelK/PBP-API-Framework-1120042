package model

type DataCelana struct {
	ID        int    `json:"ID"`
	Nama      string `json:"Name"`
	Harga     int    `json:"Harga"`
	Stok      int    `json:"Stok"`
	Deskripsi string `json:"Deskripsi"`
}

type ResponseCelana struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    DataCelana `json:"Data"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
