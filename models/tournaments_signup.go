package models

import (
	"time"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TournamentSignup struct
type TournamentSignup struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
	Team         string     `json:"team"`
	Name         string     `json:"name"`
	IGN          string     `json:"ign"`
	Leader       bool       `json:"leader"`
	Email        string     `json:"email"`
	IDTournament uint       `json:"id_tournament"`
}

// TournamentSignupManager struct
type TournamentSignupManager struct {
	db *DB
}

// NewTournamentSignupManager - Creates a new *TournamentSignupManager that can be used for managing all the tournament signups
func NewTournamentSignupManager(db *DB) (*TournamentSignupManager, error) {
	db.AutoMigrate(&TournamentSignup{})
	manager := TournamentSignupManager{}
	manager.db = db
	return &manager, nil
}

// Create - Add a TournamentSignup
func (state TournamentSignupManager) Create(t *TournamentSignup) error {
	return state.db.Create(t).Error
}

// FindAll - Get all the tournament signups
func (state TournamentSignupManager) FindAll() ([]TournamentSignup, error) {
	t := []TournamentSignup{}
	err := state.db.Find(&t).Error
	return t, err
}

// FindTeams - Get a list of all signed up teams
func (state TournamentSignupManager) FindTeams(id uint) ([]string, error) {
	t := []string{}
	err := state.db.Model(&TournamentSignup{}).Where("id_tournament = ?", id).Pluck("team", &t).Error
	return t, err
}

// FindAllTeams - Get a list of all signed up people in a specific tournament
func (state TournamentSignupManager) FindAllTeams(id uint) ([]TournamentSignup, error) {
	t := []TournamentSignup{}
	err := state.db.Where("id_tournament = ?", id).Find(&t).Error
	return t, err
}

// Exists - Checks if a team already exists
func (state TournamentSignupManager) Exists(name string) bool {
	if err := state.db.Where("team=?", name).Find(&TournamentSignup{}).Error; err != nil {
		return false
	}
	return true
}

// Delete removes an existing tournament from the database
func (state TournamentSignupManager) Delete(t *TournamentSignup) error {
	return state.db.Delete(&t).Error
}

// DeleteWithTournamentID is a wrapper for Delete the allows deleting a tournament
// with the tournament id instead of the tournament struct
func (state TournamentSignupManager) DeleteWithTournamentID(id uint) error {
	return state.db.Where("id_tournament = ?", id).Delete(TournamentSignup{}).Error
}
