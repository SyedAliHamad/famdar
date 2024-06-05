package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	Location  postgres.Hstore `gorm:"type:geometry(Point,4326);"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
}
