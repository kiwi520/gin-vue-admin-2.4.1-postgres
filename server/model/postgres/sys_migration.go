package postgres

import "time"

type SysMigration struct {
	Migrate bool `json:"migrate" gorm:"bool;comment:是否初始化"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
