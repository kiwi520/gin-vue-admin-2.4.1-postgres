package postgres

import (
	"gorm.io/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
