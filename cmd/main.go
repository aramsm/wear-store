package main

type Address struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	City   string `json:"city" binding:"required"`
	State  string `json:"state" binding:"required"`
	Street string `json:"street" binding:"required"`
}

type Employee struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Code  string `json:"code" binding:"required"`
	Name  string `json:"employee_name" binding:"required"`
	Phone string `json:"employee_phone" binding:"required"`
}

type Store struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Code       string     `json:"phone"`
	LabelBrand string     `json:"label_brand" binding:"required"`
	Name       string     `json:"store_name" binding:"required"`
	Phone      string     `json:"store_phone"`
	Address    Address    `json:"address"`
	Employees  []Employee `json:"employees"`
}

func main() {

}
