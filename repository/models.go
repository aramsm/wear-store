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
	BrandLabel string     `json:"brand_label" binding:"required"`
	Name       string     `json:"name" binding:"required"`
	Phone      string     `json:"phone"`
	Address    Address    `json:"address"`
	Employees  []Employee `json:"employees"`
}
