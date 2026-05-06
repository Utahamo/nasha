package db

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Open(dsn string) (*DB, error) {
	g, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{g}, nil
}

func (d *DB) AutoMigrate() error {
	return d.DB.AutoMigrate(&User{}, &Mount{}, &ShareLink{})
}

func (d *DB) Seed() error {
	var count int64
	d.DB.Model(&User{}).Count(&count)
	if count > 0 {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return d.DB.Create(&User{
		Username:     "admin",
		PasswordHash: string(hash),
		Role:         "admin",
	}).Error
}
