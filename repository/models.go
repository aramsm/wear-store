package repository

type Address struct {
	City   string `json:"city" binding:"required"`
	State  string `json:"state" binding:"required"`
	Street string `json:"street" binding:"required" gorm:"primaryKey"`
}

type Employee struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Code  string `json:"code" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type Store struct {
	ID         string     `json:"id" gorm:"primaryKey"`
	Code       string     `json:"phone"`
	LabelBrand string     `json:"label_brand" binding:"required"`
	Name       string     `json:"store_name" binding:"required"`
	Phone      string     `json:"store_phone"`
	Address    Address    `json:"address"`
	Employees  []Employee `json:"employees"`
}
