package models

import "time"

type Category struct {
	Id          int       `json:"Id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}
