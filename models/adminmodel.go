// models/user.go

package models

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Username string
	Password string
}
