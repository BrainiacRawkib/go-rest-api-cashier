package models

import "time"

type Product struct {
	Id               int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Image            string    `json:"image"`
	TotalFinalPrice  int       `json:"totalFinalPrice"`
	TotalNormalPrice int       `json:"totalNormalPrice"`
	DateCreated      time.Time `json:"dateCreated"`
	DateUpdated      time.Time `json:"dateUpdated"`
	CategoryId       int       `json:"categoryId"`
	DiscountId       int       `json:"discountId"`
}

type ProductResult struct {
	Id       int      `json:"productId" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku      string   `json:"sku"`
	Name     string   `json:"name"`
	Stock    int      `json:"stock"`
	Price    int      `json:"price"`
	Image    string   `json:"image"`
	Category Category `json:"category"`
	Discount Discount `json:"discount"`
}
