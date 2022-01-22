package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Parent     *Board
	ParentID   uint
	Flashcards []Flashcard
	Name       string
	Creator    string
	Users      pq.StringArray `gorm:"type:text[]"`
}

type Flashcard struct {
	gorm.Model
	BoardID uint
	Title   string
	Answer  string
}
