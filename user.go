package common

import (
	"time"

	"github.com/ecletus/roles"

	"github.com/moisespsena-go/aorm"
)

// CurrentUser is an interface, which is used for qor admin to get current logged user
type User interface {
	roles.RoleHaver
	DisplayName() string
	GetEmail() string
	GetName() string
	GetDefaultLocale() string
	GetTimeLocation() *time.Location
	GetLocales() []string
	GetRoles() roles.Roles
	GetUID() string
	IsSuper() bool
}

var DB_USER = pkg + ".user"

func SetUserToDB(db *aorm.DB, user User) *aorm.DB {
	return db.Set(DB_USER, user)
}

func GetUserFromDB(db *aorm.DB) User {
	if v, ok := db.Get(DB_USER); ok {
		return v.(User)
	}
	return nil
}

func GetUserFromScope(scope *aorm.Scope) User {
	if v, ok := scope.Get(DB_USER); ok {
		return v.(User)
	}
	return nil
}
