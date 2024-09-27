package models

type Product struct {
	Pid         uint    `gorm:"primary_key;auto_increment;not_null"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
