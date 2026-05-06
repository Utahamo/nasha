package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;not null;size:64"`
	PasswordHash string `gorm:"not null;size:255"`
	Role         string `gorm:"not null;default:viewer;size:16"`
	Disabled     bool   `gorm:"not null;default:false"`
}

type Mount struct {
	gorm.Model
	Name     string `gorm:"not null;size:64"`
	Type     string `gorm:"not null;size:16"`
	Path     string `gorm:"not null;size:255"`
	Config   string `gorm:"type:text"`
	Priority int    `gorm:"not null;default:0"`
}

type ShareLink struct {
	gorm.Model
	Token         string `gorm:"uniqueIndex;not null;size:64"`
	Path          string `gorm:"not null;size:512"`
	CreatedBy     uint
	ExpiresAt     *time.Time
	Password      string `gorm:"size:255"`
	DownloadCount uint   `gorm:"not null;default:0"`
}
