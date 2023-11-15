package models

type User struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Products represents a product

type Products struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	ProductName string `json:"productName" gorm:"column:productname"`
	Category    string
	Price       int
	Quantity    int
}
