package models

import (
	"github.com/jinzhu/gorm"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Signup struct
type Signup struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"not null"`
	Lastname   string `json:"lastname" gorm:"not null"`
	Email      string `json:"email" gorm:"not null;unique"`
	Class      string `json:"class" gorm:"not null"`
	ClassRoom  string `json:"classroom" gorm:"not null"`
}

// SignupManager struct
type SignupManager struct {
	db *DB
}

// NewSignupManager - Creates a new *SignupManager that can be used for managing all people who signed up.
func NewSignupManager(db *DB) (*SignupManager, error) {
	db.AutoMigrate(&Signup{})
	manager := SignupManager{}
	manager.db = db
	return &manager, nil
}

// Create - Add someone to the database
func (state SignupManager) Create(s *Signup) error {
	return state.db.Create(s).Error
}

// FindAll - Get all the people who signed up
func (state SignupManager) FindAll() ([]Signup, error) {
	s := []Signup{}
	err := state.db.Order("class desc").Find(&s).Error
	return s, err
}
