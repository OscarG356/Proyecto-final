package models

type Carro struct {
	Id          int    `json:"id" db:"id"`
	Marca       string `json:"marca" db:"marca"`
	Modelo      string `json:"modelo" db:"modelo"`
	Tipo        string `json:"tipo" db:"tipo"`
	Transmision string `json:"transmision" db:"transmision"`
	Combustible string `json:"combustible" db:"combustible"`
	Ubicacion   string `json:"ubicacion" db:"ubicacion"`
	Reservado   int    `json:"reservado" db:"reservado"`
	Imagen      string `json:"imagen" db:"imagen"`
}
