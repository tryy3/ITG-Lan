package models

import (
	"time"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Tournament struct
type Tournament struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Image     string     `json:"image"`
	Amount    uint       `json:"amount"`
	Name      string     `json:"name"`
	Active    bool       `json:"active"`
}

// TournamentManager struct
type TournamentManager struct {
	db *DB
}

// NewTournamentManager - Creates a new *TournamentManager that can be used for managing all the tournaments
func NewTournamentManager(db *DB) (*TournamentManager, error) {
	db.AutoMigrate(&Tournament{})
	manager := TournamentManager{}
	manager.db = db
	return &manager, nil
}

// Create - Add a tournament
func (state TournamentManager) Create(t *Tournament) error {
	return state.db.Create(t).Error
}

// FindAll - Get all the tournaments
func (state TournamentManager) FindAll() ([]Tournament, error) {
	t := []Tournament{}
	err := state.db.Find(&t).Error
	return t, err
}

// FindActive - Get all active tournaments
func (state TournamentManager) FindActive() ([]Tournament, error) {
	t := []Tournament{}
	err := state.db.Where("active = 1").Find(&t).Error
	return t, err
}

// Find - Find a specific tournament based on the ID
func (state TournamentManager) Find(id uint) (*Tournament, error) {
	t := Tournament{}
	err := state.db.Where("id = ?", id).Find(&t).Error
	return &t, err
}

// Save - Save a tournament
func (state TournamentManager) Save(t *Tournament) error {
	return state.db.Save(t).Error
}

// Delete removes an existing tournament from the database
func (state TournamentManager) Delete(t *Tournament) error {
	return state.db.Delete(&t).Error
}

// DeleteWithID is a wrapper for Delete the allows deleting a tournament
// with the id instead of the tournament struct
func (state TournamentManager) DeleteWithID(id uint) error {
	t := &Tournament{ID: id}
	return state.Delete(t)
}
