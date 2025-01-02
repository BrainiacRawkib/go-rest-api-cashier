package models

import "time"

type Cashier struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}

type Category struct {
	Id          int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}

type Discount struct {
	Id              int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       int       `json:"expiredAt"`
	ExpiredAtFormat string    `json:"expiredAtFormat"`
	StringFormat    string    `json:"stringFormat"`
	DateCreated     time.Time `json:"dateCreated"`
	DateUpdated     time.Time `json:"dateUpdated"`
}

type Order struct {
	Id             int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	CashierId      int       `json:"cashierId"`
	PaymentTypesId int       `json:"paymentTypesId"`
	TotalPrice     int       `json:"totalPrice"`
	TotalPaid      int       `json:"totalPaid"`
	TotalReturn    int       `json:"totalReturn"`
	ReceiptId      string    `json:"receiptId"`
	IsDownload     int       `json:"isDownload"`
	ProductId      string    `json:"productId"`
	Quantities     string    `json:"quantities"`
	DateCreated    time.Time `json:"dateCreated"`
	DateUpdated    time.Time `json:"dateUpdated"`
}
