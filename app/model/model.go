package model

import (
	"github.com/jinzhu/gorm"
	// For using mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Employee struct
type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    string `json:"age"`
	Status bool   `json:"status"`
}

// Disable false an eployee status
func (e *Employee) Disable() {
	e.Status = false
}

// Enable true and employee status
func (e *Employee) Enable() {
	e.Status = true
}

// DBMigrate will create and migrate the table, and then make some relationship if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}
