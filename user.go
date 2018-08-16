package common

import (
	"github.com/jinzhu/gorm"
)

// CurrentUser is an interface, which is used for qor admin to get current logged user
type User interface {
	WithID
	DisplayName() string
	GetEmail() string
}

var DB_USER = PREFIX + ".user"

func SetUserToDB(db *gorm.DB, user User) *gorm.DB {
	return db.Set(DB_USER, user)
}

func GetUserFromDB(db *gorm.DB) User {
	if v, ok := db.Get(DB_USER); ok {
		return v.(User)
	}
	return nil
}

func GetUserFromScope(scope *gorm.Scope) User {
	if v, ok := scope.Get(DB_USER); ok {
		return v.(User)
	}
	return nil
}
