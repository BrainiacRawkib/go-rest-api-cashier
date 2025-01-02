package models

import "time"

type Payment struct {
	Id            uint      `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	PaymentTypeId int       `json:"paymentTypeId"`
	Logo          string    `json:"logo"`
	DateCreated   time.Time `json:"dateCreated"`
	DateUpdated   time.Time `json:"dateUpdated"`
}
