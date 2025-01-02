package models

import "time"

type Cashier struct {
	Id          uint      `json:"Id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}
