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
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}
