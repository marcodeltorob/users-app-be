package model

import (
	"github.com/jinzhu/gorm"
)

// User represents the columns we use from the 'users' table
type User struct {
	gorm.Model
	ID     int
	name   string
	email  string `gorm:"type:varchar(100);unique_index"`
	status bool
}
