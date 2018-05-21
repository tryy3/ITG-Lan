package models

import (
	"time"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// FAQ struct
type FAQ struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Question  string     `json:"question" gorm:"not null"`
	Answer    string     `json:"answer"`
}

// FAQManager struct
type FAQManager struct {
	db *DB
}

// NewFAQManager - Creates a new *FAQManager that can be used for managing all people who sent a question
func NewFAQManager(db *DB) (*FAQManager, error) {
	db.AutoMigrate(&FAQ{})
	manager := FAQManager{}
	manager.db = db
	return &manager, nil
}

// Create - Add a question to the database
func (state FAQManager) Create(faq *FAQ) error {
	return state.db.Create(faq).Error
}

// FindAll - Get all questions
func (state FAQManager) FindAll() ([]FAQ, error) {
	faqs := []FAQ{}
	err := state.db.Find(&faqs).Error
	return faqs, err
}

// FindAnswered - Get all questions that is answered
func (state FAQManager) FindAnswered() ([]FAQ, error) {
	faqs := []FAQ{}
	err := state.db.Not("answer", "").Find(&faqs).Error
	return faqs, err
}

// Find - Find a specific question
func (state FAQManager) Find(id uint) (*FAQ, error) {
	faq := FAQ{}
	err := state.db.Where("id = ?", id).Find(&faq).Error
	return &faq, err
}

// Save - saves a question to the database
func (state FAQManager) Save(faq *FAQ) error {
	return state.db.Save(&faq).Error
}
